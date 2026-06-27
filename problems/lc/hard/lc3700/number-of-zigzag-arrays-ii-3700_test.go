package lc3700

import (
	"reflect"
	"testing"
)

func TestNumberOfZigZagArraysII(t *testing.T) {
	tests := []struct {
		name string
		n    int
		l    int
		r    int
		want int
	}{
		{
			name: "example 1",
			n:    3,
			l:    4,
			r:    5,
			want: 2,
		},
		{
			name: "example 2",
			n:    3,
			l:    1,
			r:    3,
			want: 10,
		},
		{
			name: "example 3",
			n:    425450,
			l:    37,
			r:    44,
			want: 49489607,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := zigZagArraysII(tt.n, tt.l, tt.r)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
