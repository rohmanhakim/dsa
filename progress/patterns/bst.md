# BST

**Status:** ⏳ 0/5 sessions toward mastery (guided).
**Sessions:** LC 530 (guided, sev 3), LC 669 (guided, sev 3).

## Variants

### In-order traversal (sorted access) — LC 530
In-order traversal gives sorted values; minimum absolute difference is always between consecutive in-order nodes. Initially attempted bottom-up DFS without leveraging BST property.

### Recursive pruning — LC 669
When trimming/filtering a BST, use the search property at each node: `val < low` → discard root + left subtree, recurse right; `val > high` → discard root + right subtree, recurse left; in range → keep root, trim both children. Initially attempted complex parent-pointer rewiring; key insight is that BST invariant makes parent tracking unnecessary — a single comparison determines which subtrees to skip entirely.

## Decision Framework

"Am I **accessing sorted order** (in-order traversal) or **pruning/searching** (use search property to make local decisions)?"

## Lessons

- **LC 530:** BST + in-order = sorted order — eliminates brute-force pair comparison. Must internalize: when input is a BST, always ask "what does in-order give me for free?"
- **LC 669:** Over-engineered with parent-pointer DFS when the BST search property enables a simple recursive return pattern. Ask "what does the input's structure give me for free?" before building mechanics.