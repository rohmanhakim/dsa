package lc219

import (
	"reflect"
	"testing"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want bool
	}{

		{
			name: "example 1",
			nums: []int{1, 2, 3, 1},
			k:    3,
			want: true,
		},

		{
			name: "example 2",
			nums: []int{1, 0, 1, 1},
			k:    1,
			want: true,
		},

		{
			name: "example 3",
			nums: []int{1, 2, 3, 1, 2, 3},
			k:    2,
			want: false,
		},

		{
			name: "example 4",
			nums: []int{1, 2, 1},
			k:    0,
			want: false,
		},

		{
			name: "example 5",
			nums: []int{1, 4, 2, 3, 1, 2},
			k:    3,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := containsNearbyDuplicate(tt.nums, tt.k)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
