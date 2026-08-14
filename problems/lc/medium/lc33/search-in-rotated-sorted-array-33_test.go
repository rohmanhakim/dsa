package lc33

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{

		{
			name:   "example 1",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},

		{
			name:   "example 2",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			want:   -1,
		},

		{
			name:   "example 3",
			nums:   []int{1},
			target: 0,
			want:   -1,
		},

		{
			name:   "example 4",
			nums:   []int{1, 3},
			target: 1,
			want:   0,
		},

		{
			name:   "example 5",
			nums:   []int{3, 1},
			target: 3,
			want:   0,
		},

		{
			name:   "example 6",
			nums:   []int{3, 5, 1},
			target: 3,
			want:   0,
		},

		{
			name:   "example 7",
			nums:   []int{3, 5, 1},
			target: 5,
			want:   1,
		},

		{
			name:   "example 8",
			nums:   []int{5, 1, 3},
			target: 5,
			want:   0,
		},

		{
			name:   "example 9",
			nums:   []int{5, 1, 3},
			target: 2,
			want:   -1,
		},

		{
			name:   "example 10",
			nums:   []int{4, 5, 6, 7, 8, 1, 2, 3},
			target: 8,
			want:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := search(tt.nums, tt.target)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
