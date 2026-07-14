package lc3336

import (
	"reflect"
	"testing"
)

func TestSubsequencePairCount(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{

		{
			name: "example 1",
			nums: []int{1, 2, 3, 4},
			want: 10,
		},

		{
			name: "example 2",
			nums: []int{10, 20, 30},
			want: 2,
		},

		{
			name: "example 3",
			nums: []int{1, 1, 1, 1},
			want: 50,
		},
		{
			name: "example 4",
			nums: []int{26, 30, 27, 20, 25, 22, 28, 28, 20, 27, 20, 20, 26, 28, 30, 24, 24, 20, 25, 29, 29, 23, 26, 20, 26, 24, 28, 20, 27, 30},
			want: 437617923,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := subsequencePairCount(tt.nums)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
