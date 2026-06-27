package lc1833

import "fmt"

func maxIceCream(costs []int, coins int) int {
	freq := make([]int, 100001)

	for _, cost := range costs {
		freq[cost]++
	}

	i := 1
	total := 0
	for coins > 0 && coins >= i && i < len(freq) {
		if freq[i] == 0 {
			i++
			continue
		}

		canBuy := min(freq[i], coins/i)
		coins -= canBuy * i
		total += canBuy
		i++
	}

	return total
}

func MaximumIceCreamBar() {
	cost := []int{4, 7, 6, 4, 4, 2, 2, 4, 8, 8}

	coins := 41
	out := maxIceCream(cost, coins)
	fmt.Println(out)
}
