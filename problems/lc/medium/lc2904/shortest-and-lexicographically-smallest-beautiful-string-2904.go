package lc2904

func shortestBeautifulSubstring(s string, k int) string {
	n := len(s)
	left := 0

	ones := 0
	smallest := s

	for right := range n {
		ch := s[right]
		if ch == '1' {
			ones++
		}

		for ones > k || s[left] == '0' && left < right {
			if s[left] == '1' {
				ones--
			}
			left++
		}

		if ones == k {
			candidate := s[left : right+1]

			if len(candidate) < len(smallest) || len(candidate) == len(smallest) && candidate < smallest {
				smallest = candidate
			}
		}

	}

	if ones == k {
		return smallest
	} else {
		return ""
	}
}
