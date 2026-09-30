package lc1111

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

// The problem accepts ANY split into two VPSs whose max depth equals the
// optimum ceil(maxDepth/2), so we validate that contract instead of
// exact-matching one arbitrary answer array (LeetCode's example outputs
// are just one of many valid answers; e.g. for "(()(()))" both
// [0,1,1,1,0,0,1,0] and [0,0,0,0,1,0,0,1] are correct).

// seqMaxDepth returns the nesting depth of the original sequence.
func seqMaxDepth(seq string) int {
	depth, hi := 0, 0
	for i := 0; i < len(seq); i++ {
		if seq[i] == '(' {
			depth++
			if depth > hi {
				hi = depth
			}
		} else {
			depth--
		}
	}
	return hi
}

// vpsDepth returns the nesting depth of sub if it is a valid parentheses
// string, or -1 if it is not (unbalanced or a closing bracket whose
// matching open landed in the other group).
func vpsDepth(sub string) int {
	depth, hi := 0, 0
	for i := 0; i < len(sub); i++ {
		if sub[i] == '(' {
			depth++
			if depth > hi {
				hi = depth
			}
		} else {
			depth--
			if depth < 0 {
				return -1
			}
		}
	}
	if depth != 0 {
		return -1
	}
	return hi
}

// checkSplit returns nil if out is a correct answer for seq, or an error
// describing the first contract violation:
//  1. out has one binary label per character
//  2. each group's subsequence is a VPS (pairs stay intact)
//  3. max(depth(A), depth(B)) equals the optimum ceil(maxDepth/2)
//     (catches degenerate answers like "all zeros" that are valid
//     splits but not minimal)
func checkSplit(seq string, out []int) error {
	if len(out) != len(seq) {
		return fmt.Errorf("len(out) = %d, want %d", len(out), len(seq))
	}

	var subs [2]string
	for i := 0; i < len(seq); i++ {
		if out[i] != 0 && out[i] != 1 {
			return fmt.Errorf("out[%d] = %d, want 0 or 1", i, out[i])
		}
		subs[out[i]] += string(seq[i])
	}

	var depths [2]int
	for g := 0; g < 2; g++ {
		d := vpsDepth(subs[g])
		if d < 0 {
			return fmt.Errorf("group %d = %q is not a valid parentheses string", g, subs[g])
		}
		depths[g] = d
	}

	got := max(depths[0], depths[1])
	target := (seqMaxDepth(seq) + 1) / 2 // ceil(maxDepth/2)
	if got != target {
		return fmt.Errorf("max(depth(A), depth(B)) = %d, want optimal %d (A=%q d=%d B=%q d=%d)",
			got, target, subs[0], depths[0], subs[1], depths[1])
	}
	return nil
}

// validateSplit fails the test if out violates the problem contract.
func validateSplit(t *testing.T, seq string, out []int) {
	t.Helper()
	if err := checkSplit(seq, out); err != nil {
		t.Fatalf("%v (seq=%q out=%v)", err, seq, out)
	}
}

func TestMaxDepthAfterSplit(t *testing.T) {
	tests := []struct {
		name string
		seq  string
	}{
		{name: "example 1", seq: "(()())"},
		{name: "example 2", seq: "()(())()"},
		{name: "example 3", seq: "(()(()))"},
		{name: "single pair", seq: "()"},
		{name: "pure nesting depth 3", seq: "((()))"},
		{name: "pure nesting depth 4", seq: "(((())))"},
		{name: "flat concatenation", seq: "()()()()"},
		{name: "mixed", seq: "()((())())(()())"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validateSplit(t, tt.seq, maxDepthAfterSplit(tt.seq))
		})
	}
}

// randomVPS generates a valid parentheses string with the given number of
// pairs: the first '(' closes after k-1 nested pairs, and the remaining
// pairs form the suffix. k is chosen uniformly, giving a random Dyck word.
func randomVPS(rng *rand.Rand, pairs int) string {
	if pairs == 0 {
		return ""
	}
	k := 1 + rng.IntN(pairs)
	return "(" + randomVPS(rng, k-1) + ")" + randomVPS(rng, pairs-k)
}

// TestMaxDepthAfterSplitRandom runs the contract validator against random
// valid sequences (fixed seed for reproducibility).
func TestMaxDepthAfterSplitRandom(t *testing.T) {
	rng := rand.New(rand.NewPCG(42, 43))

	for i := 0; i < 200; i++ {
		seq := randomVPS(rng, 1+rng.IntN(30)) // 1..30 pairs
		validateSplit(t, seq, maxDepthAfterSplit(seq))
	}
}

// TestValidatorRejectsBadSplits proves the validator itself can fail:
// each out below is a wrong answer (including the degenerate all-zeros
// output of an unfinished implementation) and must be rejected.
func TestValidatorRejectsBadSplits(t *testing.T) {
	cases := []struct {
		name string
		seq  string
		out  []int
	}{
		{name: "all zeros (non-optimal)", seq: "((()))", out: []int{0, 0, 0, 0, 0, 0}},
		{name: "orphan closer (not a VPS)", seq: "(()())", out: []int{0, 0, 0, 0, 0, 1}},
		{name: "wrong length", seq: "()", out: []int{0}},
		{name: "non-binary label", seq: "(())", out: []int{0, 2, 2, 0}},
	}

	for _, c := range cases {
		if err := checkSplit(c.seq, c.out); err == nil {
			t.Errorf("validator accepted bad split: seq=%q out=%v", c.seq, c.out)
		}
	}
}
