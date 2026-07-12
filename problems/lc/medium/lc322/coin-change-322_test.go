package lc322

import (
	"reflect"
	"testing"
)

func TestCoinChange(t *testing.T) {
	tests := []struct {
		name   string
		coins  []int
		amount int
		want   int
	}{

		{
			name:   "example 1",
			coins:  []int{1, 2, 5},
			amount: 11,
			want:   3,
		},

		{
			name:   "example 2",
			coins:  []int{2},
			amount: 3,
			want:   -1,
		},

		{
			name:   "example 3",
			coins:  []int{1},
			amount: 0,
			want:   0,
		},
		{
			name:   "example 4",
			coins:  []int{2, 5, 10, 1},
			amount: 27,
			want:   4,
		},
		{
			name:   "example 5",
			coins:  []int{186, 419, 83, 408},
			amount: 6249,
			want:   20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := coinChange(tt.coins, tt.amount)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
