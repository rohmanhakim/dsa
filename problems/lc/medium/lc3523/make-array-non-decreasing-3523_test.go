package lc3523

import (
	"reflect"
	"testing"
)

func TestMaximumPossibleSize(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{

		{
			name: "example 1",
			nums: []int{4, 2, 5, 3, 5},
			want: 3,
		},

		{
			name: "example 2",
			nums: []int{1, 2, 3},
			want: 3,
		},

		{
			name: "example 3",
			nums: []int{5, 4, 3, 2, 1},
			want: 1,
		},

		{
			name: "example 4",
			nums: []int{1, 2, 3, 4, 5, 6, 8, 7},
			want: 7,
		},

		{
			name: "example 5",
			nums: []int{19, 80, 63, 74},
			want: 2,
		},

		{
			name: "example 6",
			nums: []int{4, 3, 2, 1, 5, 4, 5, 3},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := maximumPossibleSize(tt.nums)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
