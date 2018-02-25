package merkletree

import (
	"testing"
	"fmt"
)

func assertEqual(t *testing.T, obj0 interface{}, obj1 interface{}, message string) {

	// TODO: the limitation here for now is that we convert each object to
	// a string which is a bit of a hack for now. We should use a library that
	// has support for a better assertEquals function.

	str0 := fmt.Sprintf("%+v", obj0)
	str1 := fmt.Sprintf("%+v", obj1)

	if str0 == str1 {
		return
	}

	if len(message) == 0 {
		message = fmt.Sprintf("%s != %s", str0, str1)
	}

	t.Fatal(message)

}

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
