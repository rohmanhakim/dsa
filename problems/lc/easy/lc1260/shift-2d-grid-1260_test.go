package lc1260

import (
	"reflect"
	"testing"
)

func TestShiftGrid(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		k    int
		want [][]int
	}{

		{
			name: "example 1",
			grid: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			k:    1,
			want: [][]int{{9, 1, 2}, {3, 4, 5}, {6, 7, 8}},
		},

		{
			name: "example 2",
			grid: [][]int{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}},
			k:    4,
			want: [][]int{{12, 0, 21, 13}, {3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}},
		},

		{
			name: "example 3",
			grid: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			k:    9,
			want: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := shiftGrid(tt.grid, tt.k)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
