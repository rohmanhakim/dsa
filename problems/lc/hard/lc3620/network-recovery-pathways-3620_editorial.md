# LeetCode's Official Editorial

### Approach 1: Binary Answer + Shortest Path (Dijkstra)

#### Intuition

Let's first summarize the key observations from the problem:

1. We are given a directed graph, where the online status of each node is determined by the array online.
2. An edge is considered **valid** only if both of its endpoints are online.
3. We need to find a path from node 0 to node n−1 whose total edge weight does not exceed k, and every node on the path must be online.
4. The **score** of a path is defined as the minimum edge weight along that path. Our goal is to maximize this score.
5. If no valid path exists, return −1.

Problems that ask us to **maximize the minimum value** often exhibit a monotonic property, making binary search a natural solution.

Suppose there exists a path from node 0 to node n−1 such that

1. its total weight is at most k, and
2. its minimum edge weight is at least x.

Then, for any threshold y≤x, the same path also satisfies

1. its total weight is still at most k, and
2. every edge weight is at least y.

Therefore,

- if check(x) is feasible, then check(y) is also feasible for every y≤x;
- if check(x) is infeasible, then check(z) is also infeasible for every z\>x.

This monotonicity allows us to binary search the answer.

For each candidate value mid, we check whether a valid path exists under the following restriction:

1. Only edges whose weights are at least mid may be used.
2. The total weight of the path must not exceed k.

If check(mid) returns `true`, then a path exists whose minimum edge weight is at least mid, so we try a larger threshold.

Otherwise, no such path exists, and we decrease the threshold.

The remaining question is how to implement `check()`. In this approach, we use **Dijkstra's algorithm**.

The procedure is straightforward:

1. Build the graph using only edges whose endpoints are both online.
2. For each candidate threshold mid:
 - Ignore every edge whose weight is smaller than mid.
 - Run Dijkstra's algorithm to compute the shortest path from node 0 to node n−1.
 - If the shortest distance is at most k, then mid is feasible.

#### Implementation

#### Complexity Analysis

Let V be the number of nodes, E be the number of edges, and U be the maximum edge weight.

- Time complexity: O((E+V)logVlogU)
 
 Constructing the graph takes O(E) time. Binary search performs O(logU) iterations, and each iteration runs Dijkstra's algorithm in O((V+E)logV) time.
 
- Space complexity: O(V+E).
 
 The adjacency list requires O(V+E) space, while the temporary arrays used in `check()` require O(V) space.
 

---

### Approach 2: Binary Answer + Memoization Search

#### Intuition

In this approach, we implement `check()` using memoized depth-first search. As before, we first construct the graph using only online nodes.

Define dfs(u) as the minimum total path weight from node u to node n−1, considering only edges whose weights are at least mid.

For every outgoing edge (u,v) with weight w≥mid, we have the transition

dfs(u)\=minv​(dfs(v)+w).

The recursion starts from node 0, while the base case is 
dfs(n−1)\=0.

To avoid repeated computation, we use a memoization array, where `memo[u]` stores the minimum total path weight from node u to node n−1.

Finally, if dfs(0)≤k, then a valid path exists under the current threshold. Otherwise, no valid path exists.

#### Implementation

#### Complexity Analysis

Let V be the number of nodes, E the number of edges, and U the maximum edge weight.

- Time complexity: O((V+E)logU).
 
- Space complexity: O(V+E).
 

---

### Approach 3: Binary Answer + Topological Sorting + Dynamic Programming

#### Intuition

Instead of memoized DFS, we can compute the minimum path weight using **topological sorting** combined with dynamic programming.

Before processing the graph, we first remove nodes that are unreachable from node 0.

Let `dp[u]` denote the minimum total path weight from node 0 to node u.

For every edge u→v whose weight satisfies w≥mid, the transition is dp\[v\]\=min(dp\[v\],dp\[u\]+w).

The topological order lets us process vertices from node 0 toward node n−1, updating each state exactly once.

Conceptually, this approach is equivalent to the memoized DFS in Approach 2. Memoized DFS computes the states recursively from the destination back to the source, while topological DP computes them iteratively from the source to the destination.

#### Implementation

#### Complexity Analysis

Let V be the number of nodes, E the number of edges, and U be the maximum edge weight.

- Time complexity: O((V+E)logU).
 
- Space complexity: O(V+E).
 
