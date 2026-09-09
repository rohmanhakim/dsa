package lc3871

import (
	"reflect"
	"testing"
)

func TestCountCommas(t *testing.T) {
	tests := []struct {
		name string
		n    int64
		want int64
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

		{
			name: "example 3",
			n:    100001,
			want: 99002,
		},

		{
			name: "example 4",
			n:    1000_000,
			want: 999002,
		},

		{
			name: "example 5",
			n:    1_004_590,
			want: 1008182,
		},

		{
			name: "example 6",
			n:    1_409_752_114,
			want: 3228255345,
		},

		{
			name: "example 7",
			n:    1000_000_000_000_000,
			want: 3998998998999005,
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
