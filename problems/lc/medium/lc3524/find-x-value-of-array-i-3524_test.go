package lc3524

import (
	"reflect"
	"testing"
)

func TestResultArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int64
	}{

		{
			name: "example 1",
			nums: []int{1, 2, 3, 4, 5},
			k:    3,
			want: []int64{9, 2, 4},
		},

		{
			name: "example 2",
			nums: []int{1, 2, 4, 8, 16, 32},
			k:    4,
			want: []int64{18, 1, 2, 0},
		},

		{
			name: "example 3",
			nums: []int{1, 1, 2, 1, 1},
			k:    2,
			want: []int64{9, 6},
		},

		{
			name: "example 4",
			nums: []int{1},
			k:    2,
			want: []int64{0, 1},
		},

		{
			name: "example 5",
			nums: []int{78, 71, 26, 24, 81, 51, 46, 29, 45, 77, 83, 72, 79, 74, 8, 34, 49, 63, 69, 39, 81, 61, 88, 35, 59, 45, 21, 72, 98, 10, 48, 34, 3, 11, 16, 50, 53, 21, 83, 96, 53, 5, 94},
			k:    5,
			want: []int64{767, 46, 33, 49, 51},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := resultArray(tt.nums, tt.k)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
