package lc3498

import (
	"reflect"
	"testing"
)

func TestReverseDegree(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{

		{
			name: "example 1",
			s:    "abc",
			want: 148,
		},

		{
			name: "example 2",
			s:    "zaza",
			want: 160,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := reverseDegree(tt.s)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
