/*
https://leetcode.com/problems/rotate-image/

Accepted:
Runtime: 1ms
Memory: 2.26MB
*/
package rotateimage

func rotate(matrix [][]int) [][]int {
	for i := 0; i < (len(matrix) - 1); i++ {
		opp_val := len(matrix) - i - 1
		for j := i; j < opp_val; j++ {
			rot_vals := []int{
				matrix[i][j],
				matrix[j][opp_val],
				matrix[opp_val][opp_val-j+i],
				matrix[opp_val-j+i][i],
			}
			matrix[j][opp_val] = rot_vals[0]
			matrix[opp_val][opp_val-j+i] = rot_vals[1]
			matrix[opp_val-j+i][i] = rot_vals[2]
			matrix[i][j] = rot_vals[3]
		}
	}

	return matrix
}
