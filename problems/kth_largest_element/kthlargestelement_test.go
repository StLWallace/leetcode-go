package kthlargestelement

import "testing"

func slicesEqual(s1 []int, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	} else {
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				return false
			}
		}
	}
	return true
}

func TestFindKthLargest(t *testing.T) {
	test_nums := []int{1,2,3,4}
	test_k := 2
	wanted := 3
	test_result := findKthLargest(test_nums, test_k)

	if wanted != test_result {
		t.Fatalf("wanted: %d, test_result: %d", wanted, test_result)
	}
}

func TestInsertIfTop(t *testing.T){
	test_top_k := []int{1,2,3,5}
	test_num := 4
	test_k := 4
	test_i := 4
	wanted := []int{2,3,4,5}
	test_result := insertIfTop(test_top_k, test_num, test_k, test_i)

	if !slicesEqual(test_result, wanted) {
		t.Fatalf("wanted: %v, test_result: %v", wanted, test_result)
	}

}

func TestInsertIfTop2(t *testing.T){
	test_top_k := []int{1,2}
	test_num := 4
	test_k := 2
	test_i := 2
	wanted := []int{2,4}
	test_result := insertIfTop(test_top_k, test_num, test_k, test_i)

	if !slicesEqual(test_result, wanted) {
		t.Fatalf("wanted: %v, test_result: %v", wanted, test_result)
	}

}

func TestAddIfLarger(t *testing.T) {
	test_top_k := []int{1,2,3,5}
	test_num := 4
	wanted := []int{2,3,4,5}
	test_result := addIfLarger(test_top_k, test_num)

	if !slicesEqual(test_result, wanted) {
		t.Fatalf("wanted: %v, test_result: %v", wanted, test_result)
	}

}

func TestAddIfLargerEnd(t *testing.T) {
	test_top_k := []int{1,2,3,5}
	test_num := 6
	wanted := []int{2,3,5,6}
	test_result := addIfLarger(test_top_k, test_num)

	if !slicesEqual(test_result, wanted) {
		t.Fatalf("wanted: %v, test_result: %v", wanted, test_result)
	}

}

func TestAddToList(t *testing.T) {
	test_top_k := []int{1,2,4}
	test_num := 3
	wanted := []int{1,2,3,4}
	test_result := addToList(test_top_k, test_num)
	if !slicesEqual(wanted, test_result) {
		t.Fatalf("wanted: %v, test_result: %v", wanted, test_result)
	}

}

func TestAddToListEnd(t *testing.T) {
	test_top_k := []int{1,2,4}
	test_num := 5
	wanted := []int{1,2,4,5}
	test_result := addToList(test_top_k, test_num)
	if !slicesEqual(wanted, test_result) {
		t.Fatalf("wanted: %v, test_result: %v", wanted, test_result)
	}

}

func TestAddToListStart(t *testing.T) {
	test_top_k := []int{1,2,4}
	test_num := 0
	wanted := []int{0,1,2,4}
	test_result := addToList(test_top_k, test_num)
	if !slicesEqual(wanted, test_result) {
		t.Fatalf("wanted: %v, test_result: %v", wanted, test_result)
	}

}