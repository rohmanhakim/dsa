package medium

import (
	"reflect"
	"testing"
)

func TestNumberOfProvinces(t *testing.T) {
	tests := []struct {
		name        string
		isConnected [][]int
		want        int
	}{
		{
			name:        "test 1",
			isConnected: [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}},
			want:        2,
		},

		{
			name:        "test 2",
			isConnected: [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			want:        3,
		},

		{
			name:        "test 3",
			isConnected: [][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}},
			want:        1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := findCircleNum(tt.isConnected)
			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
