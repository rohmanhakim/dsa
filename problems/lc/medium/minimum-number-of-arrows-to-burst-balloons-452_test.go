package medium

import (
	"reflect"
	"testing"
)

func TestMinimumNumberOfArrowsToBurstBalloons(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   int
	}{
		{
			name: "test 1",
			points: [][]int{
				{10, 16},
				{2, 8},
				{1, 6},
				{7, 12},
			},
			want: 2,
		},
		{
			name: "test 2",
			points: [][]int{
				{1, 2},
				{3, 4},
				{5, 6},
				{7, 8},
			},
			want: 4,
		},
		{
			name: "test 3",
			points: [][]int{
				{1, 2},
				{2, 3},
				{3, 4},
				{4, 5},
			},
			want: 2,
		},
		{
			name: "test 4",
			points: [][]int{
				{2, 3},
				{2, 3},
			},
			want: 1,
		},
		{
			name: "test 5",
			points: [][]int{
				{1, 2147483647},
			},
			want: 1,
		},
		{
			name: "test 6",
			points: [][]int{
				{1, 10},
				{2, 11},
				{3, 12},
				{4, 13},
				{5, 14},
			},
			want: 1,
		},
		{
			name: "test 7",
			points: [][]int{
				{-1, 1}, {0, 1}, {2, 3}, {1, 2},
			},
			want: 2,
		},
		{
			name: "test 8",
			points: [][]int{
				{3, 9}, {7, 12}, {3, 8}, {6, 8}, {9, 10}, {2, 9}, {0, 9}, {3, 9}, {0, 6}, {2, 8},
			},
			want: 2,
		},
		{
			name: "test 9",
			points: [][]int{
				{0, 9}, {1, 8}, {7, 8}, {1, 6}, {9, 16}, {7, 13}, {7, 10}, {6, 11}, {6, 9}, {9, 13}},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := findMinArrowShots(tt.points)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
