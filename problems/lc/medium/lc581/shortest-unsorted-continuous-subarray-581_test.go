package lc581

import (
	"reflect"
	"testing"
)

func TestFindUnsortedSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{

		{
			name: "example 1",
			nums: []int{2, 6, 4, 8, 10, 9, 15},
			want: 5,
		},

		{
			name: "example 2",
			nums: []int{1, 2, 3, 4},
			want: 0,
		},

		{
			name: "example 3",
			nums: []int{1},
			want: 0,
		},

		{
			name: "example 4",
			nums: []int{1, 3, 2, 2, 2},
			want: 4,
		},

		{
			name: "example 5",
			nums: []int{2, 1},
			want: 2,
		},

		{
			name: "example 6",
			nums: []int{2, 1, 3, 2},
			want: 4,
		},

		{
			name: "example 7",
			nums: []int{2, 1, 3, 4},
			want: 2,
		},

		{
			name: "example 8",
			nums: []int{1, 3, 2, 4, 5},
			want: 2,
		},

		{
			name: "example 9",
			nums: []int{1, 2, 3, 3, 3},
			want: 0,
		},

		{
			name: "example 10",
			nums: []int{1, 3, 2, 3, 3},
			want: 2,
		},

		{
			name: "example 11",
			nums: []int{2, 3, 3, 2, 4},
			want: 3,
		},

		{
			name: "example 12",
			nums: []int{5, 4, 3, 2, 1},
			want: 5,
		},
		{
			name: "example 13",
			nums: []int{1, 1},
			want: 0,
		},

		{
			name: "example 14",
			nums: []int{3, 2, 3, 2, 4},
			want: 4,
		},

		{
			name: "example 15",
			nums: []int{1, 2, 5, 3, 4},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := findUnsortedSubarray(tt.nums)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
