package lc714

import (
	"reflect"
	"testing"
)

func TestBestTimetoBuyandSellStockWithTransactionFee(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		fee    int
		want   int
	}{
		{
			name:   "example 1",
			prices: []int{1, 3, 2, 8, 4, 9},
			fee:    2,
			want:   8,
		},
		{
			name:   "example 2",
			prices: []int{1, 3, 7, 5, 10, 3},
			fee:    3,
			want:   6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := maxProfit(tt.prices, tt.fee)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
