package lc32

import (
	"reflect"
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{

		{
			name: "example 1",
			s:    "(()",
			want: 2,
		},

		{
			name: "example 2",
			s:    ")()())",
			want: 4,
		},

		{
			name: "example 3",
			s:    "",
			want: 0,
		},

		{
			name: "example 4",
			s:    "()(()",
			want: 2,
		},

		{
			name: "example 5",
			s:    "()(())",
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := longestValidParentheses(tt.s)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
