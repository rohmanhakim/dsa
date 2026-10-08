package lc1021

import (
	"reflect"
	"testing"
)

func TestRemoveOuterParentheses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{

		{
			name: "example 1",
			s:    "(()())(())",
			want: "()()()",
		},

		{
			name: "example 2",
			s:    "(()())(())(()(()))",
			want: "()()()()(())",
		},

		{
			name: "example 3",
			s:    "()()",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := removeOuterParentheses(tt.s)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
