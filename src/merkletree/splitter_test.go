package merkletree

import (
	"testing"
)

func TestComputeSlabReferencesWithExactBounds(t *testing.T) {
	
	actual := ComputeSlabReferences(10, 10)
	expected := []SlabReference{{0,10}}

	assertEqual(t, actual, expected, "")

}

func TestComputeSlabReferencesWithTruncatedbounds(t *testing.T) {

	actual := ComputeSlabReferences(12, 10)
	expected := []SlabReference{{0,10},{10,2}}

	assertEqual(t, actual, expected, "")

}
