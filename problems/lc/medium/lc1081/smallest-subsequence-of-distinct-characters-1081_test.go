package lc1081

import (
	"reflect"
	"testing"
)

func TestSmallestSubsequence(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{

		{
			name: "example 1",
			s:    "bcabc",
			want: "abc",
		},

		{
			name: "example 2",
			s:    "cbacdcbc",
			want: "acdb",
		},

		{
			name: "example 3",
			s:    "baababaaaaababbbbbbaababaababa",
			want: "ab",
		},
		{
			name: "example 4",
			s:    "bcbcbcababa",
			want: "bca",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := smallestSubsequence(tt.s)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
