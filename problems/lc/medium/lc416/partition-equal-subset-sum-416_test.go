package lc416

import (
	"reflect"
	"testing"
)

func TestCanPartition(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{

		{
			name: "example 1",
			nums: []int{1, 5, 11, 5},
			want: true,
		},

		{
			name: "example 2",
			nums: []int{1, 2, 3, 5},
			want: false,
		},

		{
			name: "example 3",
			nums: []int{1},
			want: false,
		},

		{
			name: "example 4",
			nums: []int{1, 1},
			want: true,
		},

		{
			name: "example 5",
			nums: []int{1, 2},
			want: false,
		},

		{
			name: "example 6",
			nums: []int{1, 3, 7, 5},
			want: true,
		},

		{
			name: "example 7",
			nums: []int{1, 3, 3, 6, 1},
			want: true,
		},

		{
			name: "example 8",
			nums: []int{3, 3, 3, 4, 5},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := canPartition(tt.nums)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
