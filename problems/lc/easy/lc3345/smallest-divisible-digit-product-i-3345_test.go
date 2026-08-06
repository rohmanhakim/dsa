package lc3345

import (
	"reflect"
	"testing"
)

func TestSmallestNumber(t *testing.T) {
	tests := []struct {
		name string
		n    int
		t    int
		want int
	}{

		{
			name: "example 1",
			n:    10,
			t:    2,
			want: 10,
		},

		{
			name: "example 2",
			n:    15,
			t:    3,
			want: 16,
		},

		{
			name: "example 3",
			n:    1,
			t:    10,
			want: 10,
		},

		{
			name: "example 4",
			n:    11,
			t:    10,
			want: 20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := smallestNumber(tt.n, tt.t)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
