package merkletree

import (
	"testing"
	"fmt"
)

func assertTrue(t *testing.T, expected bool) {
	if ! expected {
		t.Fatal("Value is not true")
	}
}

func assertNil(t *testing.T, expected interface{}) {
	if expected != nil {
		t.Fatal("Value is not nil")
	}
}

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {

	// TODO: the limitation here for now is that we convert each object to
	// a string which is a bit of a hack for now. We should use a library that
	// has support for a better assertEquals function.

	strExpected := format(expected)
	strActual := format(actual)

	if strExpected == strActual {
		return
	}

	message := fmt.Sprintf("Assert equals failed:\n expected: %s\n actual:   %s", strExpected, strActual)

	t.Fatal(message)

}

// Format the input to avoid double encoding strings but gracefully handle other types.
func format(value interface{}) string {

	switch v := value.(type) {
	case string:
		return v
	default:
		return fmt.Sprintf("%#v", value)
	}

}


