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

	assertTrue(t, first.IsPartial())

	firstJSON, err := ToJSON(first)
	assertNil(t, err)

	assertEqual(t, firstJSON, `{"Left":{"Hashcode":"AAE=","Left":null,"Right":null},"Right":{"Hashcode":null,"Left":null,"Right":null}}`)

	actualJSON, err := ToJSON(actual)
	assertNil(t, err)

	expectedJSON := `[{"Left":{"Hashcode":"AAE=","Left":null,"Right":null},"Right":{"Hashcode":null,"Left":null,"Right":null}}]`

	assertEqual(t, expectedJSON, actualJSON )

}


func TestComputeHashNodePairsWithTwoInputNodes(t *testing.T) {

	input := make([]HashNode, 0)
	input = append(input, HashNode{[]byte{0x00, 0x01}, nil, nil })
	input = append(input, HashNode{[]byte{0x00, 0x02}, nil, nil })

	actual := computeHashNodePairs(input)

	assertEqual(t, 1, len(actual))
	first := actual[0]

	assertFalse(t, first.IsPartial())

	firstJSON, err := ToJSON(first)
	assertNil(t, err)

	assertEqual(t, firstJSON, `{"Left":{"Hashcode":"AAE=","Left":null,"Right":null},"Right":{"Hashcode":"AAI=","Left":null,"Right":null}}`)

	actualJSON, err := ToJSON(actual)
	assertNil(t, err)

	expectedJSON := `[{"Left":{"Hashcode":"AAE=","Left":null,"Right":null},"Right":{"Hashcode":"AAI=","Left":null,"Right":null}}]`

	assertEqual(t, expectedJSON, actualJSON )

}



func TestMergeForIntermediateTwoInputNodes(t *testing.T) {

	input := make([]HashNode, 0)
	input = append(input, HashNode{[]byte{0x00, 0x01}, nil, nil })
	input = append(input, HashNode{[]byte{0x00, 0x02}, nil, nil })

	merged := mergeForIntermediate(input)

	assertEqual(t, 1, len(merged))

	actualJSON, err := ToJSON(merged)
	assertNil(t, err)

	expectedJSON := `[{"Hashcode":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACzsZgS3VHvBE9+ABb/GaY00YEtBRWVqlVCbcbpGzEOkw==","Left":{"Hashcode":"AAE=","Left":null,"Right":null},"Right":{"Hashcode":"AAI=","Left":null,"Right":null}}]`

	assertEqual(t, actualJSON, expectedJSON )

}


func TestMergeForIntermediateFourInputNodes(t *testing.T) {

	input := make([]HashNode, 0)
	input = append(input, HashNode{[]byte{0x00, 0x01}, nil, nil })
	input = append(input, HashNode{[]byte{0x00, 0x02}, nil, nil })
	input = append(input, HashNode{[]byte{0x00, 0x03}, nil, nil })
	input = append(input, HashNode{[]byte{0x00, 0x04}, nil, nil })

	merged := mergeForIntermediate(input)

	assertEqual(t, 1, len(merged))

	actualJSON, err := ToJSON(merged)
	assertNil(t, err)

	expectedJSON := `[{"Hashcode":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAdfjQwkxemB4Jf8/DWPYJ6ZMO1RUlI3r3yTmNIPJu/NA==","Left":{"Hashcode":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACzsZgS3VHvBE9+ABb/GaY00YEtBRWVqlVCbcbpGzEOkw==","Left":{"Hashcode":"AAM=","Left":null,"Right":null},"Right":{"Hashcode":"AAQ=","Left":null,"Right":null}},"Right":{"Hashcode":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABfQa4+bVgavV/MUA9orVs9LvRv+H5Sb/6cGIrti2VIOQ==","Left":{"Hashcode":"AAM=","Left":null,"Right":null},"Right":{"Hashcode":"AAQ=","Left":null,"Right":null}}}]`

	assertEqual(t, actualJSON, expectedJSON )

}







