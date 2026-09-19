package lc1401

func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {

	clamp := func(val, minVal, maxVal int) int {
		return max(minVal, min(val, maxVal))
	}

	closestX := clamp(xCenter, x1, x2)
	closestY := clamp(yCenter, y1, y2)

	distSquared := func(ax, bx, ay, by int) int {
		xVal := (ax - bx) * (ax - bx)
		yVal := (ay - by) * (ay - by)

		return xVal + yVal
	}
	return distSquared(xCenter, closestX, yCenter, closestY) <= radius*radius
}
