# LeetCode's Official Editorial

### Approach: Dynamic Programming

#### Intuition

**Hint 1**

After any player makes a move, the total value of the remaining stones does not change.

In each move, a player can merge the leftmost x stones and gain points equal to their total value. Therefore, the value of the leftmost stone after each move must be equal to some prefix sum of stones. In other words, the points obtained by a player in each round must be equal to a prefix sum of the initial array stones.

**Hint 2**

Based on Hint 1, we can transform the original game into the following equivalent form:

- First, calculate the prefix sum array pre of stones, where

pre\[i\]\=∑j\=0i​stones\[j\].

- Alice and Bob take turns choosing indices in pre, with Alice going first. In each move, the current player chooses an index u and gains pre\[u\] points.

- If the current player is Alice and this is her first move, then u cannot be 0. This corresponds to the original rule that the number of stones x chosen in the first move must be greater than 1.

- For every subsequent move, if the opponent chose index v in the previous move, then the current player must choose an index u\>v. In the original game, this corresponds to the opponent merging some stones and making the value of the leftmost remaining stone equal to pre\[v\]. The current player then merges all stones from that leftmost stone through the stone at index u, making the value of the new leftmost stone equal to pre\[u\].

- Let the length of stones and pre be n. If a player chooses index n−1, they merge all remaining stones, so the game ends.

**Intuition**

We can use dynamic programming based on game theory to solve the transformed game.

Let f\[i\] denote the maximum score difference between the current player and the opponent when the current player can choose an index from the range \[i,n).

When computing f\[i\], there are two possibilities:

- **The current player does not choose i.** 
 In this case, they can simply choose an index from \[i+1,n), so the state transition is

 f\[i\]\=f\[i+1\].

- **The current player chooses i.** 
 They gain pre\[i\] points, and the opponent then plays optimally in the remaining range \[i+1,n). From the current player's perspective, the opponent can achieve a score difference of f\[i+1\]. Therefore, the current player's resulting score difference is

 f\[i\]\=pre\[i\]−f\[i+1\].

Since the current player plays optimally, they choose the better of these two options:

f\[i\]\=max(f\[i+1\],pre\[i\]−f\[i+1\]).

We calculate the states in reverse order, starting from i\=n−1. The final answer is f\[1\] because Alice's first move cannot choose index 0.

**Details**

When i\=n−1, there is no state f\[i+1\]. This is a boundary case because choosing index n−1 means merging all remaining stones and ending the game. Therefore,

f\[n−1\]\=pre\[n−1\].

We can then calculate all remaining states in reverse order, starting from i\=n−2.

**Thinking**

Why is the final answer f\[1\] instead of f\[0\]?

The reason is that Alice cannot choose index 0 on her first move. Therefore, the first valid index she can choose is 1, making f\[1\] the required answer.

#### Complexity Analysis

- Time complexity: O(n).

- Space complexity: O(n).

 The arrays pre and f each require O(n) space.
