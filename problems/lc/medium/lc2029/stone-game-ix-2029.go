package lc2029

func stoneGameIX(stones []int) bool {
	cnt := [3]int{0, 0, 0}
	for _, s := range stones {
		cnt[s%3]++
	}

	if cnt[0]%2 == 0 {
		return cnt[1] > 0 && cnt[2] > 0
	} else {
		diff := cnt[1] - cnt[2]
		if diff < 0 {
			diff = -diff
		}
		return diff > 2
	}
}
