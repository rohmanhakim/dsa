package easy

import (
	"fmt"
	"math"
)

func earliestFinishTime(
	landStartTime []int,
	landDuration []int,
	waterStartTime []int,
	waterDuration []int) int {
	minFinishTime := math.MaxInt32
	for i := 0; i < len(landStartTime); i++ {
		for j := 0; j < len(waterStartTime); j++ {
			landFinish := landStartTime[i] + landDuration[i]
			waterFinish := waterStartTime[j] + waterDuration[j]

			landFirst := 0
			waterFirst := 0

			if waterStartTime[j] <= landFinish {
				landFirst = landFinish + waterDuration[j]
			} else {
				landFirst = waterFinish
			}

			if landStartTime[i] <= waterFinish {
				waterFirst = waterFinish + landDuration[i]
			} else {
				waterFirst = landFinish
			}
			minRide := min(landFirst, waterFirst)
			if minRide < minFinishTime {
				minFinishTime = minRide
			}
		}
	}
	return minFinishTime
}

func EarliestFinishTimeForLandAndWaterRidesI() {

	// landStartTime := []int{90, 12}
	// landDuration := []int{39, 86}
	// waterStartTime := []int{11, 86}
	// waterDuration := []int{44, 62} // 129

	// landStartTime := []int{37, 40}
	// landDuration := []int{53, 74}
	// waterStartTime := []int{20, 84}
	// waterDuration := []int{41, 8} // 98

	// landStartTime := []int{31, 8}
	// landDuration := []int{47, 64}
	// waterStartTime := []int{3, 7}
	// waterDuration := []int{95, 44} // 98

	// landStartTime := []int{2, 8}
	// landDuration := []int{4, 1}
	// waterStartTime := []int{6}
	// waterDuration := []int{3} // 9

	// landStartTime := []int{5}
	// landDuration := []int{3}
	// waterStartTime := []int{1}
	// waterDuration := []int{10} // 14

	// landStartTime := []int{99}
	// landDuration := []int{59}
	// waterStartTime := []int{99, 54}
	// waterDuration := []int{85, 20} // 158

	landStartTime := []int{94, 85}
	landDuration := []int{20, 54}
	waterStartTime := []int{20, 4}
	waterDuration := []int{2, 35} // 114

	out := earliestFinishTime(landStartTime, landDuration, waterStartTime, waterDuration)

	fmt.Println(out)
}
