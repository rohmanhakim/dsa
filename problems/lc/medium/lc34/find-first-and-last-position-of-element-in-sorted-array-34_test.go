package lc34

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{

		{
			name:   "example 1",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			want:   []int{3, 4},
		},
		{
			name:   "example 2",
			nums:   []int{},
			target: 0,
			want:   []int{-1, -1},
		},
		{
			name:   "example 3",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 5,
			want:   []int{0, 0},
		},
		{
			name:   "example 4",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 10,
			want:   []int{5, 5},
		},
		{
			name:   "example 5",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			want:   []int{-1, -1},
		},
		{
			name:   "example 6",
			nums:   []int{5},
			target: 5,
			want:   []int{0, 0},
		},
		{
			name:   "example 7",
			nums:   []int{5},
			target: 6,
			want:   []int{-1, -1},
		},
		{
			name:   "example 8",
			nums:   []int{2, 2},
			target: 3,
			want:   []int{-1, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := searchRange(tt.nums, tt.target)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
