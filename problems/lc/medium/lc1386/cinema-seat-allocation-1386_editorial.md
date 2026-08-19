# LeetCode's Official Editorial

### Approach: Bitwise Operations

#### Intuition

For each row, there are only three possible ways to arrange a family:

- Seats 2, 3, 4, 5;
- Seats 4, 5, 6, 7;
- Seats 6, 7, 8, 9.

Therefore, seats 1 and 10 in each row are irrelevant. Even if they are reserved, they do not affect the answer. Thus, we can ignore all reservations at seats 1 and 10.

We can also observe that if a row has no reserved seats among positions 2 to 9, it can accommodate exactly two families: one family in seats 2, 3, 4, 5, and another in seats 6, 7, 8, 9. If at least one of these seats is reserved, the row can accommodate at most one family.

Therefore, we can use 8 binary bits to represent the reservation status of seats 2 to 9 in a row. If seat i is reserved, then the (i−2)\-th bit is set to 1; otherwise, it is set to 0.

For example, in Example 1, the binary representation of each row is as follows:

- **First row:** Seats 2, 3, and 8 are reserved, so the binary representation is (01000011)2​.
- **Second row:** Seat 6 is reserved, so the binary representation is (00010000)2​.
- **Third row:** Seats 1 and 10 are reserved, so the binary representation is (00000000)2​.

We can use a `HashMap` to store the reservation status of each row. For each key-value pair in the `HashMap`, the key represents the row number, and the value represents the corresponding bitmask.

If a row has no reserved seats among positions 2 to 9, such as the third row above, we already know that it can accommodate two families. Therefore, there is no need to store such rows in the `HashMap`. In other words, we only store a row if at least one seat between positions 2 and 9 is reserved.

After processing all reservations, we traverse the `HashMap`. For each key-value pair (row,bitmask), we determine whether the row can accommodate a family. Since every row stored in the `HashMap` has at least one reserved seat, it can accommodate at most one family.

There are three possible arrangements:

- **Seats 2, 3, 4, 5:** If the first four bits of the bitmask are all 0, these seats can be allocated to a family. Equivalently, performing a bitwise OR between the bitmask and (11110000)2​ leaves the latter unchanged.

- **Seats 4, 5, 6, 7:** If the corresponding four bits are all 0, these seats can be allocated to a family. Equivalently, performing a bitwise OR between the bitmask and (11000011)2​ leaves the latter unchanged.

- **Seats 6, 7, 8, 9:** If the last four bits are all 0, these seats can be allocated to a family. Equivalently, performing a bitwise OR between the bitmask and (00001111)2​ leaves the latter unchanged.

Thus, for each row, we perform a bitwise OR operation between the bitmask and each of (11110000)2​, (11000011)2​, and (00001111)2​. If any of these operations leaves the corresponding mask unchanged, the row can accommodate one family.

Finally, there are n−∣S∣ rows with no reserved seats among positions 2 to 9, where ∣S∣ is the number of rows stored in the `HashMap`. Each of these rows can accommodate two families, so we add 2(n−∣S∣) to the answer.

#### Complexity Analysis

Let r be the length of the array reservedSeats.

- Time complexity: O(r).

 We first traverse reservedSeats to record the reservation information in the hash map, and then traverse the hash map to compute the answer. Both operations take O(r) time in total.

- Space complexity: O(r).

 The hash map stores the reservation information for each row that has at least one reserved seat among positions 2 to 9. In the worst case, it can contain O(r) entries.

---
