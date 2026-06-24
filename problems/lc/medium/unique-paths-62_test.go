package medium

import (
	"reflect"
	"testing"
)

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name string
		m    int
		n    int
		want int
	}{
		{
			name: "example 1",
			m:    3,
			n:    7,
			want: 28,
		},
		{
			name: "example 2",
			m:    3,
			n:    2,
			want: 3,
		},
		{
			name: "example 3",
			m:    19,
			n:    13,
			want: 86493225,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := uniquePathsBottomUp(tt.m, tt.n)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
