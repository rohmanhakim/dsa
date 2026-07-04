package lc2492

import (
	"reflect"
	"testing"
)

func TestMinScore(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		roads [][]int
		want  int
	}{
		{
			name:  "example 1",
			n:     4,
			roads: [][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}},
			want:  5,
		},
		{
			name:  "example 2",
			n:     4,
			roads: [][]int{{1, 2, 2}, {1, 3, 4}, {3, 4, 7}},
			want:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := minScore(tt.n, tt.roads)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
