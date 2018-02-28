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

	assertEqual(t, 1, len(actual))

	first := actual[0]

	firstJSON, err := ToJSON(first)
	assertNil(t, err)

	assertEqual(t, firstJSON, `{"Left":{"Hashcode":"AAE=","Left":null,"Right":null},"Right":{"Hashcode":null,"Left":null,"Right":null}}`)

	actualJSON, err := ToJSON(actual)
	assertNil(t, err)

	expectedJSON := `[{"Left":{"Hashcode":"AAE=","Left":null,"Right":null},"Right":{"Hashcode":null,"Left":null,"Right":null}}]`

	assertEqual(t, expectedJSON, actualJSON )

}



