package lc1358

// numberOfSubstringsMentor - Alternative sliding window approach
// Instead of counting n-right extensions, this counts left+1 substrings
// ending at each right position.
//
// Key insight: when window [left, right] is valid, all substrings
// s[0:right+1], s[1:right+1], ..., s[left:right+1] are valid.
// That's left+1 substrings ending at right.
func numberOfSubstringsMentor(s string) int {
	count := [3]int{} // count[0]=a, count[1]=b, count[2]=c
	left := 0
	result := 0

	for right := 0; right < len(s); right++ {
		count[s[right]-'a']++

		// Shrink window until it's no longer valid
		for count[0] > 0 && count[1] > 0 && count[2] > 0 {
			count[s[left]-'a']--
			left++
		}
		// After shrinking, left is the first invalid position
		// So all substrings s[k:right+1] for k=0..left-1 are valid = left substrings
		result += left
	}
	return result
}
