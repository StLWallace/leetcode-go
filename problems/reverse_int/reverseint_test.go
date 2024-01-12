package reverseint

import (
	"testing"
)

func TestReverseStr(t *testing.T) {
	test_s := "321"
	wanted := "123"

	test_rev := reverseStr(test_s)

	if test_rev != wanted {
		t.Fatalf("Output val: %s. Should be: %s", test_rev, wanted)
	}
}

func TestReverse(t *testing.T) {
	test_int := -123
	wanted := -321

	test_output := reverse(test_int)

	if test_output != wanted {
		t.Fatalf("Output val: %d. Should be: %d", test_output, wanted)
	}
}
