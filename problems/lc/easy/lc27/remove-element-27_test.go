package lc27

import (
	"reflect"
	"slices"
	"sort"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		val      int
		want     int
		wantNums []int
	}{

		{
			name:     "example 1",
			nums:     []int{3, 2, 2, 3},
			val:      3,
			want:     2,
			wantNums: []int{2, 2},
		},

		{
			name:     "example 2",
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			want:     5,
			wantNums: []int{0, 1, 4, 0, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := removeElement(tt.nums, tt.val)

			if !reflect.DeepEqual(k, tt.want) {
				t.Fatalf("got %v want %v", k, tt.want)
			}
			firstK := tt.nums[:k]
			sort.Ints(firstK)
			sort.Ints(tt.wantNums)

			if slices.Compare(firstK, tt.wantNums) != 0 {
				t.Fatalf("got %v want %v", firstK, tt.wantNums)
			}
		})
	}
}
