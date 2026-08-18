package lc3471

import (
	"reflect"
	"testing"
)

func TestLargestInteger(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{

		{
			name: "example 1",
			nums: []int{3, 9, 2, 1, 7},
			k:    3,
			want: 7,
		},

		{
			name: "example 2",
			nums: []int{3, 9, 7, 2, 1, 7},
			k:    4,
			want: 3,
		},

		{
			name: "example 3",
			nums: []int{0, 0},
			k:    1,
			want: -1,
		},

		{
			name: "example 4",
			nums: []int{0, 0},
			k:    2,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := largestInteger(tt.nums, tt.k)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
