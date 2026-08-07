package lc3348

import (
	"sort"
	"strconv"
	"strings"
)

func smallestNumber(num string, t int64) string {
	n := len(num)
	rem := make([]int64, n+1)

	gcd := func(a, b int64) int64 {
		for b > 0 {
			a, b = b, a%b
		}
		return a
	}

	rem[0] = t
	for i := 1; i <= n; i++ {
		rem[i] = rem[i-1] / gcd(rem[i-1], int64(num[i-1]-'0'))
	}
	hasZero := false
	start := n - 1
	for i := 0; i < len(num); i++ {
		if num[i] == '0' {
			hasZero = true
			start = i
			break
		}
	}

	if rem[n] == 1 && !hasZero {
		return num
	}
	for i := start; i >= 0; i-- {
		curr := int64(num[i] - '0')
		available := n - i - 1
		new := curr + 1
		res := []int64{}
		for new <= 9 {
			tNow := rem[i] / gcd(rem[i], new)
			if available > 0 {
				pos := []int64{}
				for j := available; j >= 1; j-- {
					for k := int64(9); k >= 1; k-- {
						if tNow%k == 0 {
							tNow = tNow / k
							pos = append(pos, k)
							break
						}
					}

					if tNow == 1 {
						for len(pos) < available {
							pos = append(pos, 1)
						}
						for l, r := 0, len(pos)-1; l < r; l, r = l+1, r-1 {
							pos[l], pos[r] = pos[r], pos[l]
						}
						res = pos
						break
					}
				}
			}

			if tNow == 1 {
				var sb strings.Builder
				for j := 0; j < i; j++ {
					sb.WriteByte(num[j])
				}
				sb.WriteString(strconv.Itoa(int(new)))
				for j := 0; j < len(res); j++ {
					sb.WriteString(strconv.Itoa(int(res[j])))
				}
				return sb.String()
			} else {
				new++
			}
		}
	}

	rem2 := t
	digits := []int{}
	for rem2 > 1 {
		before := rem2
		for i := 9; i >= 2; i-- {
			if rem2%int64(i) == 0 {
				rem2 = rem2 / int64(i)
				digits = append(digits, i)
				break
			}
		}
		if rem2 == before {
			return "-1"
		}
	}
	sort.Ints(digits)

	totalLen := max(len(digits), n+1)
	var sb strings.Builder
	for i := 0; i < totalLen-len(digits); i++ {
		sb.WriteByte('1')
	}
	for _, d := range digits {
		sb.WriteByte(byte('0' + d))
	}
	return sb.String()
}
