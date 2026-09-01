# LeetCode's Official Editorial

### Approach: Breadth-First Search

#### Intuition

We can use Breadth-First Search (BFS) to calculate the shortest path. In addition to the basic position coordinates x and y, our BFS state must also track the current energy e, the number of steps taken steps, and the set of already collected litter. We can efficiently represent the collected litter using a bitmask, where each piece of litter is assigned a unique bit identifier starting from 0.

During the search process, move in four directions from the current state (x,y,mask,e,steps), and update the state as follows:

- Current position: (nx,ny).

- Collected litter mask: If the current position is ’L’, add the identifier of this piece of litter to mask; otherwise, the mask remains unchanged.

- If the current position is ’R’, update the energy of the current position to energy; otherwise, update it to e−1.

- Update the steps of the current position to one more than the previous step.

The starting point of our BFS is (sx,sy,0,energy,0), where (sx,sy) is the student's starting position. The target endpoint is any valid state where the mask equals fullMask (representing the collection of all litter).

During the BFS, we use a 3D array bestEnergy\[x\]\[y\]\[mask\] to record the maximum energy value reached at position (x,y) for each possible mask state. We only enqueue a new state and update bestEnergy\[x\]\[y\]\[mask\] if the remaining energy is strictly greater than the previously recorded energy for that configuration. This pruning prevents redundant cycles, such as repeatedly jumping back and forth between two positions.

#### Complexity Analysis

Let m and n be the number of rows and columns of the matrix respectively, energy be the total energy, and k be the number of occurrences of ’L’ in the matrix.

- Time complexity: O(mn⋅energy⋅2k).

 The state is determined by position and a binary mask, and the corresponding energy bestEnergy\[x\]\[y\]\[mask\] for the same position and mask may take up to energy different values.

- Space complexity: O(mn⋅energy⋅2k).

 This is the space required by the queue.
