package medium

import "fmt"

func isVowel(r rune) bool {
	return r == 'a' || r == 'i' || r == 'u' || r == 'e' || r == 'o'
}

func maxVowels(s string, k int) int {
	left := 0
	right := 0
	max := 0
	count := 0
	queue := []int{}
	for right < len(s) {
		newChar := s[right]
		if isVowel(rune(newChar)) {
			queue = append(queue, right)
			count++
		}
		if len(queue) > 0 {
			front := queue[0]
			if front < left {
				queue = queue[1:]
				count--
			}
		}

		if count > max {
			max = count
		}
		if right-left+1 == k {
			left++
		}
		right++
	}
	return max
}

func maxVowelsClaudeSonnet(s string, k int) int {
	count := 0
	max := 0

	for i := 0; i < len(s); i++ {
		// Expand: add incoming character
		if isVowel(rune(s[i])) {
			count++
		}

		// Shrink: remove outgoing character once window is full
		if i >= k {
			if isVowel(rune(s[i-k])) {
				count--
			}
		}

		if count > max {
			max = count
		}
	}

	return max
}

func MaximumNumberOfVolwelsInASubstringOfGivenLength() {
	// s := "abciiidef"
	// k := 3

	// s := "aeiou"
	// k := 2

	// s := "leetcode"
	// k := 3

	s := "weallloveyou"
	k := 7

	out := maxVowels(s, k)
	fmt.Println(out)
}
