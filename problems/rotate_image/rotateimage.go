/*
https://leetcode.com/problems/rotate-image/

Accepted:
Runtime: 1ms
Memory: 2.28MB
*/
package rotateimage

// Processes the rotation by starting with the outer values and iterates across the columns
// Continues to work inward until the next matrix is < 2x2
func rotate(matrix [][]int) [][]int {
	for i := 0; i < (len(matrix) - 1); i++ {
		opp_val := len(matrix) - i - 1
		for j := i; j < opp_val; j++ {
			rotateSet(matrix, i, j, opp_val)
		}
	}

	return matrix
}


// Gathers a set of 4 values from a "matrix" and reassigns them according to a 90 degree rotation
// Args:
// matrix - the original slices of slices
// i - the starting row index of the sub-matrix
// j - the starting column index of the sub-matrix
// opp_val - the ending column/row (square matrix) index
func rotateSet(matrix [][]int, i int, j int, opp_val int) {
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