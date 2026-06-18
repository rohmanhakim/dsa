package medium

import (
	"fmt"
	"math"
)

func angleClock(hour int, minutes int) float64 {
	perMinuteAngle := float64(6)
	getAngleFromMin := func(m int) float64 {
		a := float64(m) * perMinuteAngle
		if a > 180 {
			return a - 180
		}
		return a
	}

	if hour == 12 {
		hour = 0
	}
	angleFromMin := getAngleFromMin(minutes)
	baseAngleFromHour := (float64(hour) * (5 * perMinuteAngle))
	adv := (float64(minutes) / 2)
	if adv > 180 {
		adv = adv - 180
	}
	if minutes > 30 {
		angleFromMin = angleFromMin - 180
	}

	res := math.Abs(angleFromMin - baseAngleFromHour - adv)
	if res > 180 {
		return math.Abs(360 - res)
	}
	return res

}

func angleClockClaudeSonnet(hour int, minutes int) float64 {
	minuteAngle := float64(minutes) * 6.0
	hourAngle := float64(hour%12)*30.0 + float64(minutes)*0.5

	diff := math.Abs(minuteAngle - hourAngle)
	if diff > 180 {
		diff = 360 - diff
	}
	return diff
}

func AngleBetweenHandsOfAClock() {
	hour := 9
	minutes := 41
	out := angleClock(hour, minutes)
	fmt.Println(out)
}
