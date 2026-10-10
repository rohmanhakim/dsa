package lc703

import (
	"slices"
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		k      int
		nums   []int
		values []int
		want   []int
	}{
		{
			name:   "example 1",
			k:      3,
			nums:   []int{4, 5, 8, 2},
			values: []int{3, 5, 10, 9, 4},
			want:   []int{4, 5, 5, 8, 8},
		},
		{
			name:   "example 2",
			k:      4,
			nums:   []int{7, 7, 7, 7, 8, 3},
			values: []int{2, 10, 9, 9},
			want:   []int{7, 7, 7, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := Constructor(tt.k, tt.nums)

			res := []int{}

			for _, num := range tt.values {
				res = append(res, obj.Add(num))
			}

			if !slices.Equal(res, tt.want) {
				t.Fatalf("got %v want %v", res, tt.want)
			}
		})
	}
}
