package lc3622

import (
	"reflect"
	"testing"
)

func TestCheckDivisibility(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{

		{
			name: "example 1",
			n:    99,
			want: true,
		},

		{
			name: "example 2",
			n:    23,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := checkDivisibility(tt.n)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
