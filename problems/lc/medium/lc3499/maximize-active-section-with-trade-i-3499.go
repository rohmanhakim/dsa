package lc3499

/*
func maxActiveSectionsAfterTrade(s string) int {

	type region struct {
		lZeros int
		lOnes  int
		ones   int
		rOnes  int
		rZeros int
	}

	n := len(s)
	var prev *region
	curr := &region{}

	pointer := 0

	regions := []region{}
	for pointer < n {
		ch := s[pointer]

		if ch == '0' {
			if prev != nil {
				prev.rOnes = curr.ones
			}
			if curr.ones == 0 {
				curr.lZeros++
			} else {
				curr.rZeros++
			}
		} else if ch == '1' && curr.lZeros != 0 {
			if curr.rZeros == 0 {
				curr.ones++
			} else {
				regions = append(regions, *curr)
				prev = curr
				curr = &region{
					lZeros: prev.rZeros,
					ones:   1,
				}
			}
		} else if ch == '1' {
			curr.lOnes++
		}
		pointer++
	}
	regions = append(regions, *curr)

	fmt.Println(regions)
	fmt.Println(curr)

	// sort.Slice(regions, func(i, j int) bool {
	// 	iCount := regions[i].lZeros + regions[i].ones + regions[j].rZeros
	// 	jCount := regions[j].lZeros + regions[j].ones + regions[j].rZeros
	// 	return iCount > jCount
	// })

	maxRegion := 0
	for _, r := range regions {
		count := r.lZeros + r.ones + r.rZeros
		if count > maxRegion {
			maxRegion = count
		}
	}

	fmt.Println(maxRegion)
	return -1
}
*/

func maxActiveSectionsAfterTrade(s string) int {

	maxActive := 0
	n := len(s)

	prefix := make([]int, n+1)
	for i := 1; i <= n; i++ {
		val := 0
		if s[i-1] == '1' {
			val = 1
		}
		prefix[i] = prefix[i-1] + val
	}
	for i := 1; i < n; i++ {
		ch := s[i]
		prev := s[i-1]
		toTheLeft := 0
		toTheRight := 0
		count := 0

		toTheLeft = prefix[i+1]
		if ch == '1' && prev == '0' && i < n-1 {
			curr := i - 1

			toTheLeft = 0
			for curr >= 0 {
				count++
				if curr > 0 && s[curr] == '1' && s[curr-1] == '0' {
					break
				}
				curr--
			}
			if curr >= 0 {
				toTheLeft += prefix[curr]
			}
			for curr = i; curr < len(s) && s[curr] == '1'; curr++ {
				count++
			}

			if curr == n {
				count = prefix[n]
				continue
			}

			for curr < len(s) {
				count++
				if curr > 0 && (curr < n-1 && s[curr] == '1' && s[curr+1] == '0') || (curr == n-1 && s[curr] == '1') {
					break
				}
				curr++
			}

			if curr < n-1 {
				toTheRight = prefix[n] - prefix[curr+1]
			}
		}

		if count+toTheLeft+toTheRight > maxActive {
			maxActive = count + toTheLeft + toTheRight
		}
	}

	return maxActive
}
