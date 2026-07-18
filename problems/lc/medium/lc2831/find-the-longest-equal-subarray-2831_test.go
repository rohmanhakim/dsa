package lc2831

import (
	"reflect"
	"testing"
)

func TestLongestEqualSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{

		{
			name: "example 1",
			nums: []int{1, 3, 2, 3, 1, 3},
			k:    3,
			want: 3,
		},

		{
			name: "example 2",
			nums: []int{1, 1, 2, 2, 1, 1},
			k:    2,
			want: 4,
		},

		{
			name: "example 3",
			nums: []int{1},
			k:    0,
			want: 1,
		},

		{
			name: "example 4",
			nums: []int{6, 3, 7, 5, 1, 3, 3, 7, 3},
			k:    1,
			want: 3,
		},

		{
			name: "example 5",
			nums: []int{7, 5, 7, 6, 6, 5, 5},
			k:    1,
			want: 2,
		},

		{
			name: "example 6",
			nums: []int{3, 1, 1},
			k:    2,
			want: 2,
		},

		{
			name: "example 7",
			nums: []int{1, 2, 1},
			k:    0,
			want: 1,
		},

		{
			name: "example 8",
			nums: []int{3, 2, 8, 4, 5, 5, 6, 7, 10, 7, 10, 6, 2, 5, 5},
			k:    6,
			want: 2,
		},

		{
			name: "example 9",
			nums: []int{3, 2, 4, 2},
			k:    1,
			want: 2,
		},

		{
			name: "example 10",
			nums: []int{5, 3, 2, 3, 4},
			k:    1,
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := longestEqualSubarray(tt.nums, tt.k)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
