package easy

import "fmt"

func check(nums []int) bool {

	if len(nums) == 1 {
		return true
	}

	inc := 0
	dec := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			dec++
		} else {
			inc++
		}
		if i == len(nums)-1 {
			if dec == 1 && nums[0] >= nums[len(nums)-1] {
				return true
			} else if nums[0] <= nums[len(nums)-1] && dec == 0 {
				return true
			}
		}
	}

	return false
}

func checkClaudeSonnet(nums []int) bool {
	/*
		The core idea : By using (i+1) % n, the loop naturally checks the
		wrap-around between nums[n-1] and nums[0] without any special-casing.
		A valid sorted+rotated array looks like:

		[3, 4, 5, 1, 2]
	           ↑
	   	only 1 drop allowed (5 → 1)

		If there are 2 or more drops, the array can't be a rotation of a
		sorted array — return false immediately.
	*/
	drops := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		if nums[i] > nums[(i+1)%n] {
			drops++
		}
		if drops > 1 {
			return false
		}
	}

	return true
}

func CheckIfArrayIsSortedAndRotated() {
	// nums := []int{3, 4, 5, 1, 2}
	// out := check(nums) // true
	nums := []int{2, 1, 3, 4}
	out := check(nums) // false
	// nums := []int{1, 2, 3}
	// out := check(nums) // true
	// nums := []int{1, 2, 1}
	// out := check(nums) // true
	// nums := []int{1, 1, 1}
	// out := check(nums) // true
	fmt.Println(out)
}
