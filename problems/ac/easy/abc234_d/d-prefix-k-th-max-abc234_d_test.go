package abc234d

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// runWithStdio drives a solution entry point end-to-end like the
// online judge: stdin comes from a temp input file, stdout is
// captured to a temp output file, then returned.
func runWithStdio(t *testing.T, input string, fn func()) string {
	t.Helper()
	dir := t.TempDir()

	inFile, err := os.Create(filepath.Join(dir, "input"))
	if err != nil {
		t.Fatalf("create input file: %v", err)
	}
	if _, err := io.WriteString(inFile, input); err != nil {
		t.Fatalf("write input file: %v", err)
	}
	if _, err := inFile.Seek(0, io.SeekStart); err != nil {
		t.Fatalf("rewind input file: %v", err)
	}

	outFile, err := os.Create(filepath.Join(dir, "output"))
	if err != nil {
		t.Fatalf("create output file: %v", err)
	}

	oldStdin, oldStdout := os.Stdin, os.Stdout
	os.Stdin, os.Stdout = inFile, outFile
	defer func() {
		os.Stdin, os.Stdout = oldStdin, oldStdout
		inFile.Close()
		outFile.Close()
	}()

	fn()

	// restore before reading so all output reached the file
	os.Stdin, os.Stdout = oldStdin, oldStdout
	if err := outFile.Close(); err != nil {
		t.Fatalf("close output file: %v", err)
	}

	got, err := os.ReadFile(outFile.Name())
	if err != nil {
		t.Fatalf("read output file: %v", err)
	}
	return string(got)
}

// runPushEvict parses stdin exactly like DPrefixKthMax() but invokes the
// push-then-evict variant, so dPrefixKthMax stays covered even though the
// entry point currently wires to dPrefixKthMaxAlternative.
func runPushEvict() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	perm := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var s int
		fmt.Scanf("%d", &s)
		perm = append(perm, s)
	}
	dPrefixKthMax(n, k, perm)
}

// solutions are the two variants kept in this file. They share the same
// invariant (min-heap = the k largest of the prefix, root = k-th largest)
// but update it differently: unconditional push + evict vs conditional
// replace-root + heap.Fix.
var solutions = []struct {
	name string
	fn   func()
}{
	{"push-then-evict", runPushEvict},
	{"replace-root-fix", DPrefixKthMax},
}

// TestDPrefixKthMax runs fixed cases (samples + edge cases) against both
// variants.
func TestDPrefixKthMax(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "sample 1",
			input: "3 2\n1 2 3",
			want:  "1\n2\n",
		},

		{
			name:  "sample 2",
			input: "11 5\n3 7 2 5 11 6 1 9 8 10 4",
			want:  "2\n3\n3\n5\n6\n7\n7",
		},

		{
			// k=1: answer is the running maximum of the prefix
			name:  "k=1 edge",
			input: "5 1\n3 1 4 5 2",
			want:  "3\n3\n4\n5\n5",
		},

		{
			// k=n: single answer = minimum of the whole permutation
			name:  "k=n edge",
			input: "4 4\n4 3 2 1",
			want:  "1",
		},

		{
			name:  "single element",
			input: "1 1\n5",
			want:  "5",
		},

		{
			// decreasing: the new element is never worthy of the top-k,
			// so the guard fails every iteration and the root (2nd largest
			// = 4) never changes
			name:  "strictly decreasing",
			input: "5 2\n5 4 3 2 1",
			want:  "4\n4\n4\n4",
		},

		{
			// increasing: every new element displaces the root
			name:  "strictly increasing",
			input: "4 3\n1 2 3 4",
			want:  "1\n2",
		},
	}

	for _, tt := range tests {
		for _, sol := range solutions {
			t.Run(tt.name+"/"+sol.name, func(t *testing.T) {
				got := runWithStdio(t, tt.input, sol.fn)
				if acNormalize(got) != acNormalize(tt.want) {
					t.Fatalf("output mismatch:\ngot:  %q\nwant: %q", got, tt.want)
				}
			})
		}
	}
}

