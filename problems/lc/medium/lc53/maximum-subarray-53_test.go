package lc53

import (
	"reflect"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{

		{
			name: "example 1",
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},

		{
			name: "example 2",
			nums: []int{1},
			want: 1,
		},

		{
			name: "example 3",
			nums: []int{5, 4, -1, 7, 8},
			want: 23,
		},

		{
			name: "example 4",
			nums: []int{8, -19, 5, -4, 20},
			want: 21,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := maxSubArray(tt.nums)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
