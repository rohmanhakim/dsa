LeetCode's Official Editorial

### Approach 1: Dijkstra

#### Intuition

The problem asks whether it is possible to start from the origin (0,0) with a positive health value and reach the destination (m−1,n−1). Since entering a cell with value 1 reduces the health by 1, the problem can be viewed as finding the path from the start to the destination with the minimum total cost, where the cost of a path is the sum of the values of all visited cells.

Since all cell values are non-negative, we can apply Dijkstra's algorithm to compute the minimum cost to reach every cell. If the minimum cost to reach the destination is less than the initial health value health, then the destination is reachable; otherwise, it is not.

#### Complexity Analysis

Let m and n be the numbers of rows and columns in the matrix grid, respectively.

- Time complexity: O(mnlog(mn)).
 
 There are O(mn) cells, and each cell is removed from the priority queue at most once. Each priority queue operation takes O(log(mn)) time, giving a total time complexity of O(mnlog(mn)).
 
- Space complexity: O(mn).
 
 The distance array and the priority queue together require O(mn) space.
 

---

### Approach 2: 0-1 BFS

#### Intuition

Since every cell in the grid has a value of either 0 or 1, we can use 0-1 BFS instead of Dijkstra's algorithm.

When traversing an edge with weight 0, we push the corresponding cell to the front of the deque. When traversing an edge with weight 1, we push it to the back. This ensures that cells with smaller distances are always processed first, achieving the same effect as Dijkstra's priority queue while reducing each queue operation to O(1) time.

Like Dijkstra's algorithm, 0-1 BFS guarantees that when a cell is removed from the deque for the first time, its shortest distance has already been determined. Therefore, once the destination cell is dequeued, we can immediately return `true`.

We can also prune the search when inserting a cell into the deque:

- If the cost of reaching a cell is already greater than or equal to health, then no path passing through that cell can satisfy the health requirement, so there is no need to continue exploring from it.

#### Complexity Analysis

Let m and n be the numbers of rows and columns in the matrix grid, respectively.

- Time complexity: O(mn).
 
 Each cell is inserted into the deque at most a constant number of times, and every deque operation takes O(1) time. Therefore, the total time complexity is O(mn).
 
- Space complexity: O(mn).
 
 The distance array and the deque together require O(mn) space.
 