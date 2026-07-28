package lc518

import (
	"reflect"
	"testing"
)

func TestChange(t *testing.T) {
	tests := []struct {
		name   string
		amount int
		coins  []int
		want   int
	}{

		{
			name:   "example 1",
			amount: 5,
			coins:  []int{1, 2, 5},
			want:   4,
		},

		{
			name:   "example 2",
			amount: 3,
			coins:  []int{2},
			want:   0,
		},

		{
			name:   "example 3",
			amount: 10,
			coins:  []int{10},
			want:   1,
		},

		{
			name:   "example 4",
			amount: 5,
			coins:  []int{2, 3, 5},
			want:   2,
		},

		{
			name:   "example 5",
			amount: 10,
			coins:  []int{5},
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := change(tt.amount, tt.coins)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
