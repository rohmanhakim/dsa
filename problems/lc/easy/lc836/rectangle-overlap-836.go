package lc836

func isRectangleOverlap(rec1 []int, rec2 []int) bool {

	ax1, ay1, ax2, ay2 := rec1[0], rec1[1], rec1[2], rec1[3]
	bx1, by1, bx2, by2 := rec2[0], rec2[1], rec2[2], rec2[3]

	return ax1 < bx2 && ay1 < by2 && ax2 > bx1 && ay2 > by1
}
