package lc3867

import (
	"reflect"
	"testing"
)

func TestGcdSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int64
	}{

		{
			name: "example 1",
			nums: []int{2, 6, 4},
			want: 2,
		},

		{
			name: "example 2",
			nums: []int{3, 6, 2, 8},
			want: 5,
		},

		{
			name: "example 3",
			nums: []int{9, 3, 6, 7, 1, 3, 8, 4, 5, 7, 1, 8, 3, 6, 2, 8, 4, 3},
			want: 9,
		},

		{
			name: "example 3",
			nums: []int{2324, 5675, 13, 277, 1456, 779, 34, 234, 9153, 82456, 34, 978},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := gcdSum(tt.nums)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
