package lc3471

func largestInteger(nums []int, k int) int {
	freq := make(map[int]int)

	left := 0
	sub := 1
	for right := 0; right < len(nums); right++ {
		freq[nums[right]]++
		wLen := right - left + 1
		if wLen > k {
			sub++
			left++
			for j := left; j < right; j++ {
				freq[nums[j]]++
			}
		}
	}
	largest := -1

	if sub == 1 {
		for k := range freq {
			if k > largest {
				largest = k
			}
		}
	} else {
		for k, v := range freq {
			if v == 1 {
				if k > largest {
					largest = k
				}
			}
		}
	}

	return largest
}
