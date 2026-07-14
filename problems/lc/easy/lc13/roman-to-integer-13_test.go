package lc13

import (
	"reflect"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{

		{
			name: "example 1",
			s:    "III",
			want: 3,
		},

		{
			name: "example 2",
			s:    "LVIII",
			want: 58,
		},

		{
			name: "example 3",
			s:    "MCMXCIV",
			want: 1994,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := romanToInt(tt.s)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
