package lc875

import (
	"reflect"
	"testing"
)

func TestMinEatingSpeed(t *testing.T) {
	tests := []struct {
		name  string
		piles []int
		h     int
		want  int
	}{

		{
			name:  "example 1",
			piles: []int{3, 6, 7, 11},
			h:     8,
			want:  4,
		},

		{
			name:  "example 2",
			piles: []int{30, 11, 23, 4, 20},
			h:     5,
			want:  30,
		},

		{
			name:  "example 3",
			piles: []int{30, 11, 23, 4, 20},
			h:     6,
			want:  23,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := minEatingSpeed(tt.piles, tt.h)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
