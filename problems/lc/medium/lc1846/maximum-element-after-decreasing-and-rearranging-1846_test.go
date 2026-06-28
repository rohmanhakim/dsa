package lc1846

import (
	"reflect"
	"testing"
)

func TestMaximumElementAfterDecreasingandRearranging(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{
			name: "example 1",
			arr:  []int{2, 2, 1, 2, 1},
			want: 2,
		},
		{
			name: "example 2",
			arr:  []int{100, 1, 1000},
			want: 3,
		},
		{
			name: "example 3",
			arr:  []int{1, 2, 3, 4, 5},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := maximumElementAfterDecrementingAndRearrangingNoSort(tt.arr)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
