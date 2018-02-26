package merkletree

import (
	"testing"
)

func TestComputeHashNodePairsWithNoInputNodes(t *testing.T) {

	input := make([]HashNode, 0)

	output := computeHashNodePairs(input)

	assertEqual(t, output, input)

}

func TestComputeHashNodePairsWithOneInputNodes(t *testing.T) {

	input := make([]HashNode, 0)
	input = append(input, HashNode{[]byte{0x00, 0x01}, nil, nil })

	output := computeHashNodePairs(input)

	assertEqual(t, output, input)

}
