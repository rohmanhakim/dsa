package lc3731

import (
	"reflect"
	"testing"
)

func TestFindMissingElements(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{

		{
			name: "example 1",
			nums: []int{1, 4, 2, 5},
			want: []int{3},
		},

		{
			name: "example 2",
			nums: []int{7, 8, 6, 9},
			want: []int{},
		},

		{
			name: "example 3",
			nums: []int{5, 1},
			want: []int{2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := findMissingElements(tt.nums)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
