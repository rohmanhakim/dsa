package lc1331

import (
	"reflect"
	"testing"
)

func TestArrayRankTransform(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "example 1",
			arr:  []int{40, 10, 20, 30},
			want: []int{4, 1, 2, 3},
		},

		{
			name: "example 2",
			arr:  []int{100, 100, 100},
			want: []int{1, 1, 1},
		},

		{
			name: "example 3",
			arr:  []int{37, 12, 28, 9, 100, 56, 80, 5, 12},
			want: []int{5, 3, 4, 2, 8, 6, 7, 1, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := arrayRankTransform(tt.arr)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
