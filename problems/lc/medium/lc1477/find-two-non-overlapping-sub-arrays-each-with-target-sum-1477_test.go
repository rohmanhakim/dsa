package lc1477

import (
	"reflect"
	"testing"
)

func TestMinSumOfLengths(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		target int
		want   int
	}{

		{
			name:   "example 1",
			arr:    []int{3, 2, 2, 4, 3},
			target: 3,
			want:   2,
		},

		{
			name:   "example 2",
			arr:    []int{7, 3, 4, 7},
			target: 7,
			want:   2,
		},

		{
			name:   "example 3",
			arr:    []int{4, 3, 2, 6, 2, 3, 4},
			target: 6,
			want:   -1,
		},

		{
			name:   "example 4",
			arr:    []int{7, 3, 4, 5},
			target: 7,
			want:   3,
		},

		{
			name:   "example 5",
			arr:    []int{1, 6, 1},
			target: 7,
			want:   -1,
		},

		{
			name:   "example 6",
			arr:    []int{3, 1, 1, 1, 5, 1, 2, 1},
			target: 3,
			want:   3,
		},

		{
			name:   "example 7",
			arr:    []int{1, 1, 1, 2, 2, 2, 4, 4},
			target: 6,
			want:   6,
		},

		{
			name:   "example 8",
			arr:    []int{1, 2, 2, 3, 2, 6, 7, 2, 1, 4, 8},
			target: 5,
			want:   4,
		},

		{
			name:   "example 9",
			arr:    []int{2, 1, 3, 3, 2, 3, 1},
			target: 6,
			want:   5,
		},

		{
			name:   "example 10",
			arr:    []int{78, 18, 1, 94, 1, 1, 1, 29, 58, 3, 4, 1, 2, 56, 17, 19, 4, 1, 63, 2, 16, 11, 1, 1, 2, 1, 25, 62, 10, 69, 12, 7, 1, 6, 2, 92, 4, 1, 61, 7, 26, 1, 1, 1, 67, 26, 2, 2, 70, 25, 2, 68, 13, 4, 11, 1, 34, 14, 7, 37, 4, 1, 12, 51, 25, 2, 4, 3, 56, 21, 7, 8, 5, 93, 1, 1, 2, 55, 14, 25, 1, 1, 1, 89, 6, 1, 1, 24, 22, 50, 1, 28, 9, 51, 9, 88, 1, 7, 1, 30, 32, 18, 12, 3, 2, 18, 10, 4, 11, 43, 6, 5, 93, 2, 2, 68, 18, 11, 47, 33, 17, 27, 56, 13, 1, 2, 29, 1, 17, 1, 10, 15, 18, 3, 1, 86, 7, 4, 16, 45, 3, 29, 2, 1, 1, 31, 19, 18, 16, 12, 1, 56, 4, 35, 1, 1, 36, 59, 1, 1, 16, 58, 18, 4, 1, 43, 31, 15, 6, 1, 1, 6, 49, 27, 12, 1, 2, 80, 14, 2, 1, 21, 32, 18, 15, 11, 59, 10, 1, 14, 3, 3, 7, 15, 4, 55, 4, 1, 12, 4, 1, 1, 53, 37, 2, 5, 72, 3, 6, 10, 3, 3, 83, 8, 1, 5},
			target: 97,
			want:   5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := minSumOfLengths(tt.arr, tt.target)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
