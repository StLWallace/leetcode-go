/* https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
Not Accepted Yet
*/
package lettercombo

import (
	"gonum.org/v1/gonum/stat/combin"
)


func letterCombinations(digits string) []string {

	letters_slice := makeLettersSlice(digits)
	combos := combin.Combinations(len(letters_slice), 2)
	result := []string{}
	for _, c := range combos {
		combo := letters_slice[c[0]] + letters_slice[c[1]]
		result = append(result, combo)
	}

	return result
}

func makeLettersSlice(digits string) []string {
    letters := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	letters_slice := []string{}
	for _, c := range digits {
		these_letters := letters[string(c)]
		letters_slice = append(letters_slice, these_letters...)
	}

	return letters_slice
}