package lc3525

import (
	"reflect"
	"testing"
)

func TestResultArray(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		k       int
		queries [][]int
		want    []int
	}{

		{
			name:    "example 1",
			nums:    []int{1, 2, 3, 4, 5},
			k:       3,
			queries: [][]int{{2, 2, 0, 2}, {3, 3, 3, 0}, {0, 1, 0, 1}},
			want:    []int{2, 2, 2},
		},

		{
			name:    "example 2",
			nums:    []int{1, 2, 4, 8, 16, 32},
			k:       4,
			queries: [][]int{{0, 2, 0, 2}, {0, 2, 0, 1}},
			want:    []int{1, 0},
		},

		{
			name:    "example 3",
			nums:    []int{1, 1, 2, 1, 1},
			k:       2,
			queries: [][]int{{2, 1, 0, 1}},
			want:    []int{5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := resultArray(tt.nums, tt.k, tt.queries)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
