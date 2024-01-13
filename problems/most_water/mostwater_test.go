package mostwater

import (
	"testing"
)

func TestGetArea(t *testing.T) {
	test_p1 := []int{1, 1}
	test_p2 := []int{2, 2}

	wanted := 1

	test_result := getArea(test_p1, test_p2)

	if test_result != wanted {
		t.Fatalf("test_result: %d. Should be %d", test_result, wanted)
	}
}

func TestMaxArea(t *testing.T) {
	test_height := []int{2,3,4,5,18,17,6}
	wanted := 17

	test_result := maxArea(test_height)

	if test_result != wanted {
		t.Fatalf("test_result: %d. Should be %d", test_result, wanted)
	}

}