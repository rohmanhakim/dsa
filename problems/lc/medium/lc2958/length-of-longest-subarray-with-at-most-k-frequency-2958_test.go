package lc2958

import (
	"reflect"
	"testing"
)

func TestMaxSubarrayLength(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{

		{
			name: "example 1",
			nums: []int{1, 2, 3, 1, 2, 3, 1, 2},
			k:    2,
			want: 6,
		},

		{
			name: "example 2",
			nums: []int{1, 2, 1, 2, 1, 2, 1, 2},
			k:    1,
			want: 2,
		},

		{
			name: "example 3",
			nums: []int{5, 5, 5, 5, 5, 5, 5},
			k:    4,
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := maxSubarrayLength(tt.nums, tt.k)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
