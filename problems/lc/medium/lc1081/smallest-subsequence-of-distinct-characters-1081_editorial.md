# LeetCode's Official Editorial

### Approach: Greedy + Monotonic Stack

#### Intuition

First, consider a simpler problem: given a string s, how can we remove one character to make the resulting string lexicographically smallest? The answer is to find the smallest index i such that s\[i\]\>s\[i+1\] and remove the character s\[i\]. For convenience, we refer to such a character as a **key character**.

With this observation, we can return to the original problem. A straightforward approach is to repeatedly find a **key character**, remove it, and continue until no more removals are possible. However, this approach repeatedly constructs intermediate strings, making it inefficient.

Instead, we scan the original string from left to right. At each position, we remove as many **key characters** as possible. Suppose that all **key characters** before position i have already been removed. When processing the current character s\[i\], any newly formed **key character** can only be the character immediately before s\[i\], since all earlier positions have already been processed.

To implement this efficiently, we use a monotonic stack to maintain the current resulting string. The stack is maintained in **non-decreasing** order from bottom to top. When processing the current character s\[i\], if the top of the stack is greater than s\[i\], then the top character becomes a **key character** and should be removed. After it is removed, the new top of the stack becomes adjacent to s\[i\]‘,sowecontinuecomparingthem.Thisprocessrepeatsuntilthestackbecomesemptyoritstopisnolongergreaterthans\[i\]$.

However, we have not yet considered another requirement of the problem: every distinct character in the original string must appear **exactly once** in the final string. To satisfy this requirement, we make the following two modifications.

- If the current character s\[i\] is already in the stack, we skip it. To support this, we maintain an array indicating whether each character is currently in the stack.

- When attempting to pop the top character from the stack, we must ensure that it appears again later in the string. Otherwise, removing it would cause it to disappear from the final answer. Therefore, we maintain the number of remaining occurrences of each character. A character can be popped only if its remaining count is greater than zero.


#### Complexity Analysis

Let N be the length of the string, and let Σ denote the character set. In this problem, all characters are lowercase English letters, so ∣Σ∣\=26.

- Time complexity: O(N).

 Although the code contains a nested loop, each character is pushed onto and popped from the stack at most once.

- Space complexity: O(∣Σ∣).

 Since each character can appear in the stack at most once, the stack contains at most ∣Σ∣ characters. Additionally, two arrays of size ∣Σ∣ are used to record whether each character is in the stack and the number of its remaining occurrences.
