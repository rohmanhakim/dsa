package lc1464

import (
	"reflect"
	"testing"
)

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{

		{
			name: "example 1",
			nums: []int{3, 4, 5, 2},
			want: 12,
		},

		{
			name: "example 2",
			nums: []int{1, 5, 4, 5},
			want: 16,
		},

		{
			name: "example 3",
			nums: []int{3, 7},
			want: 12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := maxProduct(tt.nums)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
