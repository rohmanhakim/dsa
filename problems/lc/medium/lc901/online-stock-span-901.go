package lc901

/*
---
title: "Online Stock Span"
source: "https://leetcode.com/problems/online-stock-span/description/?envType=study-plan-v2&envId=leetcode-75"
author:
published:
created: 2026-07-03
description: "Can you solve this real interview question? Online Stock Span - Design an algorithm that collects daily price quotes for some stock and returns the span of that stock's price for the current day.The span of the stock's price in one day is the maximum number of consecutive days (starting from that day and going backward) for which the stock price was less than or equal to the price of that day. * For example, if the prices of the stock in the last four days is [7,2,1,2] and the price of the stock today is 2, then the span of today is 4 because starting from today, the price of the stock was less than or equal 2 for 4 consecutive days. * Also, if the prices of the stock in the last four days is [7,34,1,2] and the price of the stock today is 8, then the span of today is 3 because starting from today, the price of the stock was less than or equal 8 for 3 consecutive days.Implement the StockSpanner class: * StockSpanner() Initializes the object of the class. * int next(int price) Returns the span of the stock's price given that today's price is price. Example 1:Input[\"StockSpanner\", \"next\", \"next\", \"next\", \"next\", \"next\", \"next\", \"next\"][[], [100], [80], [60], [70], [60], [75], [85]]Output[null, 1, 1, 1, 2, 1, 4, 6]ExplanationStockSpanner stockSpanner = new StockSpanner();stockSpanner.next(100); // return 1stockSpanner.next(80);  // return 1stockSpanner.next(60);  // return 1stockSpanner.next(70);  // return 2stockSpanner.next(60);  // return 1stockSpanner.next(75);  // return 4, because the last 4 prices (including today's price of 75) were less than or equal to today's price.stockSpanner.next(85);  // return 6 Constraints: * 1 <= price <= 105 * At most 104 calls will be made to next."
tags:
  - "clippings"
---
Design an algorithm that collects daily price quotes for some stock and returns **the span** of that stock's price for the current day.

The **span** of the stock's price in one day is the maximum number of consecutive days (starting from that day and going backward) for which the stock price was less than or equal to the price of that day.

- For example, if the prices of the stock in the last four days is `[7,2,1,2]` and the price of the stock today is `2`, then the span of today is `4` because starting from today, the price of the stock was less than or equal `2` for `4` consecutive days.
- Also, if the prices of the stock in the last four days is `[7,34,1,2]` and the price of the stock today is `8`, then the span of today is `3` because starting from today, the price of the stock was less than or equal `8` for `3` consecutive days.

Implement the `StockSpanner` class:

- `StockSpanner()` Initializes the object of the class.
- `int next(int price)` Returns the **span** of the stock's price given that today's price is `price`.

**Example 1:**

**Input**
["StockSpanner", "next", "next", "next", "next", "next", "next", "next"]
[[], [100], [80], [60], [70], [60], [75], [85]]
**Output**
[null, 1, 1, 1, 2, 1, 4, 6]

**Explanation**
StockSpanner stockSpanner = new StockSpanner();
stockSpanner.next(100); // return 1
stockSpanner.next(80);  // return 1
stockSpanner.next(60);  // return 1
stockSpanner.next(70);  // return 2
stockSpanner.next(60);  // return 1
stockSpanner.next(75);  // return 4, because the last 4 prices (including today's price of 75) were less than or equal to today's price.
stockSpanner.next(85);  // return 6

You will be given this empty solution as starter :
```
type StockSpanner struct {

}

func Constructor() StockSpanner {

}


func (this *StockSpanner) Next(price int) int {

}
```

Your StockSpanner object will be instantiated and called as such:
```
 obj := Constructor();
 param_1 := obj.Next(price);
```
*/

import (
	"fmt"
)

type Price struct {
	val  int
	span int
}

type StockSpanner struct {
	stack []Price
}

func Constructor() StockSpanner {
	return StockSpanner{
		stack: []Price{},
	}
}

func (this *StockSpanner) Next(price int) int {
	span := 1

	for len(this.stack) > 0 && price >= this.stack[len(this.stack)-1].val {
		top := this.stack[len(this.stack)-1]
		span += top.span
		this.stack = this.stack[:len(this.stack)-1]
	}

	this.stack = append(this.stack, Price{
		val:  price,
		span: span,
	})

	return span
}

func OnlineStockSpan() {
	/*
		100
		80                  100
		60					100 80
		60 70
		60					70
		60 75
		60 75 85
	*/

	/*
		32
		82				32
		82,73
		99				82,73
		99,91
	*/
	obj := Constructor()
	fmt.Println(obj.Next(32))
	fmt.Println(obj.Next(82))
	fmt.Println(obj.Next(73))
	fmt.Println(obj.Next(99))
	fmt.Println(obj.Next(91))
}
