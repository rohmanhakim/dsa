package lc3532

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
		want    []bool
	}{
		{
			name:    "example 1",
			n:       2,
			nums:    []int{1, 3},
			maxDiff: 1,
			queries: [][]int{{0, 0}, {0, 1}},
			want:    []bool{true, false},
		},
		{
			name:    "example 2",
			n:       4,
			nums:    []int{2, 5, 6, 8},
			maxDiff: 2,
			queries: [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}},
			want:    []bool{false, false, true, true},
		},
		{
			name:    "example 3",
			n:       4,
			nums:    []int{569, 10949},
			maxDiff: 56389,
			queries: [][]int{{0, 0}, {1, 0}},
			want:    []bool{true, true},
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
