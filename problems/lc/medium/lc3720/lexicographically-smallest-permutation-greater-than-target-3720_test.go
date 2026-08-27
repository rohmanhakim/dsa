package lc3720

import (
	"reflect"
	"testing"
)

func TestLexGreaterPermutation(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		target string
		want   string
	}{

		{
			name:   "example 1",
			s:      "abc",
			target: "bba",
			want:   "bca",
		},

		{
			name:   "example 2",
			s:      "leet",
			target: "code",
			want:   "eelt",
		},

		{
			name:   "example 3",
			s:      "baba",
			target: "bbaa",
			want:   "",
		},

		{
			name:   "example 4",
			s:      "z",
			target: "a",
			want:   "z",
		},

		{
			name:   "example 5",
			s:      "ab",
			target: "ab",
			want:   "ba",
		},

		{
			name:   "example 6",
			s:      "baab",
			target: "baab",
			want:   "baba",
		},

		{
			name:   "example 7",
			s:      "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
			target: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			want:   "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
		},

		{
			name:   "example 8",
			s:      "aab",
			target: "aab",
			want:   "aba",
		},

		{
			name:   "example 9",
			s:      "aab",
			target: "aba",
			want:   "baa",
		},

		{
			name:   "example 10",
			s:      "aaab",
			target: "abaa",
			want:   "baaa",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := lexGreaterPermutation(tt.s, tt.target)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
