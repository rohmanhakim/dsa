package lc35

import (
	"reflect"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "example 1",
			nums:   []int{1, 3, 5, 6},
			target: 5,
			want:   2,
		},
		{
			name:   "example 2",
			nums:   []int{1, 3, 5, 6},
			target: 2,
			want:   1,
		},
		{
			name:   "example 3",
			nums:   []int{1, 3, 5, 6},
			target: 7,
			want:   4,
		},
		{
			name:   "example 4",
			nums:   []int{1, 3, 5, 6},
			target: 0,
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := searchInsert(tt.nums, tt.target)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
