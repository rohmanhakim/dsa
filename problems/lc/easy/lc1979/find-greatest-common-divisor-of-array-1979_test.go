package lc1979

import (
	"reflect"
	"testing"
)

func TestFindGCD(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{

		{
			name: "example 1",
			nums: []int{2, 5, 6, 9, 10},
			want: 2,
		},

		{
			name: "example 2",
			nums: []int{7, 5, 6, 8, 3},
			want: 1,
		},

		{
			name: "example 3",
			nums: []int{3, 3},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := findGCD(tt.nums)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
