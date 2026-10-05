package lc856

import (
	"testing"
)

func TestScoreOfParentheses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{

		{
			name: "example 1",
			s:    "()",
			want: 1,
		},

		{
			name: "example 2",
			s:    "(())",
			want: 2,
		},

		{
			name: "example 3",
			s:    "()()",
			want: 2,
		},

		{
			name: "nested triple",
			s:    "((()))",
			want: 4,
		},

		{
			name: "nested quadruple",
			s:    "(((())))",
			want: 8,
		},

		{
			name: "three flat pairs",
			s:    "()()()",
			want: 3,
		},

		{
			name: "nested pair plus flat pair",
			s:    "(())()",
			want: 3,
		},
	}

	implementations := []struct {
		name string
		fn   func(string) int
	}{
		{"stack", scoreOfParentheses},
		{"alternative", scoreOfParenthesesAlternative},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, impl := range implementations {
				if out := impl.fn(tt.s); out != tt.want {
					t.Fatalf("%s: got %v want %v", impl.name, out, tt.want)
				}
			}
		})
	}
}
