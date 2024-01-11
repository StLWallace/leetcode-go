/* https://leetcode.com/problems/roman-to-integer/description/
Accepted:
Runtime: 3 ms
Memory Usage: 2.8 MB
*/

package romantoint

import (
	"strings"
)

func romanToInt(s string) int {
	thousands := 1000 * strings.Count(s, "M")
	five_hundreds := 500 * strings.Count(s, "D")
	hundreds := handleC(s)
	fifties := 50 * strings.Count(s, "L")
	tens := handleX(s)
	fives := 5 * strings.Count(s, "V")
	ones := handleI(s)

	total := thousands + five_hundreds + hundreds + fifties + tens + fives + ones
	return total
}

func handleI(s string) int {
	if !strings.Contains(s, "I") {
		return 0
	} else {
		num := strings.Count(s, "I")
		if strings.Contains(s, "IV") {
			num -= 2
		} else if strings.Contains(s, "IX") {
			num -= 2
		}
		return num
	}
}

func handleX(s string) int {
	if !strings.Contains(s, "X") {
		return 0
	} else {
		num := 10 * strings.Count(s, "X")
		if strings.Contains(s, "XL") {
			num -= 20
		} else if strings.Contains(s, "XC") {
			num -= 20
		}
		return num
	}
}

func handleC(s string) int {
	if !strings.Contains(s, "C") {
		return 0
	} else {
		num := 100 * strings.Count(s, "C")
		if strings.Contains(s, "CD") {
			num -= 200
		} else if strings.Contains(s, "CM") {
			num -= 200
		}
		return num
	}
}
