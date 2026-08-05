package lc1992

import (
	"reflect"
	"testing"
)

func TestFindFarmland(t *testing.T) {
	tests := []struct {
		name string
		land [][]int
		want [][]int
	}{

		{
			name: "example 1",
			land: [][]int{{1, 0, 0}, {0, 1, 1}, {0, 1, 1}},
			want: [][]int{{0, 0, 0, 0}, {1, 1, 2, 2}},
		},

		{
			name: "example 2",
			land: [][]int{{1, 1}, {1, 1}},
			want: [][]int{{0, 0, 1, 1}},
		},

		{
			name: "example 3",
			land: [][]int{{0}},
			want: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := findFarmland(tt.land)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
