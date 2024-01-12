/* https://leetcode.com/problems/reverse-integer/
Accepted:
Runtime: 0ms
Memory: 2.16MB
*/

package reverseint

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	is_negative := x < 0
	x_str := ""
	if is_negative {
		x_str = strconv.Itoa(-x)
	} else {
		x_str = strconv.Itoa(x)
	}
	x_rev := reverseStr(x_str)
	x_rev_int, _ := strconv.Atoi(x_rev)
	if is_negative {
		x_rev_int = -x_rev_int
	}
	lower_bound := int(math.Pow(-2, 31))
	upper_bound := int(math.Pow(2, 31) - 1)
	if (x_rev_int > upper_bound) || (x_rev_int < lower_bound) {
		return 0
	} else {
		return x_rev_int
	}

}

func reverseStr(s string) string {
	rev_str := ""
	for _, c := range s {
		rev_str = string(c) + rev_str
	}

	return rev_str
}
