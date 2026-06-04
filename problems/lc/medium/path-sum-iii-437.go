package medium

import "fmt"

func pathSumRecurse(root *TreeNode, targetSum int, count *int, path *[]int) {
	if root == nil {
		return
	}

	pathCopy := *path

	res := root.Val
	if res == targetSum {
		*count++
	}
	for i := len(pathCopy) - 1; i >= 0; i-- {
		res += pathCopy[i]
		if res == targetSum {
			*count++
		}
	}

	*path = append(*path, root.Val)

	pathSumRecurse(root.Left, targetSum, count, path)
	pathSumRecurse(root.Right, targetSum, count, path)

	a := *path
	*path = a[:len(a)-1]
}

func pathSum(root *TreeNode, targetSum int) int {

	count := 0
	path := []int{}
	pathSumRecurse(root, targetSum, &count, &path)

	return count
}

func PathSumIII() {
	// root := TreeNode{
	// 	Val: 10,
	// 	Left: &TreeNode{
	// 		Val: 5,
	// 		Left: &TreeNode{
	// 			Val: 3,
	// 			Left: &TreeNode{
	// 				Val: 3,
	// 			},
	// 			Right: &TreeNode{
	// 				Val: -2,
	// 			},
	// 		},
	// 		Right: &TreeNode{
	// 			Val: 2,
	// 			Right: &TreeNode{
	// 				Val: 1,
	// 			},
	// 		},
	// 	},
	// 	Right: &TreeNode{
	// 		Val: -3,
	// 		Right: &TreeNode{
	// 			Val: 11,
	// 		},
	// 	},
	// }
	// targetSum := 8 // 3

	/*
		                           1
							0                1
						1        2       0        -1
					  0   1    -1 0    -1 0     1    0
	*/

	// root := TreeNode{
	// 	Val: 5,
	// 	Left: &TreeNode{
	// 		Val: 4,
	// 		Left: &TreeNode{
	// 			Val: 11,
	// 			Left: &TreeNode{
	// 				Val: 7,
	// 			},
	// 			Right: &TreeNode{
	// 				Val: 2,
	// 			},
	// 		},
	// 	},
	// 	Right: &TreeNode{
	// 		Val: 8,
	// 		Left: &TreeNode{
	// 			Val: 13,
	// 		},
	// 		Right: &TreeNode{
	// 			Val: 4,
	// 			Left: &TreeNode{
	// 				Val: 5,
	// 			},
	// 			Right: &TreeNode{
	// 				Val: 1,
	// 			},
	// 		},
	// 	},
	// }
	// targetSum := 22 // 3

	// root := TreeNode{
	// 	Val: 0,
	// 	Left: &TreeNode{
	// 		Val: 1,
	// 	}, Right: &TreeNode{
	// 		Val: 1,
	// 	},
	// }
	// targetSum := 1 // 4

	root := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: -1,
				},
				Right: &TreeNode{
					Val: 0,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: -1,
				},
				Right: &TreeNode{
					Val: 0,
				},
			},
			Right: &TreeNode{
				Val: -1,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 0,
				},
			},
		},
	}
	targetSum := 2 // 13

	out := pathSum(&root, targetSum)
	fmt.Println(out)
}
