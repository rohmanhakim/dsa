package lc4056

import (
	"reflect"
	"testing"
)

func TestCountIntersectingIntervals(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      int
	}{

		{
			name:      "example 1",
			intervals: [][]int{{1, 2}, {2, 3}, {3, 4}},
			want:      2,
		},

		{
			name:      "example 2",
			intervals: [][]int{{1, 5}, {2, 4}, {3, 6}},
			want:      3,
		},

		{
			name:      "example 3",
			intervals: [][]int{{1, 2}, {3, 4}, {5, 6}},
			want:      0,
		},

		{
			name:      "example 4",
			intervals: [][]int{{97, 100}, {61, 61}},
			want:      0,
		},

		{
			name:      "example 5",
			intervals: [][]int{{61, 61}, {97, 100}},
			want:      0,
		},

		{
			name:      "example 6",
			intervals: [][]int{{1, 3}, {2, 4}},
			want:      1,
		},

		{
			name:      "example 7",
			intervals: [][]int{{2, 4}, {1, 3}},
			want:      1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := countIntersectingIntervals(tt.intervals)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
