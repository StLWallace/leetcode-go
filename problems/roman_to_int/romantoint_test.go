package romantoint

import (
	"testing"
)

func TestHandleI(t *testing.T) {
	test_val := "LIX"
	wanted := -1

	test_int := handleI(test_val)

	if test_int != wanted {
		t.Fatalf("test_int was %d, should be %d", test_int, wanted)
	}

}


func TestRomanToInt(t *testing.T) {
	test_val := "III"
	wanted := 3

	test_int := romanToInt(test_val)

	if test_int != wanted {
		t.Fatalf("test_int was %d, should be %d", test_int, wanted)
	}

}