package lc3312

import (
	"reflect"
	"testing"
)

func TestGcdValues(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		queries []int64
		want    []int
	}{

		{
			name:    "example 1",
			nums:    []int{2, 3, 4},
			queries: []int64{0, 2, 2},
			want:    []int{1, 2, 2},
		},

		{
			name:    "example 2",
			nums:    []int{4, 4, 2, 1},
			queries: []int64{5, 3, 1, 0},
			want:    []int{4, 2, 1, 1},
		},

		{
			name:    "example 3",
			nums:    []int{2, 2},
			queries: []int64{0, 0},
			want:    []int{2, 2},
		},
		{
			name:    "example 4",
			nums:    []int{5, 10, 4},
			queries: []int64{0, 1, 2},
			want:    []int{1, 2, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := gcdValues(tt.nums, tt.queries)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
