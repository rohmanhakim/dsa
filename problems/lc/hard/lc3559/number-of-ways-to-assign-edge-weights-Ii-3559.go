/*
There is an undirected tree with n nodes labeled from 1 to n, rooted at node 1. The tree is represented by a 2D integer array edges of length n - 1, where edges[i] = [ui, vi] indicates that there is an edge between nodes ui and vi.

Initially, all edges have a weight of 0. You must assign each edge a weight of either 1 or 2.

The cost of a path between any two nodes u and v is the total weight of all edges in the path connecting them.

You are given a 2D integer array queries. For each queries[i] = [ui, vi], determine the number of ways to assign weights to edges in the path such that the cost of the path between ui and vi is odd.

Return an array answer, where answer[i] is the number of valid assignments for queries[i].

Since the answer may be large, apply modulo 10^9 + 7 to each answer[i].

Note: For each query, disregard all edges not in the path between node ui and vi.

Example 1:

Input: edges = [[1,2]], queries = [[1,1],[1,2]]

Output: [0,1]

Explanation:

	Query [1,1]: The path from Node 1 to itself consists of no edges, so the cost is 0. Thus, the number of valid assignments is 0.
	Query [1,2]: The path from Node 1 to Node 2 consists of one edge (1 → 2). Assigning weight 1 makes the cost odd, while 2 makes it even. Thus, the number of valid assignments is 1.

Example 2:

Input: edges = [[1,2],[1,3],[3,4],[3,5]], queries = [[1,4],[3,4],[2,5]]

Output: [2, 1, 4]

Explanation:

	Query [1,4]: The path from Node 1 to Node 4 consists of two edges (1 → 3 and 3 → 4). Assigning weights (1,2) or (2,1) results in an odd cost. Thus, the number of valid assignments is 2.
	Query [3,4]: The path from Node 3 to Node 4 consists of one edge (3 → 4). Assigning weight 1 makes the cost odd, while 2 makes it even. Thus, the number of valid assignments is 1.
	Query [2,5]: The path from Node 2 to Node 5 consists of three edges (2 → 1, 1 → 3, and 3 → 5). Assigning (1,2,2), (2,1,2), (2,2,1), or (1,1,1) makes the cost odd. Thus, the number of valid assignments is 4.

Constraints:

	2 <= n <= 10^5
	edges.length == n - 1
	edges[i] == [ui, vi]
	1 <= queries.length <= 10^5
	queries[i] == [ui, vi]
	1 <= ui, vi <= n
	edges represents a valid tree.
*/

package lc3559

import "fmt"

const MOD_10_9_7 = 1000000007

/*
=============================================================================
ALGORITHM OVERVIEW
=============================================================================

The key insight is purely combinatorial:

Given a path of L edges where each edge weight is 1 or 2:
- Odd + Odd = Even, Odd + Even = Odd, Even + Even = Even
- The total cost is odd iff an ODD number of edges are assigned weight 1
- The number of subsets of {1..L} with odd cardinality is exactly 2^(L-1)

Proof sketch: Each subset either has odd or even size. There's a bijection
between odd and even subsets (flip inclusion of element 1), so exactly half
of the 2^L subsets have odd size → 2^(L-1).

So for each query (u, v):
  1. Find the number of edges L on the path from u to v
  2. Answer = 2^(L-1) mod (10^9 + 7)

The challenge is computing L efficiently for up to 10^5 queries on a tree
with up to 10^5 nodes.

SOLUTION PIPELINE:
  1. Build adjacency list (O(n))
  2. Precompute depth[] and binary lifting table up[][] (O(n log n))
  3. For each query: find LCA in O(log n), compute distance, apply formula
  Total: O(n log n + q log n)
=============================================================================
*/

