package lc2058

import (
	"math"

	"github.com/rohmanhakim/leetcode/lib"
)

func nodesBetweenCriticalPoints(head *lib.ListNode) []int {
	var prev *lib.ListNode
	curr := head

	idx := 0
	firstCriticalIndex := -1
	prevCriticalIndex := -1
	minDist := math.MaxInt
	maxDist := -1
	for curr != nil {
		if prev != nil && curr.Next != nil {
			if (curr.Val > prev.Val && curr.Val > curr.Next.Val) || (curr.Val < prev.Val && curr.Val < curr.Next.Val) {
				if firstCriticalIndex == -1 {
					firstCriticalIndex = idx
					prevCriticalIndex = idx
				} else {
					minDist = min(minDist, idx-prevCriticalIndex)
					prevCriticalIndex = idx
					maxDist = idx - firstCriticalIndex
				}
			}
		}
		prev = curr
		curr = curr.Next
		idx++
	}

	if maxDist == -1 {
		return []int{-1, -1}
	}

	return []int{minDist, maxDist}
}
