package lc3870

import (
	"reflect"
	"testing"
)

func TestCountCommas(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{

		{
			name: "example 1",
			n:    1002,
			want: 3,
		},

		{
			name: "example 2",
			n:    998,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := countCommas(tt.n)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
