package medium

import (
	"fmt"

	"github.com/rohmanhakim/leetcode/lib"
)

func createBinaryTree(descriptions [][]int) *lib.TreeNode {

	treeMap := make(map[int]*lib.TreeNode)
	children := make(map[int]bool)
	rootCandidates := []*lib.TreeNode{}
	// var root *lib.TreeNode

	for i := 0; i < len(descriptions); i++ {
		nodeCandidate := descriptions[i]
		parentVal := nodeCandidate[0]
		childVal := nodeCandidate[1]
		isLeft := nodeCandidate[2]

		if parent, ok := treeMap[parentVal]; !ok {
			p := &lib.TreeNode{
				Val: parentVal,
			}

			var c *lib.TreeNode
			if child, childExist := treeMap[childVal]; !childExist {
				c = &lib.TreeNode{
					Val: childVal,
				}
			} else {
				fmt.Printf("p: %d, c: %d, child is exist\n", parentVal, childVal)
				c = child
			}

			if isLeft == 1 {
				p.Left = c
			} else {
				p.Right = c
			}
			treeMap[parentVal] = p
			treeMap[childVal] = c
			children[childVal] = true
			rootCandidates = append(rootCandidates, p)
		} else {
			var c *lib.TreeNode
			if child, childExist := treeMap[childVal]; !childExist {
				c = &lib.TreeNode{
					Val: childVal,
				}
			} else {
				fmt.Printf("p: %d, c: %d, child is exist\n", parentVal, childVal)
				c = child
			}
			if isLeft == 1 {
				parent.Left = c
			} else {
				parent.Right = c
			}
			treeMap[childVal] = c
			children[childVal] = true
			if _, isChild := children[parentVal]; !isChild {
				rootCandidates = append(rootCandidates, parent)
			}
		}
		// lib.PrintTreeASCII(rootCandidates[len(rootCandidates)-1])
	}

	var root *lib.TreeNode
	for _, rc := range rootCandidates {
		if _, isChild := children[rc.Val]; !isChild {
			root = rc
		}
	}
	return root
}

func createBinaryTreeClaudeSonnet(descriptions [][]int) *lib.TreeNode {
	nodes := make(map[int]*lib.TreeNode)
	isChild := make(map[int]bool)

	getOrCreate := func(val int) *lib.TreeNode {
		if nodes[val] == nil {
			nodes[val] = &lib.TreeNode{Val: val}
		}
		return nodes[val]
	}

	for _, d := range descriptions {
		parent, child, left := d[0], d[1], d[2]
		p := getOrCreate(parent)
		c := getOrCreate(child)
		if left == 1 {
			p.Left = c
		} else {
			p.Right = c
		}
		isChild[child] = true
	}

	/*
		Set difference for source/root finding:
			given a list of edges, the root/source is:

			all nodes minus the nodes that appear as children.

			Two maps, one scan at the end.
	*/
	for val, node := range nodes {
		if !isChild[val] {
			return node
		}
	}
	return nil
}

func CreateBinaryTreeFromDescriptions() {
	// my mental model:
	// - to build the tree, i have to know the root node first
	// - to know which node is the root, maybe i can loop over the input and convert it to a hash map?
	// something like:
	/*
		[20] -> [15,17]
		[50] -> [20,80]
		[80] -> [19, -1]

		... hmm wait it's not right. let's try other approach

		- maybe i'll just lopp over the input and build the tree along the way, is it possiblle?

		- create empty hashmap of treenode pointer
		- create a variable `root` to store the root candidate
		[20,15,1] -> create a TreeNode{ val: 20, left: TreeNode { val: 15}}
		          -> put 20 and 15 on hash map
				  -> map[20] = Tree created TreeNode
				  -> map[15] -> the left of 20
				  -> root: 20
		[20,17,0] -> find 20 from the unfinished tree from previous node
		          -> if found in hashmap, put 17 on the right child
				  -> map[20].Right = TreeNode{ val: 17}
				  -> map[17] -> 20's right
				  -> 20 found in hash map, no need to chaneg the root
		[50,20,1] -> 50 not found in hashmap, is 20 (chlld) there? found it
		          -> create TreeNode { val: 50, left = map[20] }
		          -> apppend new node 50 to map -> map[50] = the just created node for 50
				  -> 50 not found on hash map, earlier, change root to 50
		[50,80,0] -> 50 found on hash map -> append a node 80 to its right
		          -> map[50].right = Treenode { val: 80}
				  -> map[80] -> 50's right
				  -> no need to change `root`
		[80,19,1] -> 80 found on the hash map
		          -> map[80].left = 19
				  -> map[19] = 80's left
				  -> no need to change root because 80 can beb found on hashmap earlier
		end of loop -> return root
		will this explode ina huge inputs? don't know, let's try first

	*/
	/*
	 commentary after submission: there's a lost changes from my former mental model
	 like addition of rootCandidates array, children hashmap, etc
	 most time spent on fixing failed testcases
	*/

	// descriptions := [][]int{
	// 	{20, 15, 1},
	// 	{20, 17, 0},
	// 	{50, 20, 1},
	// 	{50, 80, 0},
	// 	{80, 19, 1},
	// }
	// descriptions := [][]int{
	// 	{85, 82, 1},
	// 	{74, 85, 1},
	// 	{39, 70, 0},
	// 	{82, 38, 1},
	// 	{74, 39, 0},
	// 	{39, 13, 1},
	// }
	descriptions := [][]int{
		{85, 74, 0},
		{38, 82, 0},
		{39, 70, 0},
		{82, 85, 0},
		{74, 13, 0},
		{13, 39, 0},
	}
	out := createBinaryTree(descriptions)
	lib.PrintTreeASCII(out)
}
