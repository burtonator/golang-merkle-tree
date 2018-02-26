package merkletree

import (
	"testing"
)

func TestComputeHashNodePairsWithNoInputNodes(t *testing.T) {

	input := make([]HashNode, 0)

	actual := computeHashNodePairs(input)

	expected := make([]HashNodePair, 0)

	assertEqual(t, actual, expected)

}

func TestComputeHashNodePairsWithOneInputNodes(t *testing.T) {

	input := make([]HashNode, 0)
	input = append(input, HashNode{[]byte{0x00, 0x01}, nil, nil })

	actual := computeHashNodePairs(input)

	assertEqual(t, actual, input)

}
