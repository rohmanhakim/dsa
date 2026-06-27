package lc1448

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodesRecurse(root *TreeNode, parent int, goods *int, isRoot bool) {

	if root == nil {
		return
	}

	if isRoot {
		*goods = *goods + 1
	} else if !isRoot && parent <= root.Val {
		*goods = *goods + 1
	}

	goodNodesRecurse(root.Left, max(root.Val, parent), goods, false)
	goodNodesRecurse(root.Right, max(root.Val, parent), goods, false)
}

func goodNodes(root *TreeNode) int {

	goods := 0
	goodNodesRecurse(root, root.Val, &goods, true)

	return goods
}

func CountGoodNodesInBinaryTree() {
	// root := TreeNode{
	// 	Val: 3,
	// 	Left: &TreeNode{
	// 		Val: 1,
	// 		Left: &TreeNode{
	// 			Val: 3,
	// 		},
	// 	},
	// 	Right: &TreeNode{
	// 		Val: 4,
	// 		Left: &TreeNode{
	// 			Val: 1,
	// 		},
	// 		Right: &TreeNode{
	// 			Val: 5,
	// 		},
	// 	},
	// } // 4

	/*
		               -4,
		     -4                          ,5,
		null, null,	             4,             -5,
		                   -5,       2,       1,    null,
		            null,    -2, null, 5, null,null,
		                  null,-4, null,null,
		                     -4,null,
			               2,null,
		                 -1
	*/

	root := TreeNode{
		Val: -4,
		Left: &TreeNode{
			Val: -4,
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: -5,
					Right: &TreeNode{
						Val: -2,
						Right: &TreeNode{
							Val: -4,
							Left: &TreeNode{
								Val: -4,
								Left: &TreeNode{
									Val: 2,
									Left: &TreeNode{
										Val: 1,
									},
								},
							},
						},
					},
				},
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
			Right: &TreeNode{
				Val: -5,
				Left: &TreeNode{
					Val: 1,
				},
			},
		},
	} // 4

	// root := TreeNode{
	// 	Val: 3,
	// 	Left: &TreeNode{
	// 		Val: 3,
	// 		Left: &TreeNode{
	// 			Val: 4,
	// 		},
	// 		Right: &TreeNode{
	// 			Val: 2,
	// 		},
	// 	},
	// } // 3

	// root := TreeNode{
	// 	Val: 1,
	// }//1

	// root := TreeNode{
	// 	Val: 9,
	// 	Right: &TreeNode{
	// 		Val: 3,
	// 		Left: &TreeNode{
	// 			Val: 6,
	// 		},
	// 	},
	// } //1
	out := goodNodes(&root)
	fmt.Println(out)
}
