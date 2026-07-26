package lc930

import (
	"reflect"
	"testing"
)

func TestNumSubarraysWithSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		goal int
		want int
	}{

		{
			name: "example 1",
			nums: []int{1, 0, 1, 0, 1},
			goal: 2,
			want: 4,
		},

		{
			name: "example 2",
			nums: []int{0, 0, 0, 0, 0},
			goal: 0,
			want: 15,
		},

		{
			name: "example 3",
			nums: []int{0, 1, 1, 1, 1},
			goal: 3,
			want: 3,
		},

		{
			name: "example 4",
			nums: []int{0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
			goal: 0,
			want: 27,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := numSubarraysWithSum(tt.nums, tt.goal)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
