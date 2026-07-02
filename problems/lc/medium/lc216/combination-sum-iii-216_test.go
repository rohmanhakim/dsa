package lc216

import (
	"reflect"
	"testing"
)

func TestCombinationSumIII(t *testing.T) {
	tests := []struct {
		name string
		k    int
		n    int
		want [][]int
	}{
		{
			name: "example 1",
			k:    3,
			n:    7,
			want: [][]int{{1, 2, 4}},
		},
		{
			name: "example 2",
			k:    3,
			n:    9,
			want: [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}},
		},
		{
			name: "example 3",
			k:    4,
			n:    1,
			want: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := combinationSum3(tt.k, tt.n)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