func binaryLiftingPreprocess(tree map[int][]int, log int, root int, up [][]int, depth []int) {
	/*
		===========================================================================
		BINARY LIFTING PREPROCESSING
		===========================================================================

		Binary lifting is a technique that precomputes "jump tables" so we can
		find the k-th ancestor of any node in O(log n) time.

		up[v][k] = the 2^k-th ancestor of node v

		Examples for a chain 1-2-3-4-5 (rooted at 1):
		  up[5][0] = 4      (2^0 = 1st ancestor = parent)
		  up[5][1] = 3      (2^1 = 2nd ancestor)
		  up[5][2] = 1      (2^2 = 4th ancestor)
		  up[5][3] = 1      (2^3 = 8th ancestor, clamps to root)

		RECURSIVE FORMULA:
		  up[v][0] = parent(v)                          [base case]
		  up[v][k] = up[ up[v][k-1] ][k-1]             [jump 2^(k-1) twice = 2^k]

		Why it works: To jump 2^k nodes, jump 2^(k-1) to reach a midpoint,
		then jump another 2^(k-1) from there.

		TIME:  O(n log n) — each of n nodes fills log columns
		SPACE: O(n log n) — the up[][] table
		===========================================================================
	*/

	// Iterative DFS using an explicit stack (avoids recursion stack overflow
	// for deep/chain trees where n = 10^5).
	stack := []int{root}
	up[root][0] = root // Root's parent is itself (sentinel: climbing above
	// root keeps you at root, which is correct for LCA)
	depth[root] = 0

	for len(stack) > 0 {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Fill the binary lifting table for node v.
		// INVARIANT: When v is popped from the stack, its parent up[v][0]
		// has already been fully processed (all up[parent][k] are filled).
		// This holds because in DFS, parent is always popped before child.
		for k := 1; k < log; k++ {
			// up[v][k-1] = midpoint ancestor (2^(k-1) steps up)
			// up[midpoint][k-1] = another 2^(k-1) steps up from midpoint
			// Together: 2^(k-1) + 2^(k-1) = 2^k steps up from v
			up[v][k] = up[up[v][k-1]][k-1]
		}

		// Push children onto the stack for processing.
		for _, child := range tree[v] {
			if child != up[v][0] { // Skip the parent (undirected tree has back-edges)
				up[child][0] = v            // Set child's parent BEFORE pushing
				depth[child] = depth[v] + 1 // Depth = parent's depth + 1
				stack = append(stack, child)
			}
		}
	}
}

func lca(up [][]int, depth []int, log int, u, v int) int {
	/*
		===========================================================================
		LOWEST COMMON ANCESTOR (LCA) via Binary Lifting
		===========================================================================

		The LCA of two nodes u and v in a rooted tree is the deepest node
		that is an ancestor of both u and v.

		ALGORITHM (two phases):

		Phase 1 — Equalize depths:
		  If depth[u] > depth[v], lift u up by (depth[u] - depth[v]) using
		  kthAncestor. Now both are at the same depth.

		Phase 2 — Binary lift both upward:
		  Try lifting both u and v by decreasing powers of 2 (log-1 down to 0).
		  If up[u][k] != up[v][k], both nodes jump up by 2^k. This moves them
		  as high as possible WITHOUT passing the LCA.
		  After this loop, u and v are direct children of the LCA.
		  Return up[u][0] (the LCA).

		EXAMPLE: Find LCA(4, 5) in tree 1-2-3-4 and 1-2-3-5
		  Phase 1: depth[4]=3, depth[5]=3 → already equal
		  Phase 2: up[4][1]=2, up[5][1]=2 → equal, skip
		          up[4][0]=3, up[5][0]=3 → equal, skip
		  Return up[4][0] = 3 ✓

		TIME: O(log n) per query
		===========================================================================
	*/

	// Phase 1: Bring u to the same depth as v (u is the deeper one after swap)
	if depth[u] < depth[v] {
		u, v = v, u
	}
	u = kthAncestor(up, log, u, depth[u]-depth[v]) // Lift u up by the depth difference

	if u == v {
		return u // v was already an ancestor of (the original) u
	}

	// Phase 2: Binary lift both until they're just below the LCA
	for k := log - 1; k >= 0; k-- {
		// If their 2^k-th ancestors are different, both can safely jump up.
		// We use != (not <) because this is an unrooted tree — nodes don't
		// have numeric ordering, we only care about identity.
		if up[u][k] != up[v][k] {
			u = up[u][k]
			v = up[v][k]
		}
	}
	// Now u and v are children of the LCA, so the LCA is one step above
	return up[u][0]
}

