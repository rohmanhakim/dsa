package lc3514

import (
	"reflect"
	"testing"
)

func TestUniqueXorTriplets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{

		{
			name: "example 1",
			nums: []int{1, 3},
			want: 2,
		},

		{
			name: "example 2",
			nums: []int{6, 7, 8, 9},
			want: 4,
		},

		{
			name: "example 3",
			nums: []int{1280, 476, 379},
			want: 4,
		},

		{
			name: "example 4",
			nums: []int{1334, 495, 1112, 864, 133, 1200, 400, 1484},
			want: 64,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := uniqueXorTriplets(tt.nums)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
