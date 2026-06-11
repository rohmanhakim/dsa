/*
## The Core Question: What Are We Really Solving?

"Minimum removals" = "Maximum intervals we can **keep**" that are non-overlapping.

This is the classic **Activity Selection / Interval Scheduling** problem. And the greedy proof is:

> **At each step, pick the interval that finishes EARLIEST.** Why? Because it "uses up" the least amount of time on the timeline, leaving maximum room for future intervals.

## The Key Insight: Sort by END Time, Not Start Time

Think of it this way — what matters for future conflicts is **when an interval ends**, not when it starts. An interval that starts early but ends late is more "greedy" and blocks more future intervals.

**The correct algorithm:**
1. Sort by **end time** ascending
2. Walk through, always keeping the interval if it doesn't conflict with the last kept one
3. If it conflicts → remove it (the one we already kept ends earlier, so it's the better choice)

## Tracing Test 6 With the Correct Approach

Sorted by **end time**:
```
[-62,-49] [-73,-26] [-40,-26] [-65,-11] [-63,2] [-52,31] [30,47] [-31,49] [58,95] [82,97] [66,98] [95,99]
```

Walk through:
- Keep `[-62,-49]`. prevEnd = -49.
- `[-73,-26]`: start -73 < -49 → overlap! **Remove.** (total=1)
- `[-40,-26]`: start -40 ≥ -49 → **Keep.** prevEnd = -26.
- `[-65,-11]`: start -65 < -26 → overlap! **Remove.** (total=2)
- `[-63, 2]`: start -63 < -26 → overlap! **Remove.** (total=3)
- `[-52, 31]`: start -52 < -26 → overlap! **Remove.** (total=4)
- `[30, 47]`: start 30 ≥ -26 → **Keep.** prevEnd = 47.
- `[-31, 49]`: start -31 < 47 → overlap! **Remove.** (total=5)
- `[58, 95]`: start 58 ≥ 47 → **Keep.** prevEnd = 95.
- `[82, 97]`: start 82 < 95 → overlap! **Remove.** (total=6)
- `[66, 98]`: start 66 < 95 → overlap! **Remove.** (total=7)
- `[95, 99]`: start 95 ≥ 95 → **Keep.** prevEnd = 99.

**Result: 7 removals** — we kept 5 intervals: `[-62,-49], [-40,-26], [30,47], [58,95], [95,99]`.
*/

package medium

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	total := 1
	prevIdx := 0
	for i := 1; i < len(intervals); i++ {
		start := intervals[i][0]
		prevEnd := intervals[prevIdx][1]

		if start < prevEnd {
			continue
		}
		total++
		prevIdx = i

	}

	return len(intervals) - total
}

func NonOverlappingIntervals() {
	intervals := [][]int{
		{-52, 31}, {-73, -26}, {82, 97}, {-65, -11}, {-62, -49}, {95, 99}, {58, 95}, {-31, 49}, {66, 98}, {-63, 2}, {30, 47}, {-40, -26},
	} // 7
	out := eraseOverlapIntervals(intervals)
	fmt.Println(out)
}
