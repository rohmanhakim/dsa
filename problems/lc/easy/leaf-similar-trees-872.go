package easy

import "fmt"

func treeTraverseRecurse(root *TreeNode, leaves *[]int) {

	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*leaves = append(*leaves, root.Val)
		return
	}

	treeTraverseRecurse(root.Left, leaves)
	treeTraverseRecurse(root.Right, leaves)
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {

	leaves1 := []int{}
	treeTraverseRecurse(root1, &leaves1)

	leaves2 := []int{}
	treeTraverseRecurse(root2, &leaves2)

	if len(leaves1) != len(leaves2) {
		return false
	}

	for i, l := range leaves1 {
		if l != leaves2[i] {
			return false
		}
	}

	// fmt.Println(leaves)

	return true
}

func LeafSimilarTrees() {
	root1 := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}
	root2 := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
	}
	out := leafSimilar(&root1, &root2)
	fmt.Println(out)
}
