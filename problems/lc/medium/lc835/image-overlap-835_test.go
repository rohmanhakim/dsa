package lc835

import (
	"reflect"
	"testing"
)

func TestLargestOverlap(t *testing.T) {
	tests := []struct {
		name string
		img1 [][]int
		img2 [][]int
		want int
	}{

		{
			name: "example 1",
			img1: [][]int{{1, 1, 0}, {0, 1, 0}, {0, 1, 0}},
			img2: [][]int{{0, 0, 0}, {0, 1, 1}, {0, 0, 1}},
			want: 3,
		},

		{
			name: "example 2",
			img1: [][]int{{1}},
			img2: [][]int{{1}},
			want: 1,
		},

		{
			name: "example 3",
			img1: [][]int{{0}},
			img2: [][]int{{0}},
			want: 0,
		},

		{
			name: "example 4",
			img1: [][]int{{0, 0, 0, 0, 1}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
			img2: [][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := largestOverlap(tt.img1, tt.img2)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
