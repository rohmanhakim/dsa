package lc739

/*
Given an array of integers `temperatures` represents the daily temperatures, return *an array* `answer` *such that* `answer[i]` *is the number of days you have to wait after the* `i<sup>th</sup>` *day to get a warmer temperature*. If there is no future day for which this is possible, keep `answer[i] == 0` instead.

**Example 1:**

**Input:** temperatures = [73,74,75,71,69,72,76,73]
**Output:** [1,1,4,2,1,1,0,0]

**Example 2:**

**Input:** temperatures = [30,40,50,60]
**Output:** [1,1,1,0]

**Example 3:**

**Input:** temperatures = [30,60,90]
**Output:** [1,1,0]

**Constraints:**

- `1 <= temperatures.length <= 10^5`
- `30 <= temperatures[i] <= 100`
*/

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	/*
			            76
		        75
		      74
			73
			          72
			      71
		                  73
				    69
	*/

	/*
		73
		74        73
		75        74
		75,71
		75,71,69
		75,72     69,71
		76        72,75
		76 73
	*/

	stack := [][]int{}
	res := make([]int, len(temperatures))

	for i, temp := range temperatures {
		for len(stack) > 0 && temp > stack[len(stack)-1][1] {
			top := stack[len(stack)-1]
			res[top[0]] = i - top[0]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, []int{i, temp})
	}
	return res
}

func DailyTemperatures() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73} // 1,1,4,2,1,1,0,0
	// temperatures := []int{30, 40, 50, 60} // 1,1,1,0
	// temperatures := []int{30, 60, 90} // 1,1,0
	out := dailyTemperatures(temperatures)
	fmt.Println(out)
}
