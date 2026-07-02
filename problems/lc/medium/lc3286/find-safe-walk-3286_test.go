package lc3286

import (
	"reflect"
	"testing"
)

func TestFindaSafeWalkThroughaGrid(t *testing.T) {
	tests := []struct {
		name   string
		grid   [][]int
		health int
		want   bool
	}{

		{
			name:   "example 1",
			grid:   [][]int{{0, 1, 0, 0, 0}, {0, 1, 0, 1, 0}, {0, 0, 0, 1, 0}},
			health: 1,
			want:   true,
		},
		{
			name:   "example 2",
			grid:   [][]int{{0, 1, 1, 0, 0, 0}, {1, 0, 1, 0, 0, 0}, {0, 1, 1, 1, 0, 1}, {0, 0, 1, 0, 1, 0}},
			health: 3,
			want:   false,
		},
		{
			name:   "example 3",
			grid:   [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			health: 5,
			want:   true,
		},
		{
			name:   "example 4",
			grid:   [][]int{{1, 1, 1}, {1, 0, 0}, {0, 1, 1}, {0, 1, 1}, {1, 1, 1}},
			health: 6,
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := findSafeWalkDijkstra(tt.grid, tt.health)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
