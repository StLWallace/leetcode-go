/* https://leetcode.com/problems/container-with-most-water/description/
Accepted:
Runtime: 68ms
Memory: 8.89MB
I used the hints to recall the intuition on this "two pointer" solution
*/

package mostwater

import (
	"fmt"
)

func maxArea(height []int) int {
    l := 0
    r := len(height) -1
    max_area := 0
    for i := 0; l < r; i++ {
		fmt.Printf("l: %d, r: %d\n", l, r)
        this_area := getArea([]int{l+1, height[l]}, []int{r+1, height[r]})
        if this_area > max_area {
            max_area = this_area
		}
        if height[l] < height[r] {
            l += 1
        } else {
            r -= 1
        } 
        
    }
    return max_area
}


func getArea(p1 []int, p2 []int) int {
    area := 1
    length := p2[0] - p1[0]
    if p1[1] > p2[1] {
        area = p2[1]*length
    } else {
        area = p1[1]*length
    }

    return area
}