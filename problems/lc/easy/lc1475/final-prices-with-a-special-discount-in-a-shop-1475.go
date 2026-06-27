package lc1475

import "fmt"

func finalPrices(prices []int) []int {
	stack := []int{}
	for i, price := range prices {
		for len(stack) > 0 && price <= prices[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			prices[top] = prices[top] - price
			// prev++
		}
		stack = append(stack, i)
	}
	return prices
}

func FinalPricesWithASpecialDiscountInAShop() {
	// prices := []int{8, 4, 6, 2, 3} // 4,2,4,2,3
	// prices := []int{1, 2, 3, 4, 5} // 1,2,3,4,5
	// prices := []int{10, 1, 1, 6} // 9,0,1,6
	prices := []int{8, 7, 4, 2, 8, 1, 7, 7, 10, 1} // 1,3,2,1,7,0,0,6,9,1
	out := finalPrices(prices)

	fmt.Println(out)
}
