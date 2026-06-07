package temp

import (
	"fmt"
	"math"
)

/*
When evaluating stock market trading strategies, it is useful to determine the maximum pos-
sible profit that can be made by trading a certain stock.

Given the daily price of a stock, computes the maximum profit that can be made on the last day.

Assume that you are allowed to own no more than 1 share at any time, and that
you have an unlimited budget.

Example 1:
Input: prices = [2, 5, 1]
Output: 3
Explanation:
The stock price over several days is [2, 5, 1]. The best strategy is to buy a share
on the first day for price 2, then sell it on the second day for price 5, obtaining a profit of 3.

Example 2:
Input: prices = [2, 5, 1, 3]
Output: 5
Explanation:
The stock price over several days is [2, 5, 1, 3]. The best strategy is to buy a
share on the first day for price 2, then sell it on the second day for price 5, obtaining a profit
of 3; then buy it again on the third day for price 1, and sell it on the fourth day for price 3,
obtaining an overall profit of 5.

Constraints:
- 1 <= prices.length <= 1000
- 0 <= prices[i] <= 999
*/

func getBestProfit2(prices []int, day int, haveStock bool) int {
	if day < 0 {
		if !haveStock {
			// initial state before the first day: should have no stock and have no profit
			return 0
		} else {
			// not allowed to have initial stock
			// give very low profit
			return math.MinInt32
		}
	}
	price := prices[day]

	if haveStock {
		// to own a stock in a specific day, we can only either buying or holding stock
		stratBuy := getBestProfit2(prices, day-1, false) - price
		stratHold := getBestProfit2(prices, day-1, true)
		return max(stratBuy, stratHold)
	} else {
		// to not own stock in a specific day, we can only either selling or avoid buying stock
		stratSell := getBestProfit2(prices, day-1, true) + price
		stratAvoid := getBestProfit2(prices, day-1, false)
		return max(stratSell, stratAvoid)
	}
}

func maxProfit2(prices []int) int {
	lastDay := len(prices) - 1
	haveStock := false

	// at the end of the last day, no shares should be owned
	return getBestProfit2(prices, lastDay, haveStock)
}

func OptimalStockMarket2() {
	prices := []int{2, 5, 1, 3}
	out := maxProfit2(prices)
	fmt.Println(out)
}
