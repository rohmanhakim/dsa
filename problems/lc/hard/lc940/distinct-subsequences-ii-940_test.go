package lc940

import (
	"reflect"
	"testing"
)

func TestDistinctSubseqII(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{

		{
			name: "example 1",
			s:    "abc",
			want: 7,
		},

		{
			name: "example 2",
			s:    "aba",
			want: 6,
		},

		{
			name: "example 3",
			s:    "aaa",
			want: 3,
		},

		{
			name: "example 4",
			s:    "abaa",
			want: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := distinctSubseqII(tt.s)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
