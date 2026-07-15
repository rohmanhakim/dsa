package lc57

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		name        string
		intervals   [][]int
		newInterval []int
		want        [][]int
	}{

		{
			name:        "example 1",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			want:        [][]int{{1, 5}, {6, 9}},
		},

		{
			name:        "example 2",
			intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{4, 8},
			want:        [][]int{{1, 2}, {3, 10}, {12, 16}},
		},

		{
			name:        "example 3",
			intervals:   [][]int{{1, 5}},
			newInterval: []int{2, 3},
			want:        [][]int{{1, 5}},
		},
		{
			name:        "example 4",
			intervals:   [][]int{{1, 5}},
			newInterval: []int{6, 8},
			want:        [][]int{{1, 5}, {6, 8}},
		},
		{
			name:        "example 5",
			intervals:   [][]int{{1, 5}},
			newInterval: []int{0, 3},
			want:        [][]int{{0, 5}},
		},
		{
			name:        "example 6",
			intervals:   [][]int{{1, 5}},
			newInterval: []int{0, 0},
			want:        [][]int{{0, 0}, {1, 5}},
		},
		{
			name:        "example 7",
			intervals:   [][]int{{3, 5}, {12, 15}},
			newInterval: []int{6, 6},
			want:        [][]int{{3, 5}, {6, 6}, {12, 15}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := insert(tt.intervals, tt.newInterval)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
