# LeetCode's Official Editorial

### Approach 1: Stack

#### Intuition

We traverse s and use a stack to track the nesting depth of the parentheses. When we encounter (, we push it onto the stack; when we encounter ), we pop the top element from the stack. The interval from when the stack first becomes non-empty until it becomes empty again corresponds to scanning a single primitive string. Within each primitive string, the outermost opening and closing parentheses must be discarded, while all inner characters are added to the result.

In the implementation, we can streamline the conditions by ordering the operations carefully:

- For ), we pop from the stack _before_ checking whether the stack is empty (so the outermost ) leaves the stack empty and is excluded).
- For (, we check whether the stack is non-empty _before_ pushing (so the outermost ( sees an empty stack and is excluded).

This allows a single check on whether the stack is non-empty to handle both cases.

#### Complexity Analysis

Let n be the length of the input s.

- Time complexity: O(n).

 We only need to traverse s once.

- Space complexity: O(n).

 The stack stores at most n/2 characters in the worst case (e.g., when all opening parentheses are nested), taking O(n) auxiliary space.

---

### Approach 2: Counting

#### Intuition

We can maintain an integer variable `level` representing the current nesting depth, adding 1 for ( and subtracting 1 for ). Starting with level\=0, each time level drops back to 0, we have reached the end of a primitive valid parentheses string.

To omit the outermost parentheses of each primitive string without extra conditional checks:

- When encountering (, we only append it to the result if level\>0 (meaning it is not the opening parenthesis of a primitive string), and then increment level.
- When encountering ), we first decrement level, and only append it to the result if level\>0 (meaning it is not the closing parenthesis of a primitive string).

#### Complexity Analysis

Let n be the length of the input s.

- Time complexity: O(n).

 We only need to traverse s once.

- Space complexity: O(1).

 Excluding the space required to store the output string, only a constant amount of extra space is used for variables.

---
