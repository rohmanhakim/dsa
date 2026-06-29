package lc14

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".



Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"

Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.



Constraints:

    1 <= strs.length <= 200
    0 <= strs[i].length <= 200
    strs[i] consists of only lowercase English letters if it is non-empty.


*/

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	pref := strs[0]

	for i := 1; i < len(strs); i++ {
		curr := strs[i]
		sublen := len(pref)

		if len(curr) < sublen {
			sublen = len(curr)
		}
		preflen := 0

		for j := 0; j < sublen; j++ {
			if pref[j] == curr[j] {
				preflen++
			} else {
				break
			}
		}

		pref = pref[:preflen]

	}

	return pref
}

func LongestCommonPrefix() {
	strs := []string{"flower", "flow", "flight"}
	out := longestCommonPrefix(strs)
	fmt.Println(out)
}
