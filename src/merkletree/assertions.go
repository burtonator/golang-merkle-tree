package merkletree

import (
	"testing"
	"fmt"
)

func assertEqual(t *testing.T, obj0 interface{}, obj1 interface{}) {

	// TODO: the limitation here for now is that we convert each object to
	// a string which is a bit of a hack for now. We should use a library that
	// has support for a better assertEquals function.

	str0 := fmt.Sprintf("%+v", obj0)
	str1 := fmt.Sprintf("%+v", obj1)

	if str0 == str1 {
		return
	}

	message := fmt.Sprintf("%s != %s", str0, str1)

	t.Fatal(message)

}

