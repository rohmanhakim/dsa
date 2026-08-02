package lc877

import (
	"reflect"
	"testing"
)

func TestStoneGame(t *testing.T) {
	tests := []struct {
		name  string
		piles []int
		want  bool
	}{
		{
			name:  "example 1",
			piles: []int{5, 3, 45},
			want:  true,
		},
		{
			name:  "example 2",
			piles: []int{3, 7, 2, 3},
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := stoneGame(tt.piles)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v, want %v", out, tt.want)
			}
		})
	}
}
