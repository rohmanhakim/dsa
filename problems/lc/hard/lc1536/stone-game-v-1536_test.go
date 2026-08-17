package lc1536

import (
	"reflect"
	"testing"
)

func TestStoneGameV(t *testing.T) {
	tests := []struct {
		name       string
		stoneValue []int
		want       int
	}{

		{
			name:       "example 1",
			stoneValue: []int{6, 2, 3, 4, 5, 5},
			want:       18,
		},

		{
			name:       "example 2",
			stoneValue: []int{7, 7, 7, 7, 7, 7, 7},
			want:       28,
		},

		{
			name:       "example 3",
			stoneValue: []int{4},
			want:       0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := stoneGameV(tt.stoneValue)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
