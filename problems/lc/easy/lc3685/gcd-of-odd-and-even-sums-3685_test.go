package lc3685

import (
	"reflect"
	"testing"
)

func TestGcdOfOddEvenSums(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "example 1",
			n:    4,
			want: 4,
		},
		{
			name: "example 2",
			n:    5,
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := gcdOfOddEvenSums(tt.n)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
