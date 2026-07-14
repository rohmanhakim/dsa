package lc150

import (
	"reflect"
	"testing"
)

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		name   string
		tokens []string
		want   int
	}{

		{
			name:   "example 1",
			tokens: []string{"2", "1", "+", "3", "*"},
			want:   9,
		},

		{
			name:   "example 2",
			tokens: []string{"4", "13", "5", "/", "+"},
			want:   6,
		},

		{
			name:   "example 3",
			tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			want:   22,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := evalRPN(tt.tokens)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