// wantBySort computes the expected output independently of any heap:
// copy each prefix, sort descending, take the k-th element.
func wantBySort(perm []int, k int) string {
	var b strings.Builder
	prefix := make([]int, len(perm))
	for i := k; i <= len(perm); i++ {
		copy(prefix, perm[:i])
		sort.Sort(sort.Reverse(sort.IntSlice(prefix[:i])))
		b.WriteString(strconv.Itoa(prefix[k-1]))
		b.WriteByte('\n')
	}
	return b.String()
}

// buildInput renders an input file in the judge's format.
func buildInput(n, k int, perm []int) string {
	var b strings.Builder
	b.WriteString(strconv.Itoa(n))
	b.WriteByte(' ')
	b.WriteString(strconv.Itoa(k))
	b.WriteByte('\n')
	for i, v := range perm {
		if i > 0 {
			b.WriteByte(' ')
		}
		b.WriteString(strconv.Itoa(v))
	}
	return b.String()
}

// TestDPrefixKthMaxVsBruteForce checks both variants against a sort-based
// oracle on seeded random permutations. The oracle shares no code with the
// solutions, so it catches invariant bugs the samples might miss.
func TestDPrefixKthMaxVsBruteForce(t *testing.T) {
	cases := []struct {
		name string
		n, k int
		seed int64
	}{
		{"random n=300 k=7", 300, 7, 42},
		{"random n=300 k=150", 300, 150, 7},
		{"random n=50 k=49", 50, 49, 2026},
	}
	for _, tc := range cases {
		perm := rand.New(rand.NewSource(tc.seed)).Perm(tc.n)
		input := buildInput(tc.n, tc.k, perm)
		want := wantBySort(perm, tc.k)
		for _, sol := range solutions {
			t.Run(tc.name+"/"+sol.name, func(t *testing.T) {
				got := runWithStdio(t, input, sol.fn)
				if acNormalize(got) != acNormalize(want) {
					t.Fatalf("output mismatch:\ngot:  %q\nwant: %q", got, want)
				}
			})
		}
	}
}

// TestDPrefixKthMaxMaxConstraints runs N=500000, K=250000 — the upper
// constraint bound — on the identity permutation 1..N, whose answer is
// analytic: the k-th largest of prefix {1..i} is exactly i-k+1, so the
// expected output is 1,2,...,N-K+1. This exercises correctness at scale
// plus an eviction on every iteration.
func TestDPrefixKthMaxMaxConstraints(t *testing.T) {
	const n = 500000
	const k = 250000

	var inb strings.Builder
	inb.WriteString(strconv.Itoa(n))
	inb.WriteByte(' ')
	inb.WriteString(strconv.Itoa(k))
	inb.WriteByte('\n')
	for i := 1; i <= n; i++ {
		if i > 1 {
			inb.WriteByte(' ')
		}
		inb.WriteString(strconv.Itoa(i))
	}

	var wantb strings.Builder
	for v := 1; v <= n-k+1; v++ {
		wantb.WriteString(strconv.Itoa(v))
		wantb.WriteByte('\n')
	}
	input, want := inb.String(), wantb.String()

	for _, sol := range solutions {
		t.Run(sol.name, func(t *testing.T) {
			got := runWithStdio(t, input, sol.fn)
			if acNormalize(got) != acNormalize(want) {
				t.Fatalf("output mismatch:\ngot:  %.200q\nwant: %.200q", got, want)
			}
		})
	}
}

// acNormalize trims trailing whitespace per line and overall
// so cosmetic newline differences do not fail the test.
func acNormalize(s string) string {
	lines := strings.Split(s, "\n")
	for i, l := range lines {
		lines[i] = strings.TrimRight(l, " \t\r")
	}
	return strings.TrimSpace(strings.Join(lines, "\n"))
}
