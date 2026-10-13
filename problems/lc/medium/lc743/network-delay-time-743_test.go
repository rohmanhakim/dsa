package lc743

import (
	"reflect"
	"testing"
)

func TestNetworkDelayTime(t *testing.T) {
	tests := []struct {
		name  string
		times [][]int
		n     int
		k     int
		want  int
	}{

		{
			name:  "example 1",
			times: [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}},
			n:     4,
			k:     2,
			want:  2,
		},

		{
			name:  "example 2",
			times: [][]int{{1, 2, 1}},
			n:     2,
			k:     1,
			want:  1,
		},

		{
			name:  "example 3",
			times: [][]int{{1, 2, 1}},
			n:     2,
			k:     2,
			want:  -1,
		},

		{
			name:  "example 4",
			times: [][]int{{1, 2, 1}, {1, 3, 4}, {2, 3, 2}, {4, 3, 1}},
			n:     4,
			k:     1,
			want:  -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := networkDelayTime(tt.times, tt.n, tt.k)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
