package lc3090

import (
	"reflect"
	"testing"
)

func TestMaximumLengthSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{

		{
			name: "example 1",
			s:    "bcbbbcba",
			want: 4,
		},

		{
			name: "example 2",
			s:    "aaaa",
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := maximumLengthSubstring(tt.s)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
