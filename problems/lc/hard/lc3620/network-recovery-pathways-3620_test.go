package lc3620

import (
	"reflect"
	"testing"
)

func TestFindMaxPathScore(t *testing.T) {
	tests := []struct {
		name   string
		edges  [][]int
		online []bool
		k      int64
		want   int
	}{

		{
			name:   "example 1",
			edges:  [][]int{{0, 1, 5}, {1, 3, 10}, {0, 2, 3}, {2, 3, 4}},
			online: []bool{true, true, true, true},
			k:      10,
			want:   3,
		},
		{
			name:   "example 2",
			edges:  [][]int{{0, 1, 7}, {1, 4, 5}, {0, 2, 6}, {2, 3, 6}, {3, 4, 2}, {2, 4, 6}},
			online: []bool{true, true, true, false, true},
			k:      12,
			want:   6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := findMaxPathScore(tt.edges, tt.online, tt.k)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
