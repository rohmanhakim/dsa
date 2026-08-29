package lc2948

import (
	"reflect"
	"testing"
)

func TestLexicographicallySmallestArray(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		limit int
		want  []int
	}{

		{
			name:  "example 1",
			nums:  []int{1, 5, 3, 9, 8},
			limit: 2,
			want:  []int{1, 3, 5, 8, 9},
		},

		{
			name:  "example 2",
			nums:  []int{1, 7, 6, 18, 2, 1},
			limit: 3,
			want:  []int{1, 6, 7, 18, 1, 2},
		},

		{
			name:  "example 3",
			nums:  []int{1, 7, 28, 19, 10},
			limit: 3,
			want:  []int{1, 7, 28, 19, 10},
		},

		{
			name:  "example 4",
			nums:  []int{14, 1, 12},
			limit: 3,
			want:  []int{12, 1, 14},
		},

		{
			name:  "example 5",
			nums:  []int{1, 7, 6, 18, 2, 1, 16, 3, 5, 4},
			limit: 3,
			want:  []int{1, 1, 2, 16, 3, 4, 18, 5, 6, 7},
		},

		{
			name:  "example 6",
			nums:  []int{6, 7, 2, 5, 4, 9, 50, 37, 48, 8, 6, 2, 5, 4, 3},
			limit: 5,
			want:  []int{2, 2, 3, 4, 4, 5, 48, 37, 50, 5, 6, 6, 7, 8, 9},
		},

		{
			name:  "example 7",
			nums:  []int{1000000000},
			limit: 1,
			want:  []int{1000000000},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := lexicographicallySmallestArray(tt.nums, tt.limit)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
