# LeetCode's Official Editorial

### Approach: Guess + Mathematical Induction Verification

#### Intuition

**Hint 1**

If the number of question marks is **odd**, then Alice is definitely the winner.

**Hint 1 Explanation**

Since Alice goes first, the last question mark must be filled by Alice.

It is clear that, in the range \[0,9\], there is at most one digit d that can make the sum of the digits in the first half equal to the sum of the digits in the second half. Therefore, Alice can simply replace the last question mark with any digit other than d.

Thus, we only need to consider the case where the number of question marks is **even**.

**Hint 2**

If the number of question marks is 0, then Bob wins if and only if the sum of the digits in the first half equals the sum of the digits in the second half.

**Hint 3**

If the number of question marks is 2 and they appear on different sides, that is, one question mark is in the first half and the other is in the second half, then Bob wins if and only if the sum of the known digits in the first half equals the sum of the known digits in the second half.

**Hint 3 Explanation**

If the sum of the known digits in the first half is equal to that in the second half, Alice can choose either question mark and replace it with any digit d. Bob can then replace the other question mark with the same digit d. Therefore, Bob is guaranteed to win.

If the sums of the known digits in the two halves are not equal, Alice can choose the question mark in the half with the larger sum and replace it with 9. Since no digit greater than 9 can be chosen, Alice is guaranteed to win.

**Hint 4**

If the number of question marks is 2 and both appear on the same side, Bob wins if and only if the sum of the known digits on that side is exactly 9 less than the sum of the known digits on the other side.

**Hint 4 Explanation**

Bob can always ensure that the digits chosen for two consecutive question marks sum to 9. Specifically, if Alice replaces one question mark with a digit d, Bob can replace the other question mark with 9−d.

Therefore, if the above condition is satisfied, Bob is guaranteed to win. Otherwise:

- If the difference exceeds 9, Alice replaces the question mark with 0.

- If the difference does not exceed 9, Alice replaces the question mark with 9.

In either case, Bob cannot choose a digit in \[0,9\] that makes the sums of the two halves equal, so Alice is guaranteed to win.

**Hint 5**

Suppose the sum of the known digits in the first half is n0​, the number of question marks in the first half is q0​, the sum of the known digits in the second half is n1​, and the number of question marks in the second half is q1​. If q0​+q1​ is even, then Bob wins if and only if

n\_0 - n\_1 = \\frac{9}{2}(q\_1 - q\_0). \\tag{1}

We can interpret this equation based on Hints 3 and 4. Without loss of generality, assume q0​≤q1​.

- For the q0​ question marks on the two sides, according to Hint 3, Alice and Bob can choose digits so that their contributions to the two halves cancel each other out.

- For the remaining q1​−q0​ question marks in the second half, according to Hint 4, Alice and Bob can pair them up, with each pair contributing a total of 9.

- Therefore, for the sums of the two halves to become equal, the first half must initially have a sum that is 29​(q1​−q0​) larger than the second half.

This gives Equation (1).

**Hint 5 Explanation**

We can prove Hint 5 rigorously using mathematical induction on the total number of question marks q0​+q1​.

- According to Hints 3 and 4, Hint 5 holds when q0​+q1​\=2. In this case, Bob is guaranteed to win if Equation (1) holds; otherwise, Alice is guaranteed to win.

- Assume that Hint 5 holds when q0​+q1​\=k. We now consider the case where q0​+q1​\=k+2. Without loss of generality, assume q0​≤q1​.

 - If Equation (1) already holds, then consider Alice's first move. If Alice chooses a question mark in the first half and replaces it with d, Bob replaces a question mark in the second half with the same digit d. If Alice instead chooses a question mark in the second half and replaces it with d, Bob replaces another question mark in the second half with 9−d. In both cases, the remaining game still satisfies Equation (1).

 Therefore, the game can be reduced to the case of q0​+q1​\=k, where, by the induction hypothesis, Bob is guaranteed to win.

 - If Equation (1) does not hold, then we can write

 n0​−n1​\=29​(q1​−q0​)+δ,

 where δ\=0.

 - If δ\>0, Alice replaces a question mark in the second half with 0.

 If Bob chooses a question mark in the first half and replaces it with d, then

 (n0​+d)−n1​\=29​(q1​−q0​)+(δ+d)\=29​((q1​−1)−(q0​−1)).

 Therefore, Equation (1) still does not hold.

 If Bob instead chooses another question mark in the second half and replaces it with d, then

 n0​−(n1​+d)\=29​((q1​−2)−q0​)+(δ+9−d)\=29​((q1​−2)−q0​).

 Therefore, Equation (1) also does not hold.

 - If δ<0, Alice replaces a question mark in the second half with 9.

 If Bob chooses a question mark in the first half and replaces it with d, then

 (n0​+d)−(n1​+9)\=29​(q1​−q0​)+(δ−9+d)\=29​((q1​−1)−(q0​−1)).

 Therefore, Equation (1) still does not hold.

 If Bob instead chooses another question mark in the second half and replaces it with d, then

 n0​−(n1​+9+d)\=29​((q1​−2)−q0​)+(δ−d)\=29​((q1​−2)−q0​).

 Therefore, Equation (1) also does not hold.

 In either case, the game can be reduced to the case of q0​+q1​\=k, where Equation (1) does not hold. By the induction hypothesis, Alice is guaranteed to win.

Therefore, Hint 5 holds for all even values of q0​+q1​.

Finally, we traverse the string num to calculate n0​, q0​, n1​, and q1​. If q0​+q1​ is odd, Alice is guaranteed to win. Otherwise, the winner is determined by Equation (1).

#### Complexity Analysis

Let n be the length of the string num.

- Time complexity: O(n).

- Space complexity: O(n).

 The implementation above creates two substrings of num, so it uses O(n) additional space. We can reduce the space complexity to O(1) by directly traversing num once and calculating n0​, q0​, n1​, and q1​.
