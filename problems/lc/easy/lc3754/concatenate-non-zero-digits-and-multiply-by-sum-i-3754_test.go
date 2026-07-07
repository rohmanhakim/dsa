package lc3754

import (
	"reflect"
	"testing"
)

func TestSumAndMultiply(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int64
	}{
		{
			name: "example 1",
			n:    10203004,
			want: 12340,
		},
		{
			name: "example 2",
			n:    1000,
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := sumAndMultiply(tt.n)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
