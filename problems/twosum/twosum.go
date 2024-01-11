/*
https://leetcode.com/problems/two-sum/description/

Accepted:
Runtime: 35 ms
Memory Usage: 3.6 MB
*/
package twosum

import (
	"fmt"
)

func TwoSum(nums []int, target int) []int {
	sol := []int{}
	for i := 0; i < len(nums); i++ {
		fmt.Printf("i = %d\n", i)
		for j := (i + 1); j < len(nums); j++ {
			if i != j {
				if (nums[i] + nums[j]) == target {
					sol = append(sol, i)
					sol = append(sol, j)
					return sol
				}
			}
		}
	}
	return sol
}
