package lc3116

import (
	"reflect"
	"testing"
)

func TestFindKthSmallest(t *testing.T) {
	tests := []struct {
		name  string
		coins []int
		k     int
		want  int64
	}{

		{
			name:  "example 1",
			coins: []int{3, 6, 9},
			k:     3,
			want:  9,
		},

		{
			name:  "example 2",
			coins: []int{5, 2},
			k:     7,
			want:  12,
		},

		{
			name:  "example 3",
			coins: []int{3, 8, 7},
			k:     6,
			want:  12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := findKthSmallest(tt.coins, tt.k)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
