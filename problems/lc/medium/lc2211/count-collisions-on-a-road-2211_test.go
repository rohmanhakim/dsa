package lc2211

import (
	"reflect"
	"testing"
)

func TestCountCollisions(t *testing.T) {
	tests := []struct {
		name       string
		directions string
		want       int
	}{

		{
			name:       "example 1",
			directions: "RLRSLL",
			want:       5,
		},

		{
			name:       "example 2",
			directions: "LLRR",
			want:       0,
		},

		{
			name:       "example 3",
			directions: "RRL",
			want:       3,
		},

		{
			name:       "example 4",
			directions: "RSLLRSSL",
			want:       5,
		},

		{
			name:       "example 5",
			directions: "SSRSSRLLRSLLRSRSSRLRRRRLLRRLSSRR",
			want:       20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := countCollisions(tt.directions)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
