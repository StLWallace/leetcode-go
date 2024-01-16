package rotateimage

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

func matrixEqual(matrix1 [][]int, matrix2 [][]int) bool {
	if len(matrix1) != len(matrix2) {
		return false
	} else {
		for i := 0; i < len(matrix1); i++ {
			if !slicesEqual(matrix1[i], matrix2[i]) {
				return false
			}
		}
	}
	return true
}

func TestRotate1(t *testing.T) {
	test_matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	test_result := rotate(test_matrix)
	wanted := [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}

	if !matrixEqual(test_result, wanted) {
		t.Fatalf("test_result: %v, wanted: %v", test_result, wanted)
	}
}

func TestRotate2(t *testing.T) {
	test_matrix := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	test_result := rotate(test_matrix)
	wanted := [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}}

	if !matrixEqual(test_result, wanted) {
		t.Fatalf("test_result: %v, wanted: %v", test_result, wanted)
	}
}
