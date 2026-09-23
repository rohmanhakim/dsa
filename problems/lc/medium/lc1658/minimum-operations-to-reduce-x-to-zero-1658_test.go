package lc1658

import (
	"reflect"
	"testing"
)

func TestMinOperations(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		x    int
		want int
	}{

		{
			name: "example 1",
			nums: []int{1, 1, 4, 2, 3},
			x:    5,
			want: 2,
		},

		{
			name: "example 2",
			nums: []int{5, 6, 7, 8, 9},
			x:    4,
			want: -1,
		},

		{
			name: "example 3",
			nums: []int{3, 2, 20, 1, 1, 3},
			x:    10,
			want: 5,
		},

		{
			name: "example 4",
			nums: []int{1, 2, 3},
			x:    6,
			want: 3,
		},

		{
			name: "example 5",
			nums: []int{1, 1},
			x:    5,
			want: -1,
		},

		{
			name: "example 6",
			nums: []int{1, 1, 1, 1, 1, 1, 1, 1},
			x:    3,
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := minOperations(tt.nums, tt.x)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
