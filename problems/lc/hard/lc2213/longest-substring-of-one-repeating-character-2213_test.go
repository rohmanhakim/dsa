package lc2213

import (
	"reflect"
	"testing"
)

func TestLongestRepeating(t *testing.T) {
	tests := []struct {
		name            string
		s               string
		queryCharacters string
		queryIndices    []int
		want            []int
	}{

		{
			name:            "example 1",
			s:               "babacc",
			queryCharacters: "bcb",
			queryIndices:    []int{1, 3, 3},
			want:            []int{3, 3, 4},
		},

		{
			name:            "example 2",
			s:               "abyzz",
			queryCharacters: "aa",
			queryIndices:    []int{2, 1},
			want:            []int{2, 3},
		},

		{
			name:            "example 3",
			s:               "mm",
			queryCharacters: "bfviuwsr",
			queryIndices:    []int{0, 0, 1, 0, 0, 1, 1, 0},
			want:            []int{1, 1, 1, 1, 1, 1, 1, 1},
		},

		{
			name:            "example 4",
			s:               "dbgmcagale",
			queryCharacters: "mf",
			queryIndices:    []int{5, 0},
			want:            []int{1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := longestRepeating(tt.s, tt.queryCharacters, tt.queryIndices)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
