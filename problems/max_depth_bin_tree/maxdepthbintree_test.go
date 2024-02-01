package maxdepthbintree

import "testing"


func TestMaxDepth(t *testing.T) {
	test_l2 := TreeNode{15, nil, nil}
	test_r2 := TreeNode{7, nil, nil}
	test_l1 := TreeNode{9, nil, nil}
	test_r1 := TreeNode{20, &test_l2, &test_r2}
	test_root := TreeNode{3, &test_r1, &test_l1}
	wanted := 3
	test_result := maxDepth(&test_root)

	if wanted != test_result {
		t.Fatalf("wanted: %d, test_result: %d", wanted, test_result)
	}

}