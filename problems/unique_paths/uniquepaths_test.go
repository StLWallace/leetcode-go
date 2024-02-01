package uniquepaths

import "testing"


func TestUniquePaths(t *testing.T) {
	test_m := 3
	test_n := 7
	test_result := uniquePaths(test_m, test_n)
	wanted := 28

	if test_result != wanted {
		t.Fatalf("test_result: %d, wanted: %d", test_result, wanted)
	}

}