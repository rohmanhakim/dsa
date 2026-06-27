package lc435

import (
	"reflect"
	"testing"
)

func TestNonOverlappingIntervals(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      int
	}{
		{
			name: "test 1",
			intervals: [][]int{
				{1, 2},
				{2, 3},
				{3, 4},
				{1, 3},
			},
			want: 1,
		},
		{
			name: "test 2",
			intervals: [][]int{
				{1, 2},
				{1, 2},
				{1, 2},
			},
			want: 2,
		},
		{
			name: "test 3",
			intervals: [][]int{
				{1, 2},
				{2, 3},
			},
			want: 0,
		},
		{
			name:      "test 4",
			intervals: [][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}},
			want:      2,
		},
		{
			name: "test 5",
			intervals: [][]int{
				{0, 2},
				{1, 3},
				{2, 4},
				{3, 5},
				{4, 6},
			},
			want: 2,
		},
		{
			name: "test 6",
			intervals: [][]int{
				{-52, 31}, {-73, -26}, {82, 97}, {-65, -11}, {-62, -49}, {95, 99}, {58, 95}, {-31, 49}, {66, 98}, {-63, 2}, {30, 47}, {-40, -26},
			},
			want: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := eraseOverlapIntervals(tt.intervals)
			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %d want %d", out, tt.want)
			}
		})
	}
}
