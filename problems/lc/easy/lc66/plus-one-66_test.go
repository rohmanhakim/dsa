package lc66

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name   string
		digits []int
		want   []int
	}{

		{
			name:   "example 1",
			digits: []int{1, 2, 3},
			want:   []int{1, 2, 4},
		},
		{
			name:   "example 2",
			digits: []int{4, 3, 2, 1},
			want:   []int{4, 3, 2, 2},
		},
		{
			name:   "example 3",
			digits: []int{9},
			want:   []int{1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := plusOne(tt.digits)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
