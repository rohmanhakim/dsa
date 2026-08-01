package lc486

import (
	"reflect"
	"testing"
)

func TestPredictTheWinner(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{

		{
			name: "example 1",
			nums: []int{1, 5, 2},
			want: false,
		},

		{
			name: "example 2",
			nums: []int{1, 5, 233, 7},
			want: true,
		},

		{
			name: "example 3",
			nums: []int{0},
			want: true,
		},

		{
			name: "example 4",
			nums: []int{1, 2},
			want: true,
		},

		{
			name: "example 5",
			nums: []int{2, 4, 55, 6, 8},
			want: false,
		},

		{
			name: "example 6",
			nums: []int{1, 567, 1, 1},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := predictTheWinner(tt.nums)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
