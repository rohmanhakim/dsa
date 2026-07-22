package lc3501

import (
	"math/bits"
	"slices"
	"sort"
)

func maxActiveSectionsAfterTrade(s string, queries [][]int) []int {

	n := len(s)
	zeros := 0
	zeroBlocks := []int{}
	blockLeft := []int{}
	blockRight := []int{}
	countOne := 0

	for i, ch := range s {
		if ch == '0' {
			zeros++
			if i == 0 || (i > 0 && s[i-1] == '1') {
				blockLeft = append(blockLeft, i)
			}
			if i == n-1 || (i < n-1 && s[i+1] == '1') {
				blockRight = append(blockRight, i)
			}
		} else {
			countOne++
		}
		if i == n-1 || ch == '1' {
			if zeros != 0 {
				zeroBlocks = append(zeroBlocks, zeros)
				zeros = 0
			}
		}
	}
	if len(s) == 1 {
		return []int{countOne}
	}
	m := len(zeroBlocks)

	pow2 := func(x int) int {
		return 1 << x
	}

	makeMaxSt := func(tmpSum []int) [][]int {
		maxst := [][]int{}
		maxst = append(maxst, tmpSum)
		j := 1
		for j < bits.Len(uint(len(tmpSum))) {
			maxst = append(maxst, make([]int, len(tmpSum)-(pow2(j)-1)))
			for i := 0; i < len(maxst[j]); i++ {
				maxst[j][i] = max(maxst[j-1][i], maxst[j-1][i+pow2(j-1)])
			}
			j++
		}
		return maxst
	}

	var tmpSum []int
	if m > 0 {
		tmpSum = make([]int, m-1)
		for i := 0; i < m-1; i++ {
			tmpSum[i] = zeroBlocks[i] + zeroBlocks[i+1]
		}
	}

	var maxSt [][]int
	if len(tmpSum) > 0 {
		maxSt = makeMaxSt(tmpSum)
	}

	ans := make([]int, len(queries))

	for qi, q := range queries {
		if len(zeroBlocks) < 2 {
			ans[qi] = countOne
			continue
		}

		l := q[0]
		r := q[1]
		bestGain := 0
		vals := []int{}

		// Step 1: find first block overlapping [l, r]
		ridx := sort.Search(len(blockRight), func(i int) bool {
			return blockRight[i] >= l
		})
		if ridx >= len(blockRight) {
			ans[qi] = countOne
			continue
		}
		li := blockLeft[ridx]
		ri := blockRight[ridx]
		zzi := ri - max(li, l) + 1

		// Step 2: find last block overlapping [l, r]
		lidx := sort.Search(len(blockLeft), func(j int) bool {
			return blockLeft[j] > r
		}) - 1
		if lidx < 0 || lidx >= len(blockLeft) {
			ans[qi] = countOne
			continue
		}
		lj := blockLeft[lidx]
		rj := blockRight[lidx]
		zzj := min(rj, r) - lj + 1

		// Step 3: at most one block → no gain
		if ridx >= lidx {
			ans[qi] = countOne
			continue
		}

		// Step 4: compute val1 and val2
		// val1 = first block truncated + next block
		val1 := zzi + zeroBlocks[ridx+1]
		if ridx+1 == lidx {
			// next block IS the last block — use truncated length
			val1 = zzi + zzj
		}
		vals = append(vals, val1)

		// val2 = prev block + last block truncated
		val2 := zeroBlocks[lidx-1] + zzj
		if lidx-1 == ridx {
			// prev block IS the first block — use truncated length
			val2 = zzi + zzj
		}
		vals = append(vals, val2)

		// Step 5: val3 = interior pairs (exclude boundary pairs covered by val1/val2)
		if lidx-1 > ridx+1 {
			left := ridx + 1
			right := lidx - 2
			k := bits.Len(uint(right-left+1)) - 1
			val3 := max(maxSt[k][left], maxSt[k][right-(1<<k)+1])
			vals = append(vals, val3)
		}
		bestGain = slices.Max(vals)

		ans[qi] = bestGain + countOne

	}

	return ans
}
