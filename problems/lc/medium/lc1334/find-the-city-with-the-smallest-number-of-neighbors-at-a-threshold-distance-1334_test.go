package lc1334

import (
	"reflect"
	"testing"
)

func TestFindTheCity(t *testing.T) {
	tests := []struct {
		name string
		n int
		edges [][]int
		distanceThreshold int
		want int
	}{

		{
			name: "example 1",
			n: 4,
			edges: [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}},
			distanceThreshold: 4,
			want: 3,
		},

		{
			name: "example 2",
			n: 5,
			edges: [][]int{{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1}},
			distanceThreshold: 2,
			want: 0,
		},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := findTheCity(tt.n, tt.edges, tt.distanceThreshold)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
