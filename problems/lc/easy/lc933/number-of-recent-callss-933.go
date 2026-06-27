package lc933

import "fmt"

type RecentCounter struct {
	requests []int
}

func RCConstructor() RecentCounter {
	rc := RecentCounter{}
	// fmt.Println(rc.requests)
	return rc
}

func (this *RecentCounter) Ping(t int) int {
	this.requests = append(this.requests, t)
	num := 0
	minT := t - 3000
	for _, r := range this.requests {
		if r >= minT && r <= t {
			num++
		}
	}
	return num
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

func NumberOfRecentCalls() {
	recentCounter := RCConstructor()

	fmt.Println(recentCounter.Ping(1))
	fmt.Println(recentCounter.Ping(100))
	fmt.Println(recentCounter.Ping(3001))
	fmt.Println(recentCounter.Ping(3002))
}
