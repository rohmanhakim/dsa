package lc636

import (
	"fmt"
	"strconv"
	"strings"
)

// func  exclusiveTime(n int, logs []string) []int {

// 	stack := []int{}
// 	exc := make([]int, n)
// 	start := make([]int, n)

// 	for i := 0; i < len(logs); i++ {
// 		parts := strings.Split(logs[i], ":")
// 		id, _ := strconv.Atoi(parts[0])
// 		state := parts[1]
// 		ts, _ := strconv.Atoi(parts[2])
// 		top := -1
// 		if len(stack) > 0 {
// 			top = stack[len(stack)-1]
// 		}
// 		if state == "start" {
// 			stack = append(stack, id)
// 			start[id] = ts
// 			if top >= 0 {
// 				fmt.Printf("start -> exc + %d\n", ts)
// 				exc[top] = ts
// 			}
// 		} else {
// 			duration := ts - start[id] + 1
// 			fmt.Printf("end -> exc + %d\n", duration)
// 			exc[id] += duration
// 			stack = stack[:len(stack)-1]
// 			if len(stack) > 0 {
// 				top = stack[len(stack)-1]
// 				start[top] = ts + 1
// 			}
// 		}
// 	}

// 	return exc
// }

func exclusiveTime(n int, logs []string) []int {
	stack := []int{}
	exc := make([]int, n)

	time := 0

	for _, log := range logs {
		parts := strings.Split(log, ":")
		id, _ := strconv.Atoi(parts[0])
		state := parts[1]
		ts, _ := strconv.Atoi(parts[2])
		top := -1

		if len(stack) > 0 {
			top = stack[len(stack)-1]
		}
		if state == "start" {
			if top != -1 {
				exc[top] += ts - time
				time = ts
			}
			stack = append(stack, id)
		} else {
			duration := ts - time + 1
			exc[id] += duration
			time = ts + 1
			if top != -1 {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return exc
}

func exclusiveTimeClaudeSonnet(n int, logs []string) []int {
	stack := []int{}
	exc := make([]int, n)
	prevTime := 0

	for _, log := range logs {
		parts := strings.Split(log, ":")
		id, _ := strconv.Atoi(parts[0])
		ts, _ := strconv.Atoi(parts[2])

		if parts[1] == "start" {
			// Credit elapsed time to whoever was running
			if len(stack) > 0 {
				exc[stack[len(stack)-1]] += ts - prevTime
			}
			stack = append(stack, id)
			prevTime = ts
		} else {
			// Credit this function for its final slice, then pop
			exc[id] += ts - prevTime + 1
			stack = stack[:len(stack)-1]
			prevTime = ts + 1
		}
	}

	return exc
}

func ExclusiveTimeOfFunctions() {
	// n := 2
	// logs := []string{
	// 	"0:start:0", "1:start:2", "1:end:5", "0:end:6",
	// } // [3,4]
	// n := 1
	// logs := []string{
	// 	"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7",
	// } // [8]
	// n := 2
	// logs := []string{
	// 	"0:start:0", "0:start:2", "0:end:5", "1:start:6", "1:end:6", "0:end:7",
	// } // [7,1]
	n := 8
	logs := []string{
		"0:start:0", "1:start:5", "2:start:6", "3:start:9", "4:start:11", "5:start:12", "6:start:14", "7:start:15", "1:start:24", "1:end:29", "7:end:34", "6:end:37", "5:end:39", "4:end:40", "3:end:45", "0:start:49", "0:end:54", "5:start:55", "5:end:59", "4:start:63", "4:end:66", "2:start:69", "2:end:70", "2:start:74", "6:start:78", "0:start:79", "0:end:80", "6:end:85", "1:start:89", "1:end:93", "2:end:96", "2:end:100", "1:end:102", "2:start:105", "2:end:109", "0:end:114",
	} // [20, 14, 35, 7, 6, 9, 10, 14]
	out := exclusiveTime(n, logs)
	fmt.Println(out)
}
