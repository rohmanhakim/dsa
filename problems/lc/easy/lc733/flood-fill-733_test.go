package lc733

import (
	"reflect"
	"testing"
)

func TestFloodFill(t *testing.T) {
	tests := []struct {
		name  string
		image [][]int
		sr    int
		sc    int
		color int
		want  [][]int
	}{

		{
			name:  "example 1",
			image: [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
			sr:    1,
			sc:    1,
			color: 2,
			want:  [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
		},

		{
			name:  "example 2",
			image: [][]int{{0, 0, 0}, {0, 0, 0}},
			sr:    0,
			sc:    0,
			color: 0,
			want:  [][]int{{0, 0, 0}, {0, 0, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := floodFill(tt.image, tt.sr, tt.sc, tt.color)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
