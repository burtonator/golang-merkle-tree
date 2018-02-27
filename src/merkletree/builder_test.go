package merkletree

import (
	"testing"
	"fmt"
)

//
//import (
//	"testing"
//)
//
//func TestComputeHashNodePairsWithNoInputNodes(t *testing.T) {
//
//	input := make([]HashNode, 0)
//
//	actual := computeHashNodePairs(input)
//
//	expected := make([]HashNodePair, 0)
//
//	assertEqual(t, actual, expected)
//
//}
//

func TestComputeHashNodePairsWithOneInputNodes(t *testing.T) {

	input := make([]HashNode, 0)
	input = append(input, HashNode{[]byte{0x00, 0x01}, nil, nil })

	actual := computeHashNodePairs(input)

	assertEqual(t, 1, len(actual))

	// FIXME: my ToJSON method isn't working properly.

	fmt.Printf("FIXME2 actual: %+v\n", actual)

	actualJSON, err := ToJSON(actual)
	assertNil(t, err)
	fmt.Printf("FIXME3: %+v\n", actualJSON)


	expected := make([]HashNodePair, 0)


	assertEqual(t, expected, actualJSON )

}

