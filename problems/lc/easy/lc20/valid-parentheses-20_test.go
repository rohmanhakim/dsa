package lc20

import (
	"reflect"
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "example 1",
			s:    "()",
			want: true,
		},

		{
			name: "example 2",
			s:    "()[]{}",
			want: true,
		},

		{
			name: "example 3",
			s:    "(]",
			want: false,
		},

		{
			name: "example 4",
			s:    "([])",
			want: true,
		},

		{
			name: "example 5",
			s:    "([)]",
			want: false,
		},
		{
			name: "example 6",
			s:    "[",
			want: false,
		},
		{
			name: "example 7",
			s:    "((",
			want: false,
		},
		{
			name: "example 8",
			s:    "){",
			want: false,
		},
		{
			name: "example 9",
			s:    "))",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := isValid(tt.s)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
