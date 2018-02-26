package merkletree

import (
	"testing"
)

func TestComputeHashNodePairsWithNoInputNodes(t *testing.T) {

	input := make([]HashNode, 0)

	output := computeHashNodePairs(input)

	assertEqual(t, output, input)

}
