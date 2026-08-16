package lc2029

import (
	"reflect"
	"testing"
)

func TestStoneGameIX(t *testing.T) {
	tests := []struct {
		name   string
		stones []int
		want   bool
	}{

		{
			name:   "example 1",
			stones: []int{2, 1},
			want:   true,
		},

		{
			name:   "example 2",
			stones: []int{2},
			want:   false,
		},

		{
			name:   "example 3",
			stones: []int{5, 1, 2, 4, 3},
			want:   false,
		},

		{
			name:   "example 4",
			stones: []int{3, 3},
			want:   false,
		},

		{
			name:   "example 5",
			stones: []int{20, 3, 20, 17, 2, 12, 15, 17, 4},
			want:   true,
		},

		{
			name:   "example 6",
			stones: []int{20, 3, 20, 17, 2, 12, 15, 17, 4, 5},
			want:   true,
		},

		{
			name:   "example 7",
			stones: []int{20, 20, 17, 2, 17, 4},
			want:   true,
		},

		{
			name:   "example 8",
			stones: []int{20, 20, 17, 12, 2, 17, 4, 6},
			want:   true,
		},

		{
			name:   "example 9",
			stones: []int{2, 3, 1},
			want:   false,
		},

		{
			name:   "example 10",
			stones: []int{2, 3, 1, 3},
			want:   true,
		},

		{
			name:   "example 11",
			stones: []int{2, 3},
			want:   false,
		},

		{
			name:   "example 12",
			stones: []int{2, 2, 3},
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := stoneGameIX(tt.stones)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
