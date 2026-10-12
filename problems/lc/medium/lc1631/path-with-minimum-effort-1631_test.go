package lc1631

import (
	"reflect"
	"testing"
)

func TestMinimumEffortPath(t *testing.T) {
	tests := []struct {
		name    string
		heights [][]int
		want    int
	}{

		{
			name:    "example 1",
			heights: [][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}},
			want:    2,
		},

		{
			name:    "example 2",
			heights: [][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}},
			want:    1,
		},

		{
			name:    "example 3",
			heights: [][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}},
			want:    0,
		},

		{
			name:    "example 4",
			heights: [][]int{{1, 10, 6, 7, 9, 10, 4, 9}},
			want:    9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := minimumEffortPath(tt.heights)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
