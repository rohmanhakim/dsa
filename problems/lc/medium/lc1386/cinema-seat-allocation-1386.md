![](https://assets.leetcode.com/uploads/2020/02/14/cinema_seats_1.png)

```
  1 2 3   4 5 6 7   8 9 10
1 ◻ ◻ ◻   ◻ ◻ ◻ ◻   ◻ ◻ ◻
2 ◻ ◻ ◻   ◻ ◻ ◻ ◻   ◻ ◻ ◻
3 ◻ ◻ ◻   ◻ ◻ ◻ ◻   ◻ ◻ ◻
    .        .        .
    .        .        .
    .        .        .
n ◻ ◻ ◻   ◻ ◻ ◻ ◻   ◻ ◻ ◻
```

A cinema has `n` rows of seats, numbered from 1 to `n`. Each row has 10 seats, numbered from 1 to 10.

You are given a 2D integer array `reservedSeats`, where `reservedSeats[i] = [row<sub>i</sub>, seat<sub>i</sub>]` means that seat `seat<sub>i</sub>` in row `row<sub>i</sub>` is already reserved.

A four-person group must be assigned to four seats in the **same** row. The group can be seated in one of the following seat blocks:

- seats `2, 3, 4, 5`
- seats `4, 5, 6, 7`
- seats `6, 7, 8, 9`

A block can be used only if **none** of its seats are reserved. Each seat can be assigned to **at most** one group.

Return an integer denoting the **maximum** number of four-person groups that can be assigned.

**Example 1:**

```
  1 2 3   4 5 6 7   8 9 10
1 ◻ x x   A A A A   x ◻ ◻
2 ◻ B B   B B x ◻   ◻ ◻ ◻
3 x C C   C C D D   D D x
```

![](https://assets.leetcode.com/uploads/2020/02/14/cinema_seats_3.png)

```
Input: n = 3, reservedSeats = [[1,2],[1,3],[1,8],[2,6],[3,1],[3,10]]
Output: 4
Explanation: The figure above shows an optimal allocation of four groups. Seats marked in blue are already reserved, and each set of four contiguous seats marked in orange is assigned to one group.
```

**Example 2:**

```
Input: n = 2, reservedSeats = [[2,1],[1,8],[2,6]]
Output: 2
```

**Example 3:**

```
Input: n = 4, reservedSeats = [[4,3],[1,4],[4,6],[1,7]]
Output: 4
```

**Constraints:**

- `1 <= n <= 10^9`
- `1 <= reservedSeats.length <= min(10 * n, 10^4)`
- `reservedSeats[i] == [rowᵢ, seatᵢ]`
- `1 <= rowᵢ <= n`
- `1 <= seatᵢ <= 10`
- All `reservedSeats[i]` are distinct.