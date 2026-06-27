package lc1833

import (
	"reflect"
	"testing"
)

func TestMaximumIceCreamBar(t *testing.T) {
	tests := []struct {
		name  string
		costs []int
		coins int
		want  int
	}{
		{
			name:  "test 1",
			costs: []int{1, 3, 2, 4, 1},
			coins: 7,
			want:  4,
		},
		{
			name:  "test 2",
			costs: []int{10, 6, 8, 7, 7, 8},
			coins: 5,
			want:  0,
		},
		{
			name:  "test 3",
			costs: []int{1, 6, 3, 1, 2, 5},
			coins: 20,
			want:  6,
		},
		{
			name:  "test 4",
			costs: []int{11},
			coins: 12,
			want:  1,
		},
		{
			name:  "test 5",
			costs: []int{4, 7, 6, 4, 4, 2, 2, 4, 8, 8},
			coins: 41,
			want:  9,
		},
		{
			name:  "test 6",
			costs: []int{6, 2, 8, 8, 5, 6, 6, 2, 2, 2},
			coins: 77,
			want:  10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := maxIceCream(tt.costs, tt.coins)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
