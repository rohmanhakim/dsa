package lc3904

import (
	"reflect"
	"testing"
)

func TestFirstStableIndex(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{

		{
			name: "example 1",
			nums: []int{5, 0, 1, 4},
			k:    3,
			want: 3,
		},

		{
			name: "example 2",
			nums: []int{3, 2, 1},
			k:    1,
			want: -1,
		},

		{
			name: "example 3",
			nums: []int{0},
			k:    0,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := firstStableIndex(tt.nums, tt.k)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
