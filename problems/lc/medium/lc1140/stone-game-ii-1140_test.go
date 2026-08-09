package lc1140

import (
	"reflect"
	"testing"
)

func TestStoneGameII(t *testing.T) {
	tests := []struct {
		name  string
		piles []int
		want  int
	}{

		{
			name:  "example 1",
			piles: []int{2, 7, 9, 4, 4},
			want:  10,
		},

		{
			name:  "example 2",
			piles: []int{1, 2, 3, 4, 5, 100},
			want:  104,
		},

		{
			name:  "example 3",
			piles: []int{1},
			want:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := stoneGameII(tt.piles)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
