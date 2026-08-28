# LettCode Official Editorial

### Approach: Sequential Enumeration

#### Intuition

The task requires rearranging the string s to find the lexicographically smallest palindrome that is greater than target. Therefore, we first need to determine whether s can be rearranged to form a palindrome.

A string can be rearranged into a palindrome if and only if at most one character appears an odd number of times. Otherwise, it is impossible to form a palindrome. Therefore, we first count the frequency of each character in s. If more than one character has an odd frequency, we return an empty string.

We then attempt to construct the lexicographically smallest palindrome that is greater than target. A palindrome consists of the **left half + an optional middle character + the right half**. Therefore, we only need to construct the left half, whose length is ⌊2n​⌋, because the right half is uniquely determined by the left half. To simplify the construction, we divide the frequency of each character by 2, which gives us the number of available copies for the left half.

The construction consists of the following steps:

- **Greedy Construction of the Left Half:** According to the problem, we need to find the lexicographically smallest palindrome that is greater than target. Therefore, we construct the left half from left to right.

 - At each position, we try to place the smallest currently available character, from `'a'` to `'z'`. After placing it, we **check** whether the current prefix can eventually be extended to form a palindrome greater than target.

 - **Checking method:** We combine the current prefix with the remaining characters arranged in descending order to construct the lexicographically largest palindrome that can be formed under the current prefix. If this palindrome is greater than target, then the current character is feasible.

 - If the condition is satisfied, we fix the current character and continue to the next position. Otherwise, we restore its count and try the next character. If none of the available characters can produce a palindrome greater than target, we return an empty string.

 - **Pruning:** If the current prefix is already lexicographically greater than the corresponding prefix of target, then the remaining characters no longer need to be chosen to make the palindrome greater than target. In this case, we arrange all remaining characters in ascending order to obtain the lexicographically smallest palindrome with the current prefix.

- **Final Palindrome Construction:** After constructing the left half, we concatenate the middle character, if n is odd, and the reverse of the left half to obtain the final palindrome.

#### Details

- **Greedy Correctness:** At each position, we choose the smallest character that can still lead to a valid palindrome greater than target. Since the feasibility check guarantees that the remaining characters can be arranged to produce a valid result, choosing the smallest feasible character at each position produces the lexicographically smallest possible palindrome.

- **Maximum Palindrome Check:** By arranging all remaining characters in descending order, we obtain the lexicographically largest palindrome that can be formed under the current prefix. If this palindrome is not greater than target, then no palindrome starting with the current prefix can be greater than target. Therefore, the current prefix is not feasible.

#### Complexity Analysis

Let n be the length of the given string s, and let ∣Σ∣\=26 be the size of the character set.

- Time complexity: O(n∣Σ∣×(n+∣Σ∣)).

 We first traverse s to count the frequency of each character and determine whether it can be rearranged into a palindrome, which takes O(n+∣Σ∣) time. We then construct the left half from left to right and try each possible character, resulting in at most O(n∣Σ∣) attempts. For each attempt, the `check()` function constructs a candidate palindrome in O(n+∣Σ∣) time. Therefore, the overall time complexity is O(n∣Σ∣(n+∣Σ∣)).

- Space complexity: O(n+∣Σ∣).

 The frequency array requires O(∣Σ∣) space, while the intermediate strings used during construction require O(n) space.
