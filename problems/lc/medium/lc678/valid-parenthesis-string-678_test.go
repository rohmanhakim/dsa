package lc678

import (
	"reflect"
	"testing"
)

func TestCheckValidString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{

		{
			name: "example 1",
			s:    "()",
			want: true,
		},

		{
			name: "example 2",
			s:    "(*)",
			want: true,
		},

		{
			name: "example 3",
			s:    "(*))",
			want: true,
		},

		{
			name: "example 4",
			s:    "(",
			want: false,
		},

		{
			name: "example 5",
			s:    "((((()(()()()*()(((((*)()*(**(())))))(())()())(((())())())))))))(((((())*)))()))(()((*()*(*)))(*)()",
			want: true,
		},

		{
			name: "example 6",
			s:    "(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())",
			want: false,
		},

		{
			name: "example 7",
			s:    "(((()))())))*))())()(**(((())(()(*()((((())))*())(())*(*(()(*)))()*())**((()(()))())(*(*))*))())",
			want: false,
		},

		{
			name: "example 8",
			s:    ")",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := checkValidString(tt.s)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
