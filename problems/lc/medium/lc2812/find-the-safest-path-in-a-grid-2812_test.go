package lc2812

import (
	"reflect"
	"testing"
)

func TestFindtheSafestPathinaGrid(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "example 1",
			grid: [][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}},
			want: 0,
		},
		{
			name: "example 2",
			grid: [][]int{{0, 0, 1}, {0, 0, 0}, {0, 0, 0}},
			want: 2,
		},
		{
			name: "example 3",
			grid: [][]int{{0, 0, 0, 1}, {0, 0, 0, 0}, {0, 0, 0, 0}, {1, 0, 0, 0}},
			want: 2,
		},
		{
			name: "example 4",
			grid: [][]int{{1}},
			want: 0,
		},
		{
			name: "example 5",
			grid: [][]int{{1, 1, 1}, {1, 1, 1}, {0, 0, 1}},
			want: 0,
		},
		{
			name: "example 6",
			grid: [][]int{{0, 0, 1}, {0, 1, 1}, {0, 0, 0}},
			want: 1,
		},
		{
			name: "example 7",
			grid: [][]int{{1, 1, 1}, {0, 1, 1}, {0, 0, 0}},
			want: 0,
		},
		{
			name: "example 8",
			grid: [][]int{{0, 1, 1}, {0, 0, 1}, {0, 0, 0}},
			want: 1,
		},
		{
			name: "example 9",
			grid: [][]int{{0, 1, 1}, {0, 0, 0}, {0, 0, 0}},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := maximumSafenessFactor(tt.grid)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
