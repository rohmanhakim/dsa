package lc2685

import (
	"reflect"
	"testing"
)

func TestCountCompleteComponents(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
		want  int
	}{

		{
			name:  "example 1",
			n:     6,
			edges: [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}},
			want:  3,
		},
		{
			name:  "example 2",
			n:     6,
			edges: [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}, {3, 5}},
			want:  1,
		},
		{
			name:  "example 3",
			n:     3,
			edges: [][]int{{1, 0}, {2, 1}},
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := countCompleteComponents(tt.n, tt.edges)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
