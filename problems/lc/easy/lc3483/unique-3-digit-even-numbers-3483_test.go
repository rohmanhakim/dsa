package lc3483

import (
	"reflect"
	"testing"
)

func TestTotalNumbers(t *testing.T) {
	tests := []struct {
		name   string
		digits []int
		want   int
	}{

		{
			name:   "example 1",
			digits: []int{1, 2, 3, 4},
			want:   12,
		},

		{
			name:   "example 2",
			digits: []int{0, 2, 2},
			want:   2,
		},

		{
			name:   "example 3",
			digits: []int{6, 6, 6},
			want:   1,
		},

		{
			name:   "example 4",
			digits: []int{1, 3, 5},
			want:   0,
		},

		{
			name:   "example 5",
			digits: []int{1, 2, 2, 2},
			want:   3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := totalNumbers(tt.digits)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
