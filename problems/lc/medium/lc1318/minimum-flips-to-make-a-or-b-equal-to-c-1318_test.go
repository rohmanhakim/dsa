package lc1318

import (
	"reflect"
	"testing"
)

func TestMinFlips(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		c    int
		want int
	}{
		{
			name: "example 1",
			a:    2,
			b:    6,
			c:    5,
			want: 3,
		},
		{
			name: "example 2",
			a:    4,
			b:    2,
			c:    7,
			want: 1,
		},
		{
			name: "example 3",
			a:    1,
			b:    2,
			c:    3,
			want: 0,
		},
		{
			name: "example 4",
			a:    8,
			b:    3,
			c:    5,
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := minFlips(tt.a, tt.b, tt.c)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
