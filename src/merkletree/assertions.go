package merkletree

import (
	"testing"
	"fmt"
)

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {

	// TODO: the limitation here for now is that we convert each object to
	// a string which is a bit of a hack for now. We should use a library that
	// has support for a better assertEquals function.

	strExpected := fmt.Sprintf("%+v", expected)
	strActual := fmt.Sprintf("%+v", actual)

	if strExpected == strActual {
		return
	}

	message := fmt.Sprintf("Assert equals failed:\n expected: %s\n actual:   %s", strExpected, strActual)

	t.Fatal(message)

}

