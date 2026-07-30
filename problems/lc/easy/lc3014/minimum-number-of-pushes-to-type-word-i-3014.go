package lc3014

func minimumPushes(word string) int {
	n := len(word)

	count := 0
	i := 1
	for n > 0 {
		if n >= 8 {
			count += i * 8
		} else {
			count += i * n
		}

		n = n - 8
		i++
	}

	return count
}
