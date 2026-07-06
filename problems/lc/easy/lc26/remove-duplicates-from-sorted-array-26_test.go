package lc26

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name       string
		nums       []int
		want       int
		wantedNums []int
	}{
		{
			name:       "example 1",
			nums:       []int{1, 1, 2},
			want:       2,
			wantedNums: []int{1, 2},
		},
		{
			name:       "example 2",
			nums:       []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want:       5,
			wantedNums: []int{0, 1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := removeDuplicates(tt.nums)

			if !reflect.DeepEqual(k, len(tt.wantedNums)) {
				t.Fatalf("unmatching length. got %v want %v", k, len(tt.wantedNums))
			}
			for i := 0; i < k; i++ {
				if !reflect.DeepEqual(tt.nums[i], tt.wantedNums[i]) {
					t.Fatalf("unmatching item. got %v want %v", tt.nums[i], tt.wantedNums[i])
				}
			}
		})
	}
}
