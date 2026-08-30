package lc2091

import (
	"reflect"
	"testing"
)

func TestMinimumDeletions(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{

		{
			name: "example 1",
			nums: []int{2, 10, 7, 5, 4, 1, 8, 6},
			want: 5,
		},

		{
			name: "example 2",
			nums: []int{0, -4, 19, 1, 8, -2, -3, 5},
			want: 3,
		},

		{
			name: "example 3",
			nums: []int{101},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := minimumDeletions(tt.nums)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
