package merkletree

import (
	"testing"
	"fmt"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func TestComputeSlabReferences(t *testing.T) {

	fmt.Printf("Test worked!\n")
	t.Log("worked fine")

	references := ComputeSlabReferences(10, 2)

	assertEqual(t, references, "", "")

}
