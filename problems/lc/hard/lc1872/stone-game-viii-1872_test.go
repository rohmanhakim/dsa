package lc1872

import (
	"reflect"
	"testing"
)

func TestStoneGameVIII(t *testing.T) {
	tests := []struct {
		name   string
		stones []int
		want   int
	}{

		{
			name:   "example 1",
			stones: []int{-1, 2, -3, 4, -5},
			want:   5,
		},

		{
			name:   "example 2",
			stones: []int{7, -6, 5, 10, 5, -2, -6},
			want:   13,
		},

		{
			name:   "example 3",
			stones: []int{-10, -12},
			want:   -22,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := stoneGameVIII(tt.stones)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
