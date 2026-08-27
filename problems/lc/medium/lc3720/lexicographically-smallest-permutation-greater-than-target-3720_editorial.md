# LeetCode's Official Editorial

### Approach 1: Sequential Enumeration

#### Intuition

We enumerate the characters of target from left to right to construct the smallest permutation that is lexicographically greater than target. Since we are looking for a permutation that is lexicographically greater than target, there must be some position i such that:

- The first i characters are identical to those of target.
- The i\-th character is greater than target\[i\].
- All characters after the i\-th position are arranged in the smallest lexicographical order.

Therefore, we traverse the characters from left to right. At each position, we first try to keep the character the same as target\[i\], which is the greedy choice that keeps the resulting string as small as possible. If it is not possible to form a result greater than target with this choice, we try to place a larger character and then arrange the remaining characters in ascending order.

#### Algorithm

1. **Count Character Frequencies**: Traverse s and use a count array cnt to record the number of occurrences of each character, which represents the available character pool.

2. **Process Each Position i from Left to Right**:

 - **Case 1 - Keep the Same Character**: If target\[i\] is still available in the character pool, try to place it at the current position. After placing it, check whether the remaining characters can form a string greater than the suffix target\[i+1:\].

 - **Judgment Method**: Construct the **largest lexicographical string** from the remaining characters by arranging them in descending order, and compare it with the suffix target\[i+1:\]. If the largest possible string is greater than the suffix, it is possible to continue constructing the result, so we keep the character and proceed to the next position.

 - If the largest possible string is not greater than the suffix, then even arranging all remaining characters in descending order cannot produce a string greater than target. Therefore, we backtrack by returning the character to the pool and proceed to Case 2.

 - **Case 2 - Place a Larger Character**: Find the smallest available character that is greater than target\[i\] in the character pool and place it at the current position. At this point, regardless of how the remaining characters are arranged, the resulting string will already be greater than target. To obtain the **smallest** permutation greater than target, append the remaining characters in **ascending order** and return the result.

 - **No Feasible Solution Found**: If Case 1 is not feasible and no larger character can be found in Case 2, it is impossible to construct a valid result using the current prefix, so we return an empty string.

3. **Boundary Case**: If all positions are processed without returning a result, it means that target itself is the largest permutation that can be formed from the characters in s. Therefore, we return an empty string.

#### Complexity Analysis

Let n be the length of the string, and let ∣Σ∣\=26 be the size of the character set.

- Time complexity: O(n×(n+∣Σ∣)).

 For each position, we may try up to ∣Σ∣ characters. Each feasibility check constructs the maximum string in O(n) time and compares it with the corresponding suffix in O(n) time. Therefore, the total time complexity is O(n×(n+∣Σ∣)).

- Space complexity: O(∣Σ∣).

 The character count array requires O(∣Σ∣) space.

---

### Approach 2: Reverse Greedy

#### Intuition

Unlike Approach 1, which processes the string from left to right, this approach enumerates the **boundary position** i from right to left. This allows us to avoid constructing and comparing strings during each iteration and only requires checking an O(∣Σ∣) count array.

#### Algorithm

1. **Initialize the Character Count Array**: First, we need to determine the **surplus** of each character in s relative to target. We use cnt\[c\] to represent the number of additional occurrences of character c that are available after matching the entire target.

 In other words, when we enumerate the boundary position i from right to left, cnt\[c\] represents how many more instances of character c can be used in the remaining positions after matching the prefix of the result with target. If cnt\[c\]<0, there are not enough occurrences of character c to match the required prefix.

 To initialize cnt, we traverse both s and target:

 - When encountering a character in s, increment its count by 1, indicating that the character is available.
 - When encountering a character in target, decrement its count by 1, indicating that the character is consumed.
2. **Enumerate Boundary Positions from Right to Left**: Start from the last position and traverse from right to left, that is, i\=n−1,n−2,…,0.

 - **Undo the Consumption at the Current Position**: Increment the count of target\[i\] by 1.Thisisequivalenttoundoingtheconsumptionofthecharacteratpositioni$, allowing us to change the character at this position while keeping the prefix before it unchanged.

 - **Check Prefix Feasibility**: After the rollback, check whether any value in cnt is negative:

 - If some value is negative, it means that even after removing the requirement at position i, the prefix \[0,i−1\] cannot be completely matched because some characters are insufficient. Therefore, position i cannot be used as the boundary, and we continue moving left.

 - If no value is negative, it means that the prefix \[0,i−1\] can be completely matched with the available characters, so position i is a valid candidate for the boundary.

 - **Find a Larger Character**: Enumerate the characters greater than target\[i\]:

 - Find the smallest character j such that j\>target\[i\] and cnt\[j\]\>0. Place this character at position i and decrement cnt\[j\] by 1. The prefix \[0,i−1\] remains unchanged.

 - **Construct the Minimal Suffix**: Append all remaining characters in cnt in ascending order. Since the character at position i has already been increased, the resulting string is guaranteed to be greater than target. Arranging the remaining characters in ascending order gives the lexicographically smallest possible suffix.

3. **No Solution Case**: If all positions are processed and no feasible boundary position is found, it is impossible to construct a permutation greater than target, so we return an empty string.

#### Details

- **Why do we enumerate from right to left?** 
 We are looking for the lexicographically smallest permutation greater than target. Therefore, the boundary position should be as far to the right as possible, allowing the longest possible prefix to remain identical to target. This produces the smallest possible result.

- **Why do we check min(cnt)<0 after the rollback?** 
 After rolling back position i, the cnt array represents the characters remaining after matching the prefix \[0,i−1\] with target. If any count is negative, it means that there are not enough occurrences of some character to construct this prefix, so position i cannot be used as the boundary.

- **Why do we arrange the remaining characters in ascending order?** 
 Once the character at position i is increased, the resulting string is already greater than target regardless of how the suffix is arranged. Therefore, to minimize the resulting permutation, we arrange the remaining characters in ascending order.

#### Complexity Analysis

Let n be the length of the given string, and let ∣Σ∣\=26 be the size of the character set.

- Time complexity: O(n∣Σ∣).

 We enumerate each position of target from right to left. At each position, checking whether the prefix can be matched takes O(∣Σ∣) time, and finding a larger character also takes O(∣Σ∣) time. Therefore, the total time complexity is O(n∣Σ∣).

- Space complexity: O(∣Σ∣)

 The character count array requires O(∣Σ∣) space.
