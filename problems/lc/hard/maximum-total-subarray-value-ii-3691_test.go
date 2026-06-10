package hard

import (
	"reflect"
	"testing"
)

func TestMaximumTotalSubarrayValueII(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int64
	}{
		{
			name: "test1",
			nums: []int{1, 3, 2},
			k:    2,
			want: 4,
		},
		{
			name: "test2",
			nums: []int{4, 2, 5, 1},
			k:    3,
			want: 12,
		},
		{
			name: "test3",
			nums: []int{11, 8},
			k:    2,
			want: 3,
		},
		{
			name: "test4",
			nums: []int{9, 9, 37},
			k:    1,
			want: 28,
		},
		{
			name: "test5",
			nums: []int{18, 36, 6},
			k:    3,
			want: 78,
		},
		{
			name: "test6",
			nums: []int{9, 9, 37},
			k:    3,
			want: 56,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxTotalValue(tt.nums, tt.k)

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}
