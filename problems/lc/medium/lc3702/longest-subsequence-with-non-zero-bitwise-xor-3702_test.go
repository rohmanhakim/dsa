package lc3702

import (
	"reflect"
	"testing"
)

func TestLongestSubsequence(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{

		{
			name: "example 1",
			nums: []int{1, 2, 3},
			want: 2,
		},

		{
			name: "example 2",
			nums: []int{2, 3, 4},
			want: 3,
		},

		{
			name: "example 3",
			nums: []int{1},
			want: 1,
		},

		{
			name: "example 4",
			nums: []int{0},
			want: 0,
		},

		{
			name: "example 4",
			nums: []int{2},
			want: 1,
		},

		{
			name: "example 5",
			nums: []int{100000},
			want: 1,
		},

		{
			name: "example 6",
			nums: []int{0, 1, 2},
			want: 3,
		},

		{
			name: "example 7",
			nums: []int{7, 6, 1, 9},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := longestSubsequence(tt.nums)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
