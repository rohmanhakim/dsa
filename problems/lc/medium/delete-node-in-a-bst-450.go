package medium

import (
	"github.com/rohmanhakim/leetcode/lib"
)

func bstDeleteNodeDFS(root **lib.TreeNode, key int) bool {
	if *root == nil {
		return false
	}

	if key < (*root).Val {
		return bstDeleteNodeDFS(&(*root).Left, key)
	} else if key > (*root).Val {
		return bstDeleteNodeDFS(&(*root).Right, key)
	} else {
		if (*root).Left == nil {
			*root = (*root).Right
			return true
		} else if (*root).Right == nil {
			*root = (*root).Left
			return true
		} else {
			deleteNode := *root
			largestBST := largestNode((*root).Left)
			deleteNode.Val = largestBST.Val

			return bstDeleteNodeDFS(&deleteNode.Left, largestBST.Val)
		}
	}
}

func largestNode(root *lib.TreeNode) *lib.TreeNode {
	if root == nil {
		return nil
	}

	maxNode := root

	left := largestNode(root.Left)
	right := largestNode(root.Right)

	if left != nil && left.Val > maxNode.Val {
		return left
	}

	if right != nil && right.Val > maxNode.Val {
		return right
	}

	return maxNode
}

func deleteNodeClaudeSonnet(root *lib.TreeNode, key int) *lib.TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNodeClaudeSonnet(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNodeClaudeSonnet(root.Right, key)
	} else {
		// Found the node to delete
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		// Two children: replace value with in-order predecessor,
		// then delete predecessor from left subtree
		predecessor := inOrderPredecessor(root.Left)
		root.Val = predecessor.Val
		root.Left = deleteNodeClaudeSonnet(root.Left, predecessor.Val)
	}

	return root
}

// Rightmost node of a subtree = largest value
func inOrderPredecessor(node *lib.TreeNode) *lib.TreeNode {
	for node.Right != nil {
		node = node.Right
	}
	return node
}

func deleteNode(root *lib.TreeNode, key int) *lib.TreeNode {
	bstDeleteNodeDFS(&root, key)
	return root
}

func DeleteNodeInABST() {
	root := lib.SliceToBinaryTreeNodes([]*int{
		lib.Ptr(5),
		lib.Ptr(3),
		lib.Ptr(6),
		lib.Ptr(2),
		lib.Ptr(4),
		nil,
		lib.Ptr(7),
	})

	key := 3 // 5,4,6,2,ni;,nil,7

	out := deleteNodeClaudeSonnet(root, key)
	lib.PrintTreeASCII(out)

}
