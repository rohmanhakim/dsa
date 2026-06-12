package hard

import (
	"reflect"
	"testing"
)

func TestNumberOfWaysToAssignEdgeWeightsII(t *testing.T) {
	tests := []struct {
		name    string
		edges   [][]int
		queries [][]int
		want    []int
	}{
		{
			name:    "test 1",
			edges:   [][]int{{1, 2}},
			queries: [][]int{{1, 1}, {1, 2}},
			want:    []int{0, 1},
		},
		{
			name:    "test 2",
			edges:   [][]int{{1, 2}, {1, 3}, {3, 4}, {3, 5}},
			queries: [][]int{{1, 4}, {3, 4}, {2, 5}},
			want:    []int{2, 1, 4},
		},
		{
			name:    "test 3",
			edges:   [][]int{{3, 9}, {1, 8}, {2, 3}, {13, 19}, {4, 15}, {2, 7}, {4, 27}, {9, 23}, {9, 12}, {5, 26}, {3, 16}, {1, 2}, {2, 11}, {5, 17}, {1, 6}, {2, 4}, {12, 18}, {2, 20}, {9, 21}, {9, 25}, {9, 24}, {9, 14}, {6, 22}, {4, 5}, {7, 10}, {4, 13}},
			queries: [][]int{{24, 18}, {18, 12}, {4, 15}, {2, 15}, {1, 27}, {1, 22}},
			want:    []int{4, 1, 1, 2, 4, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := assignEdgeWeights(tt.edges, tt.queries)
			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %d want %d\n", out, tt.want)
			}
		})
	}
}
