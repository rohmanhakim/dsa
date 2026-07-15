# LeetCode's Official Editorial

### Approach 1: Euclidean Algorithm

#### Intuition

To compute the greatest common divisor (GCD) of two integers, we can use the Euclidean algorithm.

The key observation behind the Euclidean algorithm is that the greatest common divisor of two integers is equal to the greatest common divisor of the **second integer** and the **remainder when the first integer is divided by the second integer**. Formally,

gcd(a,b)\=gcd(b,amodb)

We define gcd(a,0)\=a, which naturally leads to a recursive implementation. Specifically, for gcd(x,y):

- If y\=0, then gcd(x,y)\=gcd(y,xmody).
- Otherwise, gcd(x,y)\=x.

To compute the required values sumOdd and sumEven, we can directly apply the arithmetic series formulas:

sumOdd\=1+3+⋯+(2n−1)\=n2

sumEven\=2+4+⋯+2n\=n(n+1)

#### Complexity Analysis

- Time complexity: O(log(n)).

- Space complexity: O(log(n)).

 In the worst case, the recursion depth is O(logn).

---

### Approach 2: Mathematics

#### Intuition

Observe that gcd(n2,n(n+1))\=n×gcd(n,n+1)

Since two consecutive integers are always coprime, gcd(n,n+1)\=1

Therefore, gcd(n2,n(n+1))\=n

which is exactly the required answer.

#### Complexity Analysis

- Time complexity: O(1).

- Space complexity: O(1).