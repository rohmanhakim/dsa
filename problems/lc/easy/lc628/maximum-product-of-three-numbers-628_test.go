package lc628

import (
	"reflect"
	"testing"
)

func TestMaximumProduct(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{

		{
			name: "example 1",
			nums: []int{1, 2, 3},
			want: 6,
		},

		{
			name: "example 2",
			nums: []int{1, 2, 3, 4},
			want: 24,
		},

		{
			name: "example 3",
			nums: []int{-1, -2, -3},
			want: -6,
		},

		{
			name: "example 4",
			nums: []int{-100, -98, -1, 2, 3, 4},
			want: 39200,
		},

		{
			name: "example 5",
			nums: []int{-1, -2, 1, 2, 3},
			want: 6,
		},

		{
			name: "example 6",
			nums: []int{-4, -3, -2, -1, 60},
			want: 720,
		},

		{
			name: "example 7",
			nums: []int{-8, -7, -2, 10, 20},
			want: 1120,
		},

		{
			name: "example 8",
			nums: []int{-1, -2, -3, -4},
			want: -6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := maximumProduct(tt.nums)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
