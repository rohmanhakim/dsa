package lc836

import (
	"reflect"
	"testing"
)

func TestIsRectangleOverlap(t *testing.T) {
	tests := []struct {
		name string
		rec1 []int
		rec2 []int
		want bool
	}{

		{
			name: "example 1",
			rec1: []int{0, 0, 2, 2},
			rec2: []int{1, 1, 3, 3},
			want: true,
		},

		{
			name: "example 2",
			rec1: []int{0, 0, 1, 1},
			rec2: []int{1, 0, 2, 1},
			want: false,
		},

		{
			name: "example 3",
			rec1: []int{0, 0, 1, 1},
			rec2: []int{2, 2, 3, 3},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := isRectangleOverlap(tt.rec1, tt.rec2)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
