package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/miguelmota/go-solidity-sha3"

	"github.com/alexeykhan/yaygames/internal/config"
	"github.com/alexeykhan/yaygames/pkg/merkle"
	"github.com/alexeykhan/yaygames/pkg/yayvesting"
)

type (
	// Data describes tree structure.
	Data struct {
		Nodes []Node `json:"nodes"`
	}
	// Node describes node structure.
	Node struct {
		Address  string `json:"address"`
		Category int64  `json:"category"`
		Amount   int64  `json:"amount"`
	}
)

const configFilePath = "./src/config.yaml"

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg, err := config.New(configFilePath)
	if err != nil {
		return fmt.Errorf("get config: %w", err)
	}

	// Open source file.
	mercleTreeFile, err := os.Open(cfg.MerkleFilePath())
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer mercleTreeFile.Close()

	// Read from file.
	log.Println("Reading data from source file...")
	mercleTreeBytes, err := ioutil.ReadAll(mercleTreeFile)
	if err != nil {
		return fmt.Errorf("read file data: %v", err)
	}

	// Unmarshal JSON to struct.
	data := Data{}
	if err := json.Unmarshal(mercleTreeBytes, &data); err != nil {
		return fmt.Errorf("unmarshal file data: %v", err)
	}

	// Hash initial leaves.
	log.Println("Hashing initial leaves...")
	leaves := make([][]byte, len(data.Nodes))
	for i, node := range data.Nodes {
		category := strconv.FormatInt(node.Category, 10)
		amount := strconv.FormatInt(node.Amount, 10)
		leaves[i] = solsha3.SoliditySHA3(
			[]string{"address", "uint256", "uint256"},
			[]interface{}{node.Address, category, amount},
		)
	}

	// Build merkle tree from data.
	log.Println("Building Merkle Tree from leaves...")
	tree := merkle.NewTree(leaves, merkle.Options{SortPairs: true})

	// Format data into hex strings.
	treeHexProof := tree.GetHexProof(leaves[0])
	proof := make([]string, len(treeHexProof))
	for i, p := range treeHexProof {
		proof[i] = hexutil.Encode(p[:])
	}

	// Display proof leaves.
	mergedArray := strings.Join(proof, ",")
	log.Printf("Generated Proof: [%s]\n", mergedArray)

	// Check claim.
	client, err := ethclient.Dial(cfg.MerkleEndpoint())
	if err != nil {
		return fmt.Errorf("init client: %w", err)
	}

	contractAddress := common.HexToAddress(cfg.MerkleContract())
	instance, err := yayvesting.NewYayvesting(contractAddress, client)
	if err != nil {
		return fmt.Errorf("init contract instance: %w", err)
	}

	log.Println("Checking claim via smart contract...")
	valid, err := instance.CheckClaim(nil,
		common.HexToAddress(data.Nodes[0].Address),
		big.NewInt(data.Nodes[0].Category),
		big.NewInt(data.Nodes[0].Amount),
		treeHexProof)
	if err != nil {
		return fmt.Errorf("check claim: %w", err)
	}

	if valid {
		log.Println("Huray! Valid proof.")
	} else {
		log.Println("Whoops! Invalid proof.")
	}

	return nil
}