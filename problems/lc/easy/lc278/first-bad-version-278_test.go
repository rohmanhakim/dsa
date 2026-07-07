package lc278

import (
	"reflect"
	"testing"
)

func TestFirstBadVersion(t *testing.T) {
	tests := []struct {
		name string
		n    int
		bad  int
		want int
	}{
		{
			name: "example 1",
			n:    5,
			bad:  4,
			want: 4,
		},
		{
			name: "example 2",
			n:    1,
			bad:  1,
			want: 1,
		},
		{
			name: "example 3",
			n:    5,
			bad:  2,
			want: 2,
		},
		{
			name: "example 4",
			n:    5,
			bad:  3,
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BAD = tt.bad
			out := firstBadVersion(tt.n)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
