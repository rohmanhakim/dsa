package lc236

import (
	"fmt"
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lcaDFS(root *TreeNode, path *[]*TreeNode, pPath *[]*TreeNode, qPath *[]*TreeNode, p int, q int) {
	if root == nil {
		return
	}

	*path = append(*path, root)

	if root.Val == p {
		*pPath = slices.Clone(*path)
	}
	if root.Val == q {
		*qPath = slices.Clone(*path)
	}

	lcaDFS(root.Left, path, pPath, qPath, p, q)
	lcaDFS(root.Right, path, pPath, qPath, p, q)
	if len(*path) > 0 {
		a := *path
		*path = a[:len(a)-1]
	}
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	path, pPath, qPath := []*TreeNode{}, []*TreeNode{}, []*TreeNode{}
	lcaDFS(root, &path, &pPath, &qPath, p.Val, q.Val)

	n := len(qPath)
	leastPath := qPath
	if len(pPath) < len(qPath) {
		leastPath = pPath
		n = len(pPath)
	}

	lca := &TreeNode{}
	for i := 0; i < n; i++ {
		if pPath[i].Val == qPath[i].Val {
			lca = leastPath[i]
		} else {
			break
		}
	}
	return lca
}

func lowestCommonAncestorClaudeSonnet(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// If we find either target, bubble it up immediately
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// p and q are on opposite sides — this node is the LCA
	if left != nil && right != nil {
		return root
	}

	// One side found both (or one), bubble up whichever is non-nil
	if left != nil {
		return left
	}
	return right
}

func LowestCommonAncestorOfABinaryTree() {

	p := TreeNode{
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
	}

	q := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 8,
		},
	}

	root := TreeNode{
		Val:   3,
		Left:  &p,
		Right: &q,
	}
	out := lowestCommonAncestor(&root, &p, &q)
	fmt.Println(out.Val)
}
