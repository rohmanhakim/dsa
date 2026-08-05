package lc3310

import (
	"reflect"
	"testing"
)

func TestRemainingMethods(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		k           int
		invocations [][]int
		want        []int
	}{

		{
			name:        "example 1",
			n:           4,
			k:           1,
			invocations: [][]int{{1, 2}, {0, 1}, {3, 2}},
			want:        []int{0, 1, 2, 3},
		},

		{
			name:        "example 2",
			n:           5,
			k:           0,
			invocations: [][]int{{1, 2}, {0, 2}, {0, 1}, {3, 4}},
			want:        []int{3, 4},
		},

		{
			name:        "example 3",
			n:           3,
			k:           2,
			invocations: [][]int{{1, 2}, {0, 1}, {2, 0}},
			want:        []int{},
		},

		{
			name:        "example 4",
			n:           3,
			k:           0,
			invocations: [][]int{{2, 0}},
			want:        []int{0, 1, 2},
		},

		{
			name:        "example 5",
			n:           1,
			k:           0,
			invocations: [][]int{},
			want:        []int{},
		},

		{
			name:        "example 6",
			n:           2,
			k:           0,
			invocations: [][]int{},
			want:        []int{1},
		},

		{
			name:        "example 7",
			n:           3,
			k:           0,
			invocations: [][]int{{0, 2}, {1, 0}, {2, 1}, {0, 1}},
			want:        []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := remainingMethods(tt.n, tt.k, tt.invocations)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
