package lc3718

import (
	"reflect"
	"testing"
)

func TestMissingMultiple(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{

		{
			name: "example 1",
			nums: []int{8, 2, 3, 4, 6},
			k:    2,
			want: 10,
		},

		{
			name: "example 2",
			nums: []int{1, 4, 7, 10, 15},
			k:    5,
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := missingMultiple(tt.nums, tt.k)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
