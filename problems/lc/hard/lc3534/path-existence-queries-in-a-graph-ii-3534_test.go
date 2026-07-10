package lc3534

import (
	"reflect"
	"testing"
)

func TestPathExistenceQueries(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		nums    []int
		maxDiff int
		queries [][]int
		want    []int
	}{
		{
			name:    "example 1",
			n:       5,
			nums:    []int{1, 8, 3, 4, 2},
			maxDiff: 3,
			queries: [][]int{{0, 3}, {2, 4}},
			want:    []int{1, 1},
		},
		{
			name:    "example 2",
			n:       5,
			nums:    []int{5, 3, 1, 9, 10},
			maxDiff: 2,
			queries: [][]int{{0, 1}, {0, 2}, {2, 3}, {4, 3}},
			want:    []int{1, 2, -1, 1},
		},
		{
			name:    "example 3",
			n:       3,
			nums:    []int{3, 6, 1},
			maxDiff: 1,
			queries: [][]int{{0, 0}, {0, 1}, {1, 2}},
			want:    []int{0, -1, -1},
		},
		{
			name:    "example 4",
			n:       3,
			nums:    []int{19, 8, 9},
			maxDiff: 1,
			queries: [][]int{{2, 1}, {0, 0}},
			want:    []int{1, 0},
		},
		{
			name:    "example 5",
			n:       5,
			nums:    []int{18, 14, 8, 18, 0},
			maxDiff: 8,
			queries: [][]int{{1, 1}, {4, 0}},
			want:    []int{0, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := pathExistenceQueries(tt.n, tt.nums, tt.maxDiff, tt.queries)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
