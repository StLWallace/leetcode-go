/* https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
Accepted:
Runtime: 151 ms
Memory Usage: 7.3 MB
*/

package longestsubstring

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	longest := ""
	s_list := strings.Split(s, "")
	for i := 0; i < len(s_list); i++ {
		subs := s_list[i]
		for j := i + 1; j < len(s_list); j++ {
			if strings.Contains(subs, s_list[j]) {
				break
			} else {
				subs += s_list[j]
			}
		}
		if len(subs) > len(longest) {
			longest = subs
		}
	}
	return len(longest)
}
