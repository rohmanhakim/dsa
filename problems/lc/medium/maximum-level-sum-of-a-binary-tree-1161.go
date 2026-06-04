package medium

import (
	"fmt"
	"math"

	"github.com/rohmanhakim/leetcode/lib"
)

func maxLevelSum(root *lib.TreeNode) int {
	maximum := math.MinInt32
	maxLevel := 0

	queue := []*lib.TreeNode{root}
	level := 0
	for len(queue) > 0 {
		n := len(queue)
		sum := 0
		level++
		for i := 0; i < n; i++ {
			// fmt.Printf("%d ", queue[i].Val)
			sum += queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		// fmt.Printf("sum: %d, ", sum)
		// fmt.Println("")
		if sum > maximum {
			maximum = sum
			maxLevel = level
		}
		queue = queue[n:]
	}

	return maxLevel
}

func MaximumLevelSumOfABinaryTree() {
	root := lib.SliceToBinaryTreeNodes([]*int{
		lib.Ptr(1),
		lib.Ptr(7),
		lib.Ptr(0),
		lib.Ptr(7),
		lib.Ptr(-8),
		nil,
		nil,
	})
	out := maxLevelSum(root)
	fmt.Println(out)
}
