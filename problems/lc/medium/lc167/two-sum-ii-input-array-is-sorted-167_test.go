package lc167

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{

		{
			name:    "example 1",
			numbers: []int{2, 7, 11, 15},
			target:  9,
			want:    []int{1, 2},
		},

		{
			name:    "example 2",
			numbers: []int{2, 3, 4},
			target:  6,
			want:    []int{1, 3},
		},

		{
			name:    "example 3",
			numbers: []int{-1, 0},
			target:  -1,
			want:    []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := twoSum(tt.numbers, tt.target)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
