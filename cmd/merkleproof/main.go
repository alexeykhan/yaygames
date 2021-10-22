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

const (
	rpcBSCEndpoint = "https://data-seed-prebsc-1-s1.binance.org:8545"
	contractHex    = "0xCd71C520CB6358D3B455f1d018813FF667001618"
	sourceFilePath = "./src/mercleTreeData.json"
)

func main() {
	if err := run(sourceFilePath); err != nil {
		log.Fatal(err)
	}
}

func run(path string) (err error) {
	var (
		mercleTreeFile  *os.File
		mercleTreeBytes []byte
		data            Data
		valid           bool
	)

	// Open source file.
	if mercleTreeFile, err = os.Open(path); err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer func() {
		if closeErr := mercleTreeFile.Close(); closeErr != nil {
			if err == nil {
				err = fmt.Errorf("close file: %w", closeErr)
			} else {
				err = fmt.Errorf("close file: %w: %v", closeErr, err)
			}
		}
	}()

	// Read from file.
	if mercleTreeBytes, err = ioutil.ReadAll(mercleTreeFile); err != nil {
		return fmt.Errorf("read file data: %v", err)
	}

	// Unmarshal JSON to struct.
	if err = json.Unmarshal(mercleTreeBytes, &data); err != nil {
		return fmt.Errorf("unmarshal file data: %v", err)
	}

	// Hash initial leaves.
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
	tree := merkle.NewTree(leaves, merkle.Options{SortPairs: true})

	// Format data into hex strings.
	treeHexProof := tree.GetHexProof(leaves[0])
	proof := make([]string, len(treeHexProof))
	for i, p := range treeHexProof {
		proof[i] = hexutil.Encode(p[:])
	}

	// Display proof leaves.
	mergedArray := strings.Join(proof, ",")
	fmt.Printf("Generated Proof:\n[%s]\n\n", mergedArray)

	// Check claim.
	if valid, err = checkClaim(data.Nodes[0], treeHexProof); err != nil {
		return fmt.Errorf("check claim: %w", err)
	}

	if valid {
		fmt.Println("Huray! Valid proof.")
	} else {
		fmt.Println("Whoops! Invalid proof.")
	}

	return
}

func checkClaim(node Node, proof [][32]byte) (valid bool, err error) {
	var (
		client   *ethclient.Client
		instance *yayvesting.Yayvesting
	)

	if client, err = ethclient.Dial(rpcBSCEndpoint); err != nil {
		return false, err
	}

	contractAddress := common.HexToAddress(contractHex)
	if instance, err = yayvesting.NewYayvesting(contractAddress, client); err != nil {
		return false, err
	}

	return instance.CheckClaim(nil,
		common.HexToAddress(node.Address),
		big.NewInt(node.Category),
		big.NewInt(node.Amount),
		proof)
}
