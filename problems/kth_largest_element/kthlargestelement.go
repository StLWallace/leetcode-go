/* https://leetcode.com/problems/kth-largest-element-in-an-array/description/?envType=study-plan-v2&envId=leetcode-75
Not Accepted:
37/41 testcases passsed (Time Limit Exceeded)

My initial approach to this to avoid sorting is to keep a running list of the top-k values
- for the first k iterations through nums, I'll just add the value to the list in order
- after that, I'll check if it's larger than an existing value and displace the first if so
- finally, I'll return the first element of the top-k list (sorted ascending)

I read some solutions and I don't have much DSA experience, so I wasn't familiar
*/

package kthlargestelement

// I'm trying to solve this without sorting
func findKthLargest(nums []int, k int) int {
	top_k := []int{}
	for i, num := range nums {
		top_k = insertIfTop(top_k, int(num), k, i)
	}
    return top_k[0]
}


func insertIfTop(top_k []int, num int, k int, i int) []int {
	//n_elements := len(top_k)
	if k > i {
		top_k = addToList(top_k, num)
		return top_k
	} else {
			top_k = addIfLarger(top_k, num)
		return top_k
	}
}

func addIfLarger(top_k []int, num int) [] int {
	num_slice := []int{num}
	for i := 0; i < len(top_k); i++ {
		if num < top_k[i] {
			if i == 0 {
				return top_k
			} else {
				new_top_k := append(top_k[1:i], append(num_slice, top_k[(i):]...)...)
				return new_top_k
			}
		} 
	}
	new_top_k := append(top_k[1:], num_slice...)
	return new_top_k
}


// top_k will be sorted ascending
func addToList(top_k []int, num int) []int {
	for i := 0; i < len(top_k); i++ {
		if num < top_k[i] {
			num_slice := []int{num}
			new_top_k := append(top_k[:i], append(num_slice, top_k[(i):]...)...)
			return new_top_k
		}
	}
	// If no early return, it means num is largest element
	top_k = append(top_k, num)
	return top_k
}