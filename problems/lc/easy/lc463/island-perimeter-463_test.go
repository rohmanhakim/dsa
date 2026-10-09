package lc463

import (
	"reflect"
	"testing"
)

func TestIslandPerimeter(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{

		{
			name: "example 1",
			grid: [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}},
			want: 16,
		},

		{
			name: "example 2",
			grid: [][]int{{1}},
			want: 4,
		},

		{
			name: "example 3",
			grid: [][]int{{1, 0}},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := islandPerimeter(tt.grid)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}

func TestIslandPerimeterAlternative(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{

		{
			name: "example 1",
			grid: [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}},
			want: 16,
		},

		{
			name: "example 2",
			grid: [][]int{{1}},
			want: 4,
		},

		{
			name: "example 3",
			grid: [][]int{{1, 0}},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := islandPerimeterAlternative(tt.grid)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
