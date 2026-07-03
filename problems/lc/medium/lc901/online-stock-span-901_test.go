package lc901

import (
	"reflect"
	"testing"
)

func ptr(i int) *int {
	return &i
}

func TestOnlineStockSpan(t *testing.T) {
	tests := []struct {
		name string
		ops  []string
		val  [][]int
		want []*int
	}{
		{
			name: "example 1",
			ops:  []string{"StockSpanner", "next", "next", "next", "next", "next", "next", "next"},
			val:  [][]int{{}, {100}, {80}, {60}, {70}, {60}, {75}, {85}},
			want: []*int{nil, ptr(1), ptr(1), ptr(1), ptr(2), ptr(1), ptr(4), ptr(6)},
		},
		{
			name: "example 2",
			ops:  []string{"StockSpanner", "next", "next", "next", "next", "next"},
			val:  [][]int{{}, {31}, {41}, {48}, {59}, {79}},
			want: []*int{nil, ptr(1), ptr(2), ptr(3), ptr(4), ptr(5)},
		},
		{
			name: "example 3",
			ops:  []string{"StockSpanner", "next", "next", "next", "next", "next"},
			val:  [][]int{{}, {29}, {91}, {62}, {76}, {51}},
			want: []*int{nil, ptr(1), ptr(2), ptr(1), ptr(2), ptr(1)},
		},
		{
			name: "example 4",
			ops:  []string{"StockSpanner", "next", "next", "next", "next", "next"},
			val:  [][]int{{}, {32}, {82}, {73}, {99}, {91}},
			want: []*int{nil, ptr(1), ptr(2), ptr(1), ptr(4), ptr(1)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s StockSpanner
			for i, op := range tt.ops {
				if op == "StockSpanner" {
					s = Constructor()
				} else {
					out := s.Next(tt.val[i][0])
					want := *tt.want[i]

					if !reflect.DeepEqual(out, want) {
						t.Fatalf("got %v want %v", out, want)
					}
				}
			}
		})
	}
}
