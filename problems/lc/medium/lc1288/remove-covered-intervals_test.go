package lc1288

import (
	"reflect"
	"testing"
)

func TestRemoveCoveredIntervals(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      int
	}{
		{
			name:      "example 1",
			intervals: [][]int{{1, 4}, {3, 6}, {2, 8}},
			want:      2,
		},
		{
			name:      "example 2",
			intervals: [][]int{{1, 4}, {2, 3}},
			want:      1,
		},
		{
			name:      "example 3",
			intervals: [][]int{{3, 10}, {4, 10}, {5, 11}},
			want:      2,
		},
		{
			name:      "example 4",
			intervals: [][]int{{1, 2}, {1, 4}, {3, 4}},
			want:      1,
		},
		{
			name:      "example 5",
			intervals: [][]int{{34335, 39239}, {15875, 91969}, {29673, 66453}, {53548, 69161}, {40618, 93111}},
			want:      2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := removeCoveredIntervals(tt.intervals)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