func kthAncestor(up [][]int, log int, v, diff int) int {
	/*
		===========================================================================
		K-TH ANCESTOR via Binary Lifting
		===========================================================================

		Given a node v and a distance diff, find the ancestor of v that is
		`diff` edges above it in the tree.

		ALGORITHM: Decompose `diff` into binary representation and use the
		precomputed jump table.

		If diff = 13 (binary: 1101), then:
		  diff has bits 0, 2, 3 set (1 + 4 + 8 = 13)
		  → Jump up 2^0, then 2^2, then 2^3

		EXAMPLE: kthAncestor(5, 3) in chain 1-2-3-4-5
		  diff = 3 = binary 11
		  bit 0: diff & 1 = 1 → v = up[5][0] = 4
		  bit 1: diff & 2 = 2 → v = up[4][1] = 2
		  Result: 2 is the 3rd ancestor of 5 ✓

		TIME: O(log n) — one iteration per bit
		===========================================================================
	*/
	for bit := 0; bit < log; bit++ {
		// Check if the `bit`-th bit of diff is set
		if diff>>bit&1 == 1 {
			v = up[v][bit] // Jump up by 2^bit
		}
	}
	return v
}

func assignEdgeWeights(edges [][]int, queries [][]int) []int {
	/*
		===========================================================================
		MAIN SOLVER FUNCTION
		===========================================================================
	*/

	// -------------------------------------------------------------------------
	// 1. ON-DEMAND LOG CALCULATION
	// -------------------------------------------------------------------------
	// We need `log = ceil(log2(n))` where n = number of nodes.
	// This determines how many levels of binary lifting we need.
	// For n nodes, the deepest ancestor chain is n-1 edges, so we need
	// log2(n-1)+1 bits to represent any distance up to n-1.
	//
	// Instead of importing math, we compute it by counting how many times
	// we can double 1 before exceeding n:
	//   log=1: 1<<1 = 2   <= n? yes → increment
	//   log=2: 1<<2 = 4   <= n? yes → increment
	//   ...
	//   log=k: 1<<k = 2^k > n?  done → log = k
	//
	// For n=100000: 2^16=65536 <= 100000, 2^17=131072 > 100000 → log=17
	// For n=5:      2^2=4 <= 5, 2^3=8 > 5 → log=3
	// -------------------------------------------------------------------------
	n := len(edges) + 1
	log := 1
	for (1 << log) <= n {
		log++
	}

	// -------------------------------------------------------------------------
	// 2. ADJACENCY LIST — Undirected Tree Representation
	// -------------------------------------------------------------------------
	// Trees are stored as adjacency lists: each node maps to a list of its
	// neighbors. Since the tree is undirected, every edge [u, v] creates
	// entries in both directions: u→v and v→u.
	//
	// For edges = [[1,2],[1,3],[3,4],[3,5]]:
	//   tree[1] = [2, 3]
	//   tree[2] = [1]
	//   tree[3] = [1, 4, 5]
	//   tree[4] = [3]
	//   tree[5] = [3]
	//
	// TIME:  O(n) — process n-1 edges, each added twice
	// SPACE: O(n)
	// -------------------------------------------------------------------------
	tree := make(map[int][]int, n)

	for i := 0; i < len(edges); i++ {
		tree[edges[i][0]] = append(tree[edges[i][0]], edges[i][1])
		tree[edges[i][1]] = append(tree[edges[i][1]], edges[i][0])
	}

	// -------------------------------------------------------------------------
	// 3. BINARY LIFTING TABLE ALLOCATION
	// -------------------------------------------------------------------------
	// up[v][k] = 2^k-th ancestor of node v
	//
	// We allocate (n+1) rows (0 unused, nodes are 1-indexed) and `log` columns.
	// All values initialized to 0 by Go's make(). The preprocessing fills them.
	//
	// SPACE: O(n log n) — for n=100000 and log=17: ~1.7M integers (~13 MB)
	// -------------------------------------------------------------------------
	up := make([][]int, n+1)
	for i := range up {
		up[i] = make([]int, log)
	}
	depth := make([]int, n+1) // depth[v] = number of edges from root to v

	// Preprocess: single DFS from root=1 to fill depth[] and up[][]
	// TIME: O(n log n)
	binaryLiftingPreprocess(tree, log, 1, up, depth)

	// -------------------------------------------------------------------------
	// 4. ANSWER QUERIES
	// -------------------------------------------------------------------------
	// For each query (u, v):
	//   a. Find LCA(u, v) — O(log n)
	//   b. Compute path length L = depth[u] + depth[v] - 2*depth[LCA(u,v)]
	//   c. Answer = 2^(L-1) mod MOD
	// -------------------------------------------------------------------------
	out := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		u := queries[i][0]
		v := queries[i][1]
		if u == v {
			out[i] = 0 // No edges on the path → cost is 0 (even) → 0 valid assignments
			continue
		}

		// -------------------------------------------------------------------------
		// DISTANCE FORMULA ON TREES
		// -------------------------------------------------------------------------
		// The path from u to v goes: u → LCA(u,v) → v
		//   - Edges from u up to LCA: depth[u] - depth[LCA]
		//   - Edges from v up to LCA: depth[v] - depth[LCA]
		//   - Total edges L = depth[u] + depth[v] - 2*depth[LCA]
		//
		// This works because in a tree, there is exactly ONE path between
		// any two nodes. The LCA is the "pivot point" where the paths from
		// root to u and root to v diverge.
		// -------------------------------------------------------------------------
		distUV := depth[u] + depth[v] - 2*depth[lca(up, depth, log, u, v)]

		// -------------------------------------------------------------------------
		// COMBINATORICS FORMULA: 2^(L-1)
		// -------------------------------------------------------------------------
		// Each of the L edges can be assigned weight 1 or 2.
		// Total assignments = 2^L.
		//
		// We need the SUM to be ODD. Since 1 is odd and 2 is even:
		//   sum is odd ⟺ an odd number of edges get weight 1
		//
		// Out of all 2^L subsets of edges (choosing which get weight 1):
		//   - Exactly half have odd cardinality = 2^(L-1)
		//   - Exactly half have even cardinality = 2^(L-1)
		//
		// WHY? There's a bijection: given any subset S, toggle whether
		// edge 1 is included → S ⊕ {1}. This flips odd↔even cardinality,
		// so odd subsets and even subsets are equal in number.
		//
		// Edge case: L=0 (u==v) → 2^(-1) would be fractional, but we
		// handle u==v separately by returning 0.
		// -------------------------------------------------------------------------
		out[i] = modpow(2, distUV-1, MOD_10_9_7)
	}

	return out
}

