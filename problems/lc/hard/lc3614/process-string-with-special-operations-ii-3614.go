/*
You are given a string s consisting of lowercase English letters and the special characters: '*', '#', and '%'.

You are also given an integer k.

Build a new string result by processing s according to the following rules from left to right:

	If the letter is a lowercase English letter append it to result.
	A '*' removes the last character from result, if it exists.
	A '#' duplicates the current result and appends it to itself.
	A '%' reverses the current result.

Return the kth character of the final string result. If k is out of the bounds of result, return '.'.

Example 1:

Input: s = "a#b%*", k = 1

Output: "a"

Explanation:
i	s[i]	Operation	Current result
0	'a'	Append 'a'	"a"
1	'#'	Duplicate result	"aa"
2	'b'	Append 'b'	"aab"
3	'%'	Reverse result	"baa"
4	'*'	Remove the last character	"ba"

The final result is "ba". The character at index k = 1 is 'a'.

Example 2:

Input: s = "cd%#*#", k = 3

Output: "d"

Explanation:
i	s[i]	Operation	Current result
0	'c'	Append 'c'	"c"
1	'd'	Append 'd'	"cd"
2	'%'	Reverse result	"dc"
3	'#'	Duplicate result	"dcdc"
4	'*'	Remove the last character	"dcd"
5	'#'	Duplicate result	"dcddcd"

The final result is "dcddcd". The character at index k = 3 is 'd'.

Example 3:

Input: s = "z*#", k = 0

Output: "."

Explanation:
i	s[i]	Operation	Current result
0	'z'	Append 'z'	"z"
1	'*'	Remove the last character	""
2	'#'	Duplicate the string	""

The final result is "". Since index k = 0 is out of bounds, the output is '.'.

Constraints:

	1 <= s.length <= 10^5
	s consists of only lowercase English letters and special characters '*', '#', and '%'.
	0 <= k <= 10^15
	The length of result after processing s will not exceed 10^15.
*/
package lc3614

import (
	"fmt"
)

// processStr finds the k-th character of the result string without building it.
//
// ALGORITHM OVERVIEW (Two-Pass Approach):
//
// The naive approach builds the entire string in memory, but '#' (duplicate)
// causes exponential growth — with up to 10^5 operations, the result can reach
// 10^15 characters. That's ~1 petabyte of memory, which causes OOM.
//
// Key insight: since we only need the k-th character, we never need to build
// the string. Instead:
//
//	Phase 1 (Forward): Process operations left-to-right, tracking only the
//	 logical length of the string (as int64). Store each operation for Phase 2.
//	 Space: O(n) where n = len(s). No string allocation.
//
//	Phase 2 (Backward): Process operations right-to-left, starting from position
//	 k in the final string. For each operation, compute where position k "came
//	 from" before that operation was applied. If we trace back to an append
//	 operation and k matches the appended position, we've found our answer.
//
// HOW EACH OPERATION MAPS BACKWARD:
//
// Let n = logical length BEFORE the operation, length = logical length AFTER.
//
//	Append 'c':  length = n + 1
//	  Backward: n = length - 1. If k == n → answer is 'c' (it was just appended).
//	  If k < n → k stays the same (position existed before this append).
//
//	'*' (delete last):  length = n - 1  (or 0 if string was empty)
//	  Backward: n = length + 1. k stays the same (position is unchanged).
//
//	'#' (duplicate):  length = 2 * n
//	  Backward: n = length / 2. If k < n → it's in the first half, k stays.
//	  If k >= n → it's in the second half (a copy), so k = k - n.
//
//	'%' (reverse):  length = n (unchanged)
//	  Backward: n = length. Position k maps to n - 1 - k (mirror position).
//
// EXAMPLE TRACE — s = "a#b%*", k = 1:
//
//	Forward (tracking length):
//	  'a' → 1,  '#' → 2,  'b' → 3,  '%' → 3,  '*' → 2
//
//	Backward (tracing k = 1):
//	  '*' → n=3, k=1 stays     (delete was at end, doesn't affect pos 1)
//	  '%' → n=3, k=3-1-1=1     (reverse maps pos 1 → pos 1 in "baa")
//	  'b' → n=2, k=1≠2 continue(pos 1 ≠ appended pos 2)
//	  '#' → n=1, k=1≥1 → k=0   (pos 1 was in second half of duplicate)
//	  'a' → n=0, k=0==0 → 'a'  (pos 0 matches appended position) ✓
func processStr(s string, k int64) byte {

	isAlphabet := func(char rune) bool {
		return char >= 'a' && char <= 'z'
	}

	// ─── PHASE 1: Forward pass — track logical length and store operations ───
	// We only track the length as an int64 (can hold up to ~9.2×10^18, well
	// above the 10^15 constraint). No string/slice of characters is allocated.
	// This is where the O(n) space comes from (storing the operations array).
	length := int64(0)
	ops := []rune{}
	for _, char := range s {
		if isAlphabet(char) {
			length = length + 1 // append grows length by 1
		}
		if char == '*' {
			// Clamp at 0: '*' on an empty string is a no-op per the problem.
			// Without this clamp, length could go negative and corrupt
			// all subsequent length tracking.
			length = max(length-1, 0)
		}
		if char == '#' {
			length = length * 2 // duplicate doubles the length
		}
		// '%' doesn't change length, only reverses content
		ops = append(ops, char)
	}

	// ─── PHASE 2: Backward pass — trace position k to its source ─────────
	// We walk operations in reverse. At each step, `length` is the logical
	// length AFTER the current operation, and `n` is the length BEFORE it.
	// We use `n` to determine where position k came from.
	for i := len(ops) - 1; i >= 0; i-- {
		op := ops[i]
		n := int64(0)

		// If k is out of bounds for the current string length, the character
		// doesn't exist → return '.'
		if k >= length {
			return byte('.')
		}

		// APPEND: The character was placed at position n (= length - 1).
		// If k == n, this is the character we're looking for!
		// If k < n, the position existed before this append — continue tracing.
		if isAlphabet(op) {
			n = length - 1
			if k == n {
				return byte(op)
			}
		}

		// DELETE LAST: The string was one character longer before this op.
		// Position k is unchanged (delete only affected the end).
		if op == '*' {
			n = length + 1
		}

		// DUPLICATE: The string was half as long before this op.
		// If k was in the first half (k < n), it maps to position k.
		// If k was in the second half (k >= n), it's a copy → map to k - n.
		if op == '#' {
			n = length / 2
			if k >= n {
				k = k - n
			}
		}

		// REVERSE: Length is unchanged, but position k mirrors to n - 1 - k.
		if op == '%' {
			n = length
			k = n - 1 - k
		}

		// Update length to the pre-operation length for the next iteration
		length = n
	}

	// If we traced all the way back without finding a match, k is invalid
	return byte('.')
}

func ProcessStringWithSpecialOperationsII() {
	s := "a#b%*"
	k := 1

	// s := "cd%#*#"
	// k := 3

	// s := "*mi#"
	// k := 442611269896338 // .

	// s := "%%pd#jw#%c*vi##nm%e#a#oz####k##rlow%*w%f#u%%h#f*#f#g##%#cnq#"
	// k := 3580

	out := processStr(s, int64(k))
	fmt.Println(string(out))
}
