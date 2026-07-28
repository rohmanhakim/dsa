package lc3517

import (
	"reflect"
	"testing"
)

func TestSmallestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{

		{
			name: "example 1",
			s:    "z",
			want: "z",
		},

		{
			name: "example 2",
			s:    "babab",
			want: "abbba",
		},

		{
			name: "example 3",
			s:    "daccad",
			want: "acddca",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := smallestPalindrome(tt.s)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
