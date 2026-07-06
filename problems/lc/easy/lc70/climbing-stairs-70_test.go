package lc70

import (
	"reflect"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "example 1",
			n:    2,
			want: 2,
		},
		{
			name: "example 2",
			n:    3,
			want: 3,
		},
		{
			name: "example 3",
			n:    4,
			want: 5,
		},
		{
			name: "example 4",
			n:    1,
			want: 1,
		},
		{
			name: "example 5",
			n:    5,
			want: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := climbStairs(tt.n)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
