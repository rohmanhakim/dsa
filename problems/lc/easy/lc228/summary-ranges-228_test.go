package lc228

import (
	"reflect"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []string
	}{
		{
			name: "example 1",
			nums: []int{0, 1, 2, 4, 5, 7},
			want: []string{"0->2", "4->5", "7"},
		},
		{
			name: "example 2",
			nums: []int{0, 2, 3, 4, 6, 8, 9},
			want: []string{"0", "2->4", "6", "8->9"},
		},
		{
			name: "example 3",
			nums: []int{},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := summaryRanges(tt.nums)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
