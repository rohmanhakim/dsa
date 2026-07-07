package lc153

import (
	"reflect"
	"testing"
)

func TestFindMin(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{3, 4, 5, 1, 2},
			want: 1,
		},
		{
			name: "example 2",
			nums: []int{4, 5, 6, 7, 0, 1, 2},
			want: 0,
		},
		{
			name: "example 3",
			nums: []int{11, 13, 15, 17},
			want: 11,
		},
		{
			name: "example 4",
			nums: []int{7, 0, 1, 2},
			want: 0,
		},
		{
			name: "example 5",
			nums: []int{0},
			want: 0,
		},
		{
			name: "example 6",
			nums: []int{7, 8, 9, 10, 11, 20, 0},
			want: 0,
		},
		{
			name: "example 7",
			nums: []int{1, 3, 6, 9, 11, 22, 30},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := findMin(tt.nums)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
