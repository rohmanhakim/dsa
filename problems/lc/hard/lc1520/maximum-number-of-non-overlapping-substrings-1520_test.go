package lc1520

import (
	"testing"

	"github.com/rohmanhakim/leetcode/lib"
)

func TestMaxNumOfSubstrings(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{

		{
			name: "example 1",
			s:    "adefaddaccc",
			want: []string{"e", "f", "ccc"},
		},

		{
			name: "example 2",
			s:    "abbaccd",
			want: []string{"d", "bb", "cc"},
		},

		{
			name: "example 3",
			s:    "abab",
			want: []string{"abab"},
		},

		{
			name: "example 4",
			s:    "ababa",
			want: []string{"ababa"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := maxNumOfSubstrings(tt.s)

			if !lib.Equal1DUnordered(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
