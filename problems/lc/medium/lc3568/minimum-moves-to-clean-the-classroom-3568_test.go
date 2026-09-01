package lc3568

import (
	"reflect"
	"testing"
)

func TestMinMoves(t *testing.T) {
	tests := []struct {
		name      string
		classroom []string
		energy    int
		want      int
	}{

		{
			name:      "example 1",
			classroom: []string{"S.", "XL"},
			energy:    2,
			want:      2,
		},

		{
			name:      "example 2",
			classroom: []string{"LS", "RL"},
			energy:    4,
			want:      3,
		},

		{
			name:      "example 3",
			classroom: []string{"L.S", "RXL"},
			energy:    3,
			want:      -1,
		},

		{
			name:      "example 4",
			classroom: []string{"S.........", "........L.", "XXX.......", "..........", "LXX......."},
			energy:    50,
			want:      20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := minMoves(tt.classroom, tt.energy)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
