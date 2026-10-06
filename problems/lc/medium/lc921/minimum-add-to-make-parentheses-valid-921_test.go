package lc921

import (
	"reflect"
	"testing"
)

func TestMinAddToMakeValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{

		{
			name: "example 1",
			s:    "())",
			want: 1,
		},

		{
			name: "example 2",
			s:    "(((",
			want: 3,
		},

		{
			name: "example 3",
			s:    "(())())))",
			want: 3,
		},

		{
			name: "example 4",
			s:    "()))((",
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := minAddToMakeValid(tt.s)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
