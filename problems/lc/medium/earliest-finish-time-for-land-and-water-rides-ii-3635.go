package medium

import (
	"fmt"
	"math"
	"sort"
)

func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	type Ride struct {
		start int
		dur   int
	}

	makeSorted := func(startTime []int, duration []int) []Ride {
		rides := make([]Ride, len(startTime))

		for i := range len(startTime) {
			rides[i] = Ride{
				start: startTime[i],
				dur:   duration[i],
			}
		}

		sort.Slice(rides, func(i, j int) bool {
			return rides[i].start < rides[j].start
		})

		return rides
	}

	makePrefixMinDur := func(rides []Ride) []int {
		prefix := make([]int, len(rides))

		prefix[0] = rides[0].dur

		for i := 1; i < len(prefix); i++ {
			prefix[i] = min(prefix[i-1], rides[i].dur)
		}
		return prefix
	}

	makeSuffixMinFinishTime := func(rides []Ride) []int {
		suffix := make([]int, len(rides))
		n := len(suffix)
		suffix[n-1] = rides[n-1].start + rides[n-1].dur

		for i := n - 2; i >= 0; i-- {
			suffix[i] = min(suffix[i+1], rides[i].start+rides[i].dur)
		}
		return suffix
	}

	findBestRides := func(ride1Finish int, ride2 []Ride, prefixMinDur []int, suffixMinFinish []int) int {
		n := len(ride2)

		splitIndex := sort.Search(n, func(i int) bool { // fmt.Printf("r2.s: %d\n", ride2[i].start)
			return ride2[i].start > ride1Finish
		})

		bestInLateGroup := math.MaxInt32
		bestInEarlyGroup := math.MaxInt32
		if splitIndex > 0 {
			bestInEarlyGroup = ride1Finish + prefixMinDur[splitIndex-1]
		}

		if splitIndex < n {
			bestInLateGroup = suffixMinFinish[splitIndex]
		}

		return min(bestInEarlyGroup, bestInLateGroup)
	}

	tryOrder := func(ride1 []Ride, ride2 []Ride) int {
		prefixMinDur := makePrefixMinDur(ride2)
		suffixMinFinishTime := makeSuffixMinFinishTime(ride2)

		best := math.MaxInt32
		for _, r := range ride1 {
			ride1Finish := r.start + r.dur
			res := findBestRides(ride1Finish, ride2, prefixMinDur, suffixMinFinishTime)
			fmt.Printf("r1f: %d res: %d\n", ride1Finish, res)
			best = min(res, best)
		}
		return best
	}

	land := makeSorted(landStartTime, landDuration)
	water := makeSorted(waterStartTime, waterDuration)

	return min(
		tryOrder(land, water),
		tryOrder(water, land),
	)
}

/*
                                82-------------124
	   14---------44
6-------------------------------------97
                       54------------------------125
*/

func EarliestFinishTimeForLandAndWaterRidesII() {
	// landStartTime := []int{2, 8}
	// landDuration := []int{4, 1}
	// waterStartTime := []int{6}
	// waterDuration := []int{3} // 9

	// landStartTime := []int{5}
	// landDuration := []int{3}
	// waterStartTime := []int{1}
	// waterDuration := []int{10} // 14

	// landStartTime := []int{63, 11, 6}
	// landDuration := []int{30, 99, 1}
	// waterStartTime := []int{69, 38}
	// waterDuration := []int{83, 45} // 83

	// landStartTime := []int{23, 92, 7, 43, 28, 37, 63}
	// landDuration := []int{54, 62, 32, 35, 56, 58, 85}
	// waterStartTime := []int{15, 33, 5, 75, 91, 78, 76, 25, 38}
	// waterDuration := []int{69, 95, 82, 52, 11, 13, 63, 59, 74} // 91

	landStartTime := []int{82, 14}
	landDuration := []int{42, 30}
	waterStartTime := []int{6, 54}
	waterDuration := []int{91, 71} // 125

	out := earliestFinishTime(landStartTime, landDuration, waterStartTime, waterDuration)

	// out := sort.Search(1, func(i int) bool {
	// 	return []int{2, 8}[i] <= 6
	// })

	fmt.Println(out)
}
