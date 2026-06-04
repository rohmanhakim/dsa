package easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 1
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		curr := queue[0]
		if curr.Val > 100 {
			depth = curr.Val - 100
			queue = queue[1:]
			continue
		}

		queue = queue[1:]

		if curr.Left != nil {
			queue = append(queue, &TreeNode{
				Val: depth + 1 + 100,
			})
			queue = append(queue, curr.Left)

		}
		if curr.Right != nil {
			queue = append(queue, &TreeNode{
				Val: depth + 1 + 100,
			})
			queue = append(queue, curr.Right)

		}

	}

	return depth
}
*/

func maxDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	depth := 0
	for len(queue) > 0 {
		level := len(queue)
		for level > 0 {
			front := queue[0]
			if front.Left != nil {
				queue = append(queue, front.Left)
			}
			if front.Right != nil {
				queue = append(queue, front.Right)
			}
			level--
			queue = queue[1:]
		}
		depth++
	}
	return depth
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	// right := maxDepth(root.Right)

	return left + 1
}

func MaximumDepthOfBinaryTree() {
	root := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	} //3

	// root := TreeNode{
	// 	Val: 1,
	// 	Right: &TreeNode{
	// 		Val: 2,
	// 	},
	// } //2

	// root := TreeNode{
	// 	Val: 1,
	// 	Left: &TreeNode{
	// 		Val: 2,
	// 		Left: &TreeNode{
	// 			Val: 4,
	// 		},
	// 	},
	// 	Right: &TreeNode{
	// 		Val: 3,
	// 		Right: &TreeNode{
	// 			Val: 5,
	// 		},
	// 	},
	// } // 3

	// root := TreeNode{
	// 	Val: 1,
	// 	Left: &TreeNode{
	// 		Val: 2,
	// 		Left: &TreeNode{
	// 			Val: 3,
	// 			Left: &TreeNode{
	// 				Val: 4,
	// 				Left: &TreeNode{
	// 					Val: 5,
	// 				},
	// 			},
	// 		},
	// 	},
	// }

	out := maxDepth(&root)
	fmt.Println(out)
}
