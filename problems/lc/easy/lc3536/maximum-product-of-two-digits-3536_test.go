package lc3536

import (
	"reflect"
	"testing"
)

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{

		{
			name: "example 1",
			n:    31,
			want: 3,
		},

		{
			name: "example 2",
			n:    22,
			want: 4,
		},

		{
			name: "example 3",
			n:    124,
			want: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := maxProduct(tt.n)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
