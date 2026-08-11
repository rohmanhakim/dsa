package lc2996

import (
	"reflect"
	"testing"
)

func TestMissingInteger(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{

		{
			name: "example 1",
			nums: []int{1, 2, 3, 2, 5},
			want: 6,
		},

		{
			name: "example 2",
			nums: []int{3, 4, 5, 1, 12, 14, 13},
			want: 15,
		},

		{
			name: "example 3",
			nums: []int{29, 30, 31, 32, 33, 34, 35, 36, 37},
			want: 297,
		},

		{
			name: "example 4",
			nums: []int{18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 9},
			want: 253,
		},

		{
			name: "example 5",
			nums: []int{46, 8, 2, 4, 1, 4, 10, 2, 4, 10, 2, 5, 7, 3, 1},
			want: 47,
		},

		{
			name: "example 6",
			nums: []int{12, 10},
			want: 13,
		},

		{
			name: "example 7",
			nums: []int{28, 29},
			want: 57,
		},

		{
			name: "example 8",
			nums: []int{28},
			want: 29,
		},

		{
			name: "example 9",
			nums: []int{1, 2, 3, 9, 2, 10, 8, 3, 10, 2},
			want: 6,
		},

		{
			name: "example 10",
			nums: []int{19, 47, 6, 8, 9, 20},
			want: 21,
		},

		{
			name: "example 11",
			nums: []int{3, 4, 7, 6, 6, 5, 8, 2, 8, 9, 2, 6},
			want: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := missingInteger(tt.nums)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
