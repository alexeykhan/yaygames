package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/miguelmota/go-solidity-sha3"

	"github.com/alexeykhan/yaygames/pkg/merkle"
)

// TreeData describes structure of given source file.
type TreeData struct {
	Nodes []struct {
		Address  string `json:"address"`
		Category uint64 `json:"category"`
		Amount   uint64 `json:"amount"`
	} `json:"nodes"`
}

func main() {
	const path = "./src/mercleTreeData.json"

	if err := printMerkleProof(path); err != nil {
		log.Printf("parse data: %v", err)
		os.Exit(1)
	}
}

func printMerkleProof(path string) (err error) {
	var (
		mercleTreeFile  *os.File
		mercleTreeBytes []byte
		data            TreeData
	)

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

	if mercleTreeBytes, err = ioutil.ReadAll(mercleTreeFile); err != nil {
		return fmt.Errorf("read file data: %v", err)
	}

	if err = json.Unmarshal(mercleTreeBytes, &data); err != nil {
		return fmt.Errorf("unmarshal file data: %v", err)
	}

	leaves := make([][]byte, len(data.Nodes))
	for i, node := range data.Nodes {
		category := strconv.FormatUint(node.Category, 10)
		amount := strconv.FormatUint(node.Amount, 10)
		leaves[i] = solsha3.SoliditySHA3(
			[]string{"address", "uint256", "uint256"},
			[]interface{}{node.Address, category, amount},
		)
	}

	tree := merkle.NewTree(leaves, merkle.Options{SortPairs: true})
	mergedArray := strings.Join(tree.GetHexProof(leaves[0]), ",")
	fmt.Printf("[%s]", mergedArray)

	return
}
