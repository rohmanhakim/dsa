package lc162

import (
	"reflect"
	"testing"
)

func TestFindPeakElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{

		{
			name: "example 1",
			nums: []int{1, 2, 3, 1},
			want: 2,
		},

		{
			name: "example 2",
			nums: []int{1, 2, 1, 3, 5, 6, 4},
			want: 5,
		},
		{
			name: "example 3",
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			want: 6,
		},
		{
			name: "example 4",
			nums: []int{7, 6, 5, 4, 3, 2, 1},
			want: 0,
		},
		{
			name: "example 5",
			nums: []int{1},
			want: 0,
		},
		{
			name: "example 6",
			nums: []int{2, 1},
			want: 0,
		},
		{
			name: "example 7",
			nums: []int{1, 2},
			want: 1,
		},
		{
			name: "example 8",
			nums: []int{1, 3, 2, 1},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := findPeakElement(tt.nums)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
