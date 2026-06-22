package medium

import (
	"reflect"
	"testing"
)

func TestRottingOranges(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "example 1",
			grid: [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}},
			want: 4,
		},
		{
			name: "example 2",
			grid: [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}},
			want: -1,
		},
		{
			name: "example 3",
			grid: [][]int{{0, 2}},
			want: 0,
		},
		{
			name: "example 4",
			grid: [][]int{{0}},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := orangesRotting(tt.grid)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
