package lc704

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
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			want:   4,
		},
		{
			name:   "example 2",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			want:   -1,
		},
		{
			name:   "example 3",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: -1,
			want:   0,
		},
		{
			name:   "example 5",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 12,
			want:   5,
		},
		{
			name:   "example 6",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 0,
			want:   1,
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
