# LeetCode's Official Editorial

### Approach 1: Simulation

#### Intuition

From the problem statement, we know that stationary vehicles do not count toward collision times; only moving vehicles that collide are counted. We traverse all vehicles from left to right and use a variable flag to record the status of the vehicles on the left.

- If there are no vehicles on the left side or all vehicles on the left side are moving left, then flag is set to −1.
- If a collision occurs on the left and the vehicles eventually stop, then flag is set to 0.
- If there are consecutive vehicles on the left moving to the right, then flag stores the number of such vehicles.

In this case, we can divide the behavior of the current vehicle into the following three scenarios:

1. The current vehicle moves left. If flag≥0, the collision count increases by flag+1, and flag is set to 0.
2. The current vehicle is stationary. If flag\>0, the collision count increases by flag, and flag is set to 0.
3. The current vehicle moves right. If flag<0, set flag to 1; otherwise, increment flag by 1.

Finally, return the total number of collisions.

#### Complexity Analysis

Let n be the length of directions.

- Time complexity: O(n).

- Space complexity: O(1).

 Only a constant number of variables were used in the process.

### Approach 2: Counting

#### Intuition

We define a vehicle that continuously moves outward (either left or right) without being blocked by a reversing or stationary vehicle in between as an "outward-moving vehicle."

Left-moving outward vehicles and right-moving outward vehicles will not collide with each other. All other vehicles will collide exactly once.

#### Complexity Analysis

Let n be the length of directions.

- Time complexity: O(n).

- Space complexity: O(1).

 Only a constant number of variables were used in the process.
