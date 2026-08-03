package lc1406

import (
	"reflect"
	"testing"
)

func TestStoneGameIII(t *testing.T) {
	tests := []struct {
		name       string
		stoneValue []int
		want       string
	}{

		{
			name:       "example 1",
			stoneValue: []int{1, 2, 3, 7},
			want:       "Bob",
		},

		{
			name:       "example 2",
			stoneValue: []int{1, 2, 3, -9},
			want:       "Alice",
		},

		{
			name:       "example 3",
			stoneValue: []int{1, 2, 3, 6},
			want:       "Tie",
		},
		{
			name:       "example 4",
			stoneValue: []int{-1, -2, -3},
			want:       "Tie",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := stoneGameIII(tt.stoneValue)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
