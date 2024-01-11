package main

import (
	"fmt"

	"github.com/StLWallace/leetcode-go/problems/twosum"
)

func main() {
	test_list := [4]int{2,7,11,15}
	test_target := 9

	test_answer := twosum.TwoSum(test_list[:], test_target)

	fmt.Println(test_answer) 
}