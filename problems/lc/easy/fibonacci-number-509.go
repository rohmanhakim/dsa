package easy

import "fmt"

func fibTopDown(n int) int {

	memo := make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}
	memo[0], memo[1], memo[2] = 0, 1, 1

	var dp func(n int) int
	dp = func(n int) int {
		if n == 0 {
			return 0
		}
		if n == 1 {
			return 1
		}
		if n == 2 {
			return 1
		}

		if memo[n] != -1 {
			return memo[n]
		}

		res := dp(n-1) + dp(n-2)
		memo[n] = res

		return res
	}

	return dp(n)
}

func fibBottomUp(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}

	beforePrev, prev := 1, 1
	res := 0
	for i := 3; i <= n; i++ {
		res = prev + beforePrev
		beforePrev, prev = prev, res
	}

	return res
}

func fib(n int) int {
	out := fibBottomUp(n)
	return out
}

func FibonacciNumber() {
	n := 44
	out := fib(n)
	fmt.Println(out)
}
