package lc2467

import (
	"reflect"
	"testing"
)

func TestMostProfitablePath(t *testing.T) {
	tests := []struct {
		name   string
		edges  [][]int
		bob    int
		amount []int
		want   int
	}{

		{
			name:   "example 1",
			edges:  [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}},
			bob:    3,
			amount: []int{-2, 4, 2, -4, 6},
			want:   6,
		},

		{
			name:   "example 2",
			edges:  [][]int{{0, 1}},
			bob:    1,
			amount: []int{-7280, 2350},
			want:   -7280,
		},

		{
			name:   "example 3",
			edges:  [][]int{{0, 1}, {1, 2}, {2, 3}},
			bob:    3,
			amount: []int{-5644, -6018, 1188, -8502},
			want:   -11662,
		},

		{
			name:   "example 4",
			edges:  [][]int{{0, 2}, {1, 4}, {1, 6}, {2, 4}, {3, 6}, {3, 7}, {5, 7}},
			bob:    4,
			amount: []int{-6896, -1216, -1208, -1108, 1606, -7704, -9212, -8258},
			want:   -34998,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := mostProfitablePath(tt.edges, tt.bob, tt.amount)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
