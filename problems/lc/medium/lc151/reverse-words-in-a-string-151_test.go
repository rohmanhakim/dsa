package lc151

import (
	"reflect"
	"testing"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{

		{
			name: "example 1",
			s:    "the sky is blue",
			want: "blue is sky the",
		},

		{
			name: "example 2",
			s:    "  hello world  ",
			want: "world hello",
		},

		{
			name: "example 3",
			s:    "a good   example",
			want: "example good a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := reverseWords(tt.s)

			if !reflect.DeepEqual(out, tt.want) && len(out) != len(tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
