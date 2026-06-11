package hard

import (
	"reflect"
	"testing"
)

func TestConstructTargetArrayWithMultipleSums(t *testing.T) {
	tests := []struct {
		name   string
		target []int
		want   bool
	}{
		{
			// 9 3 5   -->  18 - 17 = 1
			// 1 3 5   -->  10 - 9 = 1
			// 1 3 1   -->  6 - 5 = 1
			// 1 1 1   --> 2 - 3
			name:   "test 1",
			target: []int{9, 3, 5},
			want:   true,
		},
		{
			name:   "test 2",
			target: []int{1, 1, 1, 2},
			want:   false,
		},
		{
			name:   "test 3",
			target: []int{8, 5},
			want:   true,
		},
		{
			name:   "test 4",
			target: []int{5, 2},
			want:   true,
		},
		{
			name:   "test 5",
			target: []int{1, 1000000000},
			want:   true,
		},
		{
			name:   "test 6",
			target: []int{1, 1, 2},
			want:   false,
		},
		{
			name:   "test 7",
			target: []int{2},
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPossible(tt.target)

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}
