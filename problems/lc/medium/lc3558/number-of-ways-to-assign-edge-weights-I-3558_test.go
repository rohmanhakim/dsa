package lc3558

import (
	"reflect"
	"testing"
)

func TestNumberOfWaysToAssignEdgeWeightsI(t *testing.T) {
	tests := []struct {
		name  string
		edges [][]int
		want  int
	}{
		{
			name:  "test 1",
			edges: [][]int{{1, 2}},
			want:  1,
		},
		{
			name:  "test 2",
			edges: [][]int{{1, 2}, {1, 3}, {3, 4}, {3, 5}},
			want:  2,
		},
		{
			name:  "test 3",
			edges: [][]int{{1, 2}, {2, 3}, {3, 4}},
			want:  4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := assignEdgeWeights(tt.edges)
			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %d want %d\n", out, tt.want)
			}
		})
	}
}
