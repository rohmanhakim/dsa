package lc1840

import (
	"reflect"
	"testing"
)

func TestMaximumBuildingHeight(t *testing.T) {
	tests := []struct {
		name         string
		n            int
		restricitons [][]int
		want         int
	}{
		{
			name:         "test 1",
			n:            5,
			restricitons: [][]int{{2, 1}, {4, 1}},
			want:         2,
		},
		{
			name:         "test 2",
			n:            6,
			restricitons: [][]int{},
			want:         5,
		},
		{
			name:         "test 3",
			n:            10,
			restricitons: [][]int{{5, 3}, {2, 5}, {7, 4}, {10, 3}},
			want:         5,
		},
		{
			name:         "minimal n=2 no restrictions",
			n:            2,
			restricitons: [][]int{},
			want:         1,
		},
		{
			name:         "minimal n=2 restriction forces height 0",
			n:            2,
			restricitons: [][]int{{2, 0}},
			want:         0,
		},
		{
			name:         "restriction at last building",
			n:            10,
			restricitons: [][]int{{10, 3}},
			want:         6,
		},
		{
			name:         "adjacent restrictions",
			n:            5,
			restricitons: [][]int{{2, 1}, {3, 1}},
			want:         3,
		},
		{
			name:         "tight restriction in middle",
			n:            100,
			restricitons: [][]int{{50, 0}},
			want:         50,
		},
		{
			name:         "non-binding restriction",
			n:            5,
			restricitons: [][]int{{3, 100}},
			want:         4,
		},
		{
			name:         "two constraints both tight",
			n:            10,
			restricitons: [][]int{{3, 1}, {8, 1}},
			want:         3,
		},
		{
			name:         "tight left then loose right",
			n:            10,
			restricitons: [][]int{{3, 1}, {5, 100}},
			want:         8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := maxBuilding(tt.n, tt.restricitons)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
