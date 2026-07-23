package lc213

import (
	"reflect"
	"testing"
)

func TestRob(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{

		{
			name: "example 1",
			nums: []int{2, 3, 2},
			want: 3,
		},

		{
			name: "example 2",
			nums: []int{1, 2, 3, 1},
			want: 4,
		},

		{
			name: "example 3",
			nums: []int{1, 2, 3},
			want: 3,
		},
		{
			name: "example 4",
			nums: []int{200, 3, 140, 20, 10},
			want: 340,
		},

		{
			name: "example 5",
			nums: []int{1, 1, 3, 6, 7, 10, 7, 1, 8, 5, 9, 1, 4, 4, 3},
			want: 41,
		},

		{
			name: "example 6",
			nums: []int{0},
			want: 0,
		},

		{
			name: "example 7",
			nums: []int{1},
			want: 1,
		},

		{
			name: "example 8",
			nums: []int{0, 0},
			want: 0,
		},

		{
			name: "example 9",
			nums: []int{1, 2},
			want: 2,
		},

		{
			name: "example 10",
			nums: []int{2, 2, 4, 3, 2, 5},
			want: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := rob(tt.nums)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
