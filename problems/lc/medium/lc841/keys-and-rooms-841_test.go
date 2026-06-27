package lc841

import (
	"reflect"
	"testing"
)

func TestKeysAndRooms(t *testing.T) {
	tests := []struct {
		name  string
		rooms [][]int
		want  bool
	}{
		{
			name: "test 1",
			rooms: [][]int{
				{1}, {2}, {3}, {},
			},
			want: true,
		},
		{
			name: "test 2",
			rooms: [][]int{
				{1, 3},
				{3, 0, 1},
				{2},
				{0},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := canVisitAllRooms(tt.rooms)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
