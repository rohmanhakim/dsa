package lc1510

import (
	"reflect"
	"testing"
)

func TestWinnerSquareGame(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{

		{
			name: "example 1",
			n:    1,
			want: true,
		},

		{
			name: "example 2",
			n:    2,
			want: false,
		},

		{
			name: "example 3",
			n:    4,
			want: true,
		},

		{
			name: "example 4",
			n:    13,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := winnerSquareGame(tt.n)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
