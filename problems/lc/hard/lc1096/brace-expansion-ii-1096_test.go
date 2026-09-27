package lc1096

import (
	"testing"

	"github.com/rohmanhakim/leetcode/lib"
)

func TestBraceExpansionII(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		want       []string
	}{

		{
			name:       "example 1",
			expression: "{a,b}{c,{d,e}}",
			want:       []string{"ac", "ad", "ae", "bc", "bd", "be"},
		},

		{
			name:       "example 2",
			expression: "{{a,z},a{b,c},{ab,z}}",
			want:       []string{"a", "ab", "ac", "z"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := braceExpansionII(tt.expression)

			if !lib.Equal1DUnordered(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
