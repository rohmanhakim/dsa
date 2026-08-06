package lc49

import (
	"testing"

	"github.com/rohmanhakim/leetcode/lib"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{

		{
			name: "example 1",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},

		{
			name: "example 2",
			strs: []string{""},
			want: [][]string{{""}},
		},

		{
			name: "example 3",
			strs: []string{"a"},
			want: [][]string{{"a"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := groupAnagrams(tt.strs)

			if !lib.Equal2DUnordered(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
