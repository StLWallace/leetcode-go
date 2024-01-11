/* https://leetcode.com/problems/search-insert-position/description/
Accepted:
Runtime: 4 ms
Memory Usage: 3 MB
*/

package searchinsertposition

func searchInsert(nums []int, target int) int {
	ind := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			ind = i
			break
		} else if nums[i] > target {
			ind = i
			break
		}
	}
	return ind
}
