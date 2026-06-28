package lc714

import (
	"fmt"
	"math"
)

func maxProfit(prices []int, fee int) int {
	n := len(prices)

	cashNotOwningStock := 0
	cashOwningStock := math.MinInt32

	for i := 0; i < n; i++ {
		stratBuy := cashNotOwningStock - prices[i]
		stratHold := cashOwningStock

		stratSell := cashOwningStock + prices[i] - fee
		stratAvoid := cashNotOwningStock

		cashOwningStock = max(stratBuy, stratHold)
		cashNotOwningStock = max(stratSell, stratAvoid)
	}
	return cashNotOwningStock
}

func BestTimetoBuyandSellStockWithTransactionFee() {
	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2
	out := maxProfit(prices, fee)
	fmt.Println(out)
}
