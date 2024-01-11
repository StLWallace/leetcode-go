package twosum

import (
	"testing"
)

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestTwoSum(t *testing.T) {
	test_nums := [4]int{2,7,11,15}
	test_target := 9
	want := []int{0, 1}

	test_answer := TwoSum(test_nums[:], test_target)

	are_equal := slicesEqual(want, test_answer)
    if !(are_equal) {
        t.Fatalf("test_answer was %v, should be %v", test_answer, want)
    }
}