func modpow(base, exp, mod int) int {
	/*
		===========================================================================
		MODULAR EXPONENTIATION (Binary Exponentiation / Exponentiation by Squaring)
		===========================================================================

		Computes base^exp mod mod efficiently in O(log exp) multiplications.

		ALGORITHM: Exploit the binary representation of exp.
		  If exp = 13 (binary 1101):
		    base^13 = base^1 * base^4 * base^8
		    (bits 0, 2, 3 are set)

		  We square the base at each step and multiply into result
		  when the corresponding bit is set.

		EXAMPLE: 2^13 mod 1000000007
		  Step 0: bit 0 set → result = 1 * 2 = 2,    base = 2^2 = 4
		  Step 1: bit 1 clear → result = 2,           base = 4^2 = 16
		  Step 2: bit 2 set → result = 2 * 16 = 32,   base = 16^2 = 256
		  Step 3: bit 3 set → result = 32 * 256 = 8192, base = 256^2 = ...

		Why O(log exp): exp halves (via right shift) each iteration,
		so there are at most log2(exp) iterations.

		Why modular: Each multiplication is taken mod `mod` to prevent
		integer overflow and keep numbers manageable. This is valid because
		(a * b) mod m = ((a mod m) * (b mod m)) mod m.
		===========================================================================
	*/
	result := 1
	base %= mod // Reduce base first in case it's >= mod
	for exp > 0 {
		if exp%2 == 1 { // If current bit of exp is set
			result = result * base % mod // Multiply result by current base power
		}
		base = base * base % mod // Square the base for the next bit position
		exp >>= 1                // Shift exp right (move to next bit)
	}
	return result
}

func NumberOfWaysToAssignEdgeWeightsII() {
	edges := [][]int{
		{1, 2},
		{1, 3},
		{3, 4},
		{3, 5},
	}
	queries := [][]int{
		{1, 4},
		{3, 4},
		{2, 5},
	}
	// edges := [][]int{
	// 	{1, 2},
	// }
	// queries := [][]int{
	// 	{1, 1},
	// 	{1, 2},
	// }
	out := assignEdgeWeights(edges, queries)

	fmt.Println(out)
}
