package validsudoku

import (
	"testing"
)


func TestGetColumnArrays(t *testing.T) {
	test_board := [][]byte{{1, 2, 3},}
	test_result := getColumnArrays(test_board)
	t.Fatalf("test_result: %v", test_result)
}

func TestIsValidSudoku(t *testing.T) {
	test_board := [][]byte{{1, 2, 3},}
	test_result := isValidSudoku(test_board)
	t.Fatalf("test_result: %v", test_result)
}
