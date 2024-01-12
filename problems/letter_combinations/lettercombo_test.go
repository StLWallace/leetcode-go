package lettercombo

import (
	"testing"
)

func slicesEqual(a, b []string) bool {
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

func TestLetterCombinations(t *testing.T) {
	test_digits := "23"
	test_result := letterCombinations(test_digits)
	wanted := []string{"ad","ae","af","bd","be","bf","cd","ce","cf"}

	if !slicesEqual(test_result, wanted){
		t.Fatalf("test_results: %v, should be %v", test_result, wanted)
	}

}