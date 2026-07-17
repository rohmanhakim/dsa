package lc300

import (
	"reflect"
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{

		{
			name: "example 1",
			nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			want: 4,
		},

		{
			name: "example 2",
			nums: []int{0, 1, 0, 3, 2, 3},
			want: 4,
		},

		{
			name: "example 3",
			nums: []int{7, 7, 7, 7, 7, 7, 7},
			want: 1,
		},

		{
			name: "example 4",
			nums: []int{4, 10, 4, 3, 8, 9},
			want: 3,
		},

		{
			name: "example 5",
			nums: []int{10, 9, 2, 5, 3, 4},
			want: 3,
		},

		{
			name: "example 6",
			nums: []int{1, 3, 6, 7, 9, 4, 10, 5, 6},
			want: 6,
		},

		{
			name: "example 7",
			nums: []int{2, 15, 3, 7, 8, 6, 18},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := lengthOfLIS(tt.nums)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
