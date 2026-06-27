package lc1926

import (
	"reflect"
	"testing"
)

func TestNearestExitFromEntranceInMaze(t *testing.T) {
	tests := []struct {
		name     string
		maze     [][]byte
		entrance []int
		want     int
	}{
		{
			name:     "example 1",
			maze:     [][]byte{{'+', '+', '.', '+'}, {'.', '.', '.', '+'}, {'+', '+', '+', '.'}},
			entrance: []int{1, 2},
			want:     1,
		},
		{
			name:     "example 2",
			maze:     [][]byte{{'+', '+', '+'}, {'.', '.', '.'}, {'+', '+', '+'}},
			entrance: []int{1, 0},
			want:     2,
		},
		{
			name:     "example 3",
			maze:     [][]byte{{'.', '+'}},
			entrance: []int{0, 0},
			want:     -1,
		},
		{
			name:     "example 4",
			maze:     [][]byte{{'+', '.', '+', '+', '+', '+', '+'}, {'+', '.', '+', '.', '.', '.', '+'}, {'+', '.', '+', '.', '+', '.', '+'}, {'+', '.', '.', '.', '.', '.', '+'}, {'+', '+', '+', '+', '.', '+', '.'}},
			entrance: []int{0, 1},
			want:     7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := nearestExit(tt.maze, tt.entrance)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
