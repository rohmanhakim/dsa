package lc2540

import "fmt"

func getCommon(nums1 []int, nums2 []int) int {
	barr := nums1
	sarr := nums2
	if nums1[0] > nums2[0] {
		barr = nums2
		sarr = nums1
	}

	sp := 0
	bp := 0

	for sp < len(sarr) {
		for bp < len(barr) {
			sn := sarr[sp]
			bn := barr[bp]
			if sn == bn {
				return sn
			} else if bn > sn {
				sp++
			} else {
				bp++
			}
		}
		sp++
	}

	return -1
}

func getCommonClaudeSonnet(nums1 []int, nums2 []int) int {
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			return nums1[i]
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return -1
}

func MinimumCommonValue() {
	// 2
	// nums1 := []int{1, 2, 3}
	// nums2 := []int{2, 4}

	// 2
	// nums1 := []int{1, 2, 3, 6}
	// nums2 := []int{2, 3, 4, 5}

	// 2
	// nums2 := []int{2, 4}
	// nums1 := []int{1, 2}

	// 13
	nums1 := []int{6, 13, 18, 18, 28, 34, 37, 39, 46, 50, 52, 54, 62, 63, 65, 66, 75, 80, 97, 98}
	nums2 := []int{10, 13, 13, 19, 27, 33, 40, 41, 43, 46, 56, 61, 69, 72, 78, 79, 82, 88, 91, 94}
	out := getCommon(nums1, nums2)
	fmt.Println(out)
}
