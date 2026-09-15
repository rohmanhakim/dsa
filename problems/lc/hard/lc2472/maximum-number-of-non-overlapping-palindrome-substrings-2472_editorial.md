# LeetCode's Official Editorial

### Approach 1: Dynamic Programming

#### Intuition

First, we preprocess all palindromic substrings in the string. Let the two-dimensional boolean array isPalindrome\[i\]\[j\] indicate whether the substring s\[i…j\] is a palindrome. According to the properties of palindromes, we have the following transition:

isPalindrome\[i\]\[j\]\=(s\[i\]\==s\[j\])∧(j−i≤1∨isPalindrome\[i+1\]\[j−1\])

Because determining whether an interval is a palindrome depends on the state of its inner sub-interval, we evaluate intervals in increasing order of length.

Next, let dp\[i\] denote the maximum number of non-overlapping palindromic substrings of length at least k that can be selected from the first i characters s\[0…i−1\].

For the first i characters, there are two choices:

- Do not select a palindromic substring ending at i−1. In this case, dp\[i\]\=dp\[i−1\].
- Enumerate the starting position j of a candidate palindromic substring ending at i−1. If i−j≥k and s\[j…i−1\] is a palindrome, then this substring can be appended after the optimal partition of the first j characters, updating dp\[i\]\=max(dp\[i\],dp\[j\]+1).

Finally, dp\[n\] is the answer.

#### Complexity Analysis

Let n be the length of the string s.

- Time complexity: O(n2).

 Both preprocessing the palindromic intervals and computing dp require O(n2) time.

- Space complexity: O(n2).

 The two-dimensional array isPalindrome requires O(n2) space, and the array dp requires O(n) space.

### Approach 2: Greedy

#### Intuition

This problem can be viewed as an **interval scheduling** problem: to maximize the number of non-overlapping compatible intervals (palindromic substrings of length ≥k), we should greedily select the valid palindrome that finishes earliest. Selecting the palindrome with the smallest ending index leaves the largest possible remaining suffix for future selections, which never hurts optimality.

We only need to check palindromes of length k and k+1. If a palindrome has length \>k+1, stripping characters from both ends leaves a smaller palindrome of length k or k+1 that ends earlier. Since we iterate from left to right, we would have already found that shorter palindrome.

We enumerate the ending position i from left to right and use start to denote the smallest index that the next selected substring may use. For each i, check in order:

- Whether s\[i−k+1…i\] is a palindrome.
- Whether s\[i−k…i\] is a palindrome.

If either substring is valid and does not cross start, select it, increment the answer by one, and set start\=i+1. Any subsequently selected substring is then guaranteed not to overlap it.

#### Complexity Analysis

Let n be the length of the string s.

- Time complexity: O(nk).

 There are n ending positions. For each position, we check at most two substrings of length at most k+1. Each palindrome check takes O(k) time using two pointers, giving an overall time complexity of O(nk).

- Space complexity: O(1).

 Only a constant number of additional variables are used.

---
