package lc1871

import (
	"fmt"
)

func canReach2BFS(s string, minJump int, maxJump int) bool {
	zeros := []int{0}
	left := 0
	for len(zeros) > 0 {
		left = zeros[0]
		zeros = zeros[1:]
		for j := left + minJump; j <= min(left+maxJump, len(s)-1); j++ {
			if rune(s[j]) == '0' {
				if j == len(s)-1 {
					return true
				}
				zeros = append(zeros, j)
			}
		}
	}
	return false
}

func canReach2ClaudeSonnet(s string, minJump int, maxJump int) bool {
	n := len(s)
	if s[n-1] == '1' {
		return false
	}

	reachable := make([]bool, n)
	reachable[0] = true
	windowCount := 0 // # of reachable indices in sliding window [j-maxJump, j-minJump]

	for j := 1; j < n; j++ {
		// Expand window right edge: index (j - minJump) just entered the window
		lo := j - minJump
		if lo >= 0 && reachable[lo] {
			fmt.Printf("wc++ ")
			windowCount++
		}

		// Shrink window left edge: index (j - maxJump - 1) just left the window
		expired := j - maxJump - 1
		if expired >= 0 && reachable[expired] {
			fmt.Printf("wc-- ")
			windowCount--
		}

		if s[j] == '0' && windowCount > 0 {
			reachable[j] = true
		}

		fmt.Printf("j: %d, lo: %d, exp: %d, wc: %d, s[j]: %c, reachable[%d]: %t\n", j, lo, expired, windowCount, s[j], j, reachable[j])
	}

	// fmt.Println(reachable)
	return reachable[n-1]
}

func canReach2(s string, minJump int, maxJump int) bool {
	n := len(s)
	canJump := make([]bool, n)
	canJump[0] = true

	windowSize := 0
	for j := 1; j < n; j++ {

		right := j - minJump
		if right >= 0 && canJump[right] {
			fmt.Printf("canjump[right]: %t ", canJump[right])
			windowSize++
		}
		fmt.Printf("j: %d, wc: %d\n", j, windowSize)

		left := j - maxJump - 1
		if left >= 0 && canJump[left] {
			windowSize--
		}

		if s[j] == '0' && windowSize > 0 {
			canJump[j] = true
		}
	}

	return canJump[n-1]
}

func JumpGameVII() {
	// s := "011010"
	// minJump := 2
	// maxJump := 3

	// s := "01101110"
	// minJump := 2
	// maxJump := 3

	s := "011111000111000001011111010"
	minJump := 6
	maxJump := 8

	out := canReach2(s, minJump, maxJump)
	fmt.Println(out)

	// fmt.Println(getJumpChoices(s, 3, minJump, maxJump))
}
