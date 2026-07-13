# Binary Search Tree (BST) Properties

## When to Use

The input is a BST (Binary Search Tree), or the problem can be modeled on BST properties. Key signals:
- Input is explicitly a BST
- Need sorted order from tree data
- Need to validate tree structure
- Need to search/insert efficiently in a tree

## Core Properties

1. **Left subtree < Root < Right subtree** — for every node
2. **In-order traversal = sorted order** — this is the most powerful property
3. **Search/Insert in O(h)** — go left or right based on comparison with current node

## Key Templates

### In-Order Traversal (Sorted Access)

Use when you need elements in sorted order, or need to compare consecutive elements:

```go
var traverse func(curr *TreeNode)
traverse = func(curr *TreeNode) {
    if curr == nil {
        return
    }
    traverse(curr.Left)
    // visit curr (process in sorted order)
    traverse(curr.Right)
}
traverse(root)
```

**What in-order gives you for free:**
- Adjacent nodes in traversal are adjacent in sorted order
- Minimum absolute difference between any two nodes must be between consecutive in-order nodes (LC 530)
- Check BST validity by verifying strictly increasing sequence (LC 98)
- Find kth smallest by counting during traversal (LC 230)

### BST Search

Use when you need to find/verify a specific value:

```go
func search(root *TreeNode, val int) *TreeNode {
    if root == nil || root.Val == val {
        return root
    }
    if val < root.Val {
        return search(root.Left, val)
    }
    return search(root.Right, val)
}
```

**Key insight:** You only traverse one branch at each level → O(h) time.

### BST Validation

Use when verifying tree structure:

```go
func validate(node *TreeNode, lo, hi int) bool {
    if node == nil {
        return true
    }
    if node.Val <= lo || node.Val >= hi {
        return false
    }
    return validate(node.Left, lo, node.Val) && validate(node.Right, node.Val, hi)
}
```

**Key insight:** Pass min/max bounds down. Each node must be within (lo, hi).

## BST vs General Tree

| Aspect | BST | General Tree |
|---|---|---|
| Sorted property | Yes — in-order = sorted | No |
| Search | O(h) via comparison | O(n) must check all nodes |
| Min/Max | Leftmost/Rightmost node | Must traverse all |
| In-order use | Sorted access, consecutive comparison | Not meaningful for ordering |

## Example: LC 530 (Minimum Absolute Difference in BST)

**Problem:** Find min absolute difference between any two nodes in a BST.

**Key insight:** In a BST, in-order traversal gives sorted order. The minimum difference is always between two consecutive elements in sorted order — no need to check all pairs.

```
Tree:    4
        / \
       2   6
      / \
     1   3

In-order: 1, 2, 3, 4, 6
Consecutive diffs: 1, 1, 1, 2
Min = 1 ✓
```

**Algorithm:**
1. In-order traversal with `prevVal` tracking
2. At each node: `diff = curr.Val - prevVal` (always positive since sorted)
3. Track minimum diff across all consecutive pairs

**Complexity:** O(n) time, O(h) space (call stack)

## Common Mistakes

1. **Trying all pairs** — O(n²) when in-order gives O(n). The BST sorted property eliminates the need for brute-force.
2. **Using DFS without BST property** — just counting depth (LC 104) is DFS, not BST. BST is about leveraging the sorted invariant.
3. **Forgetting that in-order = sorted** — this is the single most important BST property. Internalize it.
4. **Sentinel value issues** — when using a prev-value tracker, ensure the sentinel (e.g., -1) is outside the valid value range.

## Classic Problems to Practice

| Problem | Key Idea |
|---|---|
| **LC 530** — Minimum Absolute Difference in BST | In-order → compare consecutive for min diff |
| **LC 98** — Validate Binary Search Tree | In-order → check strictly increasing, or use bounds propagation |
| **LC 230** — Kth Smallest Element in BST | In-order → count to k |
| **LC 700** — Search in BST | BST search property: go left or right |
| **LC 701** — Insert into BST | BST insert property: find the leaf position |
| **LC 235** — LCA of a BST | BST property narrows search: both nodes in same subtree? |
| **LC 501** — Find Mode in BST | In-order → count frequencies |
| **LC 669** — Trim a BST | BST property guides recursive trimming |