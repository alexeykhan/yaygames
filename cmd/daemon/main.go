package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/alexeykhan/yaygames/internal/config"
	"github.com/alexeykhan/yaygames/pkg/btcvseth"
)

const configFilePath = "./src/config.yaml"

type (
	app struct {
		config   config.Config
		client   *ethclient.Client
		contract *btcvseth.Btcvseth
	}
)

func main() {
	cfg, err := config.New(configFilePath)
	if err != nil {
		log.Fatal(fmt.Sprintf("get config: %v", err))
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	a := &app{config: cfg}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		select {
		case <-ctx.Done():
			return
		case sig := <-c:
			log.Println("OS exit signal received:", sig.String())
			log.Println("starting graceful shutdown:", ctx.Err())
			cancel()
		}
	}()

	// Init client.
	a.client, err = ethclient.Dial(a.config.DaemonEndpoint())
	if err != nil {
		log.Fatal(fmt.Errorf("init client: %v", err))
	}

	// Connect to smart contract.
	address := common.HexToAddress(a.config.DaemonContract())
	a.contract, err = btcvseth.NewBtcvseth(address, a.client)
	if err != nil {
		log.Fatal(fmt.Errorf("init contract instance: %v", err))
	}

	if err = a.run(ctx); err != nil {
		log.Fatal(err)
	}
}

func (a *app) run(ctx context.Context) error {
	log.Println("ExecuteRound Daemon Started")

	daemonTicker := time.NewTicker(a.config.DaemonInterval())
	defer daemonTicker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-daemonTicker.C:
			if err := a.executeRound(ctx); err != nil {
				// Don't return, ignore and log local errors.
				// e.g.: "get block by number: cannot query unfinalized data"
				log.Println("error: execute round:", err)
			}
		}
	}
}

func (a *app) executeRound(ctx context.Context) error {
	// Pull header to get latest block.
	header, err := a.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return fmt.Errorf("get header: %w", err)
	}
	block, err := a.client.BlockByNumber(ctx, header.Number)
	if err != nil {
		return fmt.Errorf("get block by number: %w", err)
	}

	// Get interval blocks, current epoch and rounds.
	interval, err := a.contract.IntervalBlocks(nil)
	if err != nil {
		return fmt.Errorf("get interval blocks: %w", err)
	}
	epoch, err := a.contract.CurrentEpoch(nil)
	if err != nil {
		return fmt.Errorf("get current epoch: %w", err)
	}
	rounds, err := a.contract.Rounds(nil, epoch)
	if err != nil {
		return fmt.Errorf("get rounds: %w", err)
	}

	sum, sub := big.NewInt(0), big.NewInt(0)
	sum.Add(rounds.StartBlock, interval)
	sub.Sub(sum, block.Number())

	log.Println("——————————————————————")
	log.Println("Start Block:", rounds.StartBlock)
	log.Println("Current Block:", block.Number())
	log.Println("Next Round Border:", sum)
	log.Println("Blocks Left:", sub)

	if sum.Cmp(block.Number()) == -1 {
		log.Println("Rounds: Execute!")

		tx, err := a.contract.ExecuteRound(nil)
		if err != nil {
			return fmt.Errorf("execute round: %w", err)
		}

		log.Println("Successful Transaction:", tx.Hash().Hex())
	}

	return nil
}
