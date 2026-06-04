package lib

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SliceToBinaryTreeNodes(nodes []*int) *TreeNode {
	if len(nodes) == 0 || nodes[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *nodes[0]}
	queue := []*TreeNode{root}

	i := 1
	for len(queue) > 0 && i < len(nodes) {
		curr := queue[0]
		queue = queue[1:]

		if i < len(nodes) && nodes[i] != nil {
			curr.Left = &TreeNode{Val: *nodes[i]}
			queue = append(queue, curr.Left)
		}
		i++

		if i < len(nodes) && nodes[i] != nil {
			curr.Right = &TreeNode{Val: *nodes[i]}
			queue = append(queue, curr.Right)
		}
		i++
	}

	return root
}

func PrintBFS(root *TreeNode) {
	queue := []*TreeNode{root}
	level := 0
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			fmt.Printf("%d ", queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		fmt.Println("")
		queue = queue[n:]
		level++
	}
}

// PrintTreeASCII prints the tree using box-drawing characters (directory tree style).
// This is the easiest way to visualize tree structure without external dependencies.
func PrintTreeASCII(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	printChildren(root, "")
}

func printChildren(node *TreeNode, prefix string) {
	if node.Left == nil && node.Right == nil {
		return
	}

	hasLeft := node.Left != nil
	hasRight := node.Right != nil

	if hasLeft {
		isLast := !hasRight
		fmt.Printf("%s", prefix)
		if isLast {
			fmt.Print("└── ")
		} else {
			fmt.Print("├── ")
		}
		fmt.Println(node.Left.Val)

		newPrefix := prefix
		if isLast {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		printChildren(node.Left, newPrefix)
	}

	if hasRight {
		fmt.Printf("%s", prefix)
		fmt.Print("└── ")
		fmt.Println(node.Right.Val)

		newPrefix := prefix + "    "
		printChildren(node.Right, newPrefix)
	}
}
