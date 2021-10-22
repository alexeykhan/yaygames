package merkle

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"golang.org/x/crypto/sha3"
)

type (
	// Tree defines merkle tree basic functions.
	Tree interface {
		GetProof(leaf []byte) [][]byte
		GetHexProof(leaf []byte) []string
	}

	// HashFunc defines hash function signature.
	HashFunc func([]byte) []byte
	// HashProvider is an identifier for hash function.
	HashProvider uint8

	// Options modify merkle tree settings.
	Options struct {
		SortPairs    bool
		HashProvider HashProvider
	}

	tree struct {
		leaves [][]byte   // initial nodes of a tree
		layers [][][]byte // levels of merkle tree

		hashFunc  HashFunc
		sortPairs bool
	}
)

const (
	// Hash function providers.
	Keccack256 HashProvider = iota
)

// Tree interface compliance verification.
var _ Tree = (*tree)(nil)

// NewTree builds a new merkle tree from given leaves.
func NewTree(leaves [][]byte, opts Options) Tree {
	t := &tree{
		leaves:    leaves,
		layers:    [][][]byte{leaves},
		sortPairs: opts.SortPairs,
		hashFunc:  getHashFunc(opts.HashProvider),
	}

	nodes := make([][]byte, len(t.leaves))
	copy(nodes, t.leaves)

	for len(nodes) > 1 {
		var nextLayer [][]byte
		for i := 0; i < len(nodes); i += 2 {
			if i+1 == len(nodes) && len(nodes)%2 == 1 {
				nextLayer = append(nextLayer, nodes[i])
				continue
			}

			left, right := nodes[i], nodes[i]
			if len(nodes) > i+1 {
				right = nodes[i+1]
			}

			if t.sortPairs {
				if bytes.Compare(right, left) < 0 {
					left, right = right, left
				}
			}

			data := append(left, right...)
			hash := t.hashFunc(data)

			nextLayer = append(nextLayer, hash)
		}

		t.layers = append(t.layers, nextLayer)
		nodes = nextLayer
	}

	return t
}

// GetProof returns merkle proof for a leaf in a raw format.
func (t *tree) GetProof(leaf []byte) (proof [][]byte) {
	if len(leaf) == 0 {
		return
	}

	index := -1
	for i := 0; i < len(t.leaves); i++ {
		if bytes.Equal(leaf, t.leaves[i]) {
			index = i
		}
	}

	if index < 0 {
		return
	}

	for i := 0; i < len(t.layers); i++ {
		layer := t.layers[i]
		isRightNode := index%2 == 1
		pairIndex := index + 1
		if isRightNode {
			pairIndex = index - 1
		}

		if pairIndex < len(layer) {
			proof = append(proof, layer[pairIndex])
		}

		index /= 2
	}

	return
}

// GetHexProof returns merkle proof for a leaf in a hex format.
func (t *tree) GetHexProof(leaf []byte) []string {
	dataProof := t.GetProof(leaf)

	proof := make([]string, len(dataProof))
	for i, p := range dataProof {
		proof[i] = hexutil.Encode(p)
	}

	return proof
}

// getHashFunc returns a hash function by key. Keccack256 is returned by default.
func getHashFunc(provider HashProvider) HashFunc {
	switch provider {
	// case ...
	case Keccack256:
		fallthrough
	default:
		return func(data []byte) []byte {
			hash := sha3.NewLegacyKeccak256()
			hash.Write(data)
			return hash.Sum(nil)
		}
	}
}
