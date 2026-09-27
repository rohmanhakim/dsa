package lc1190

import (
	"reflect"
	"testing"
)

func TestReverseParentheses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{

		{
			name: "example 1",
			s:    "(abcd)",
			want: "dcba",
		},

		{
			name: "example 2",
			s:    "(u(love)i)",
			want: "iloveu",
		},

		{
			name: "example 3",
			s:    "(ed(et(oc))el)",
			want: "leetcode",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := reverseParentheses(tt.s)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
