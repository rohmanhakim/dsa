package lc1614

import (
	"reflect"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{

		{
			name: "example 1",
			s:    "(1+(2*3)+((8)/4))+1",
			want: 3,
		},

		{
			name: "example 2",
			s:    "(1)+((2))+(((3)))",
			want: 3,
		},

		{
			name: "example 3",
			s:    "()(())((()()))",
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := maxDepth(tt.s)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
