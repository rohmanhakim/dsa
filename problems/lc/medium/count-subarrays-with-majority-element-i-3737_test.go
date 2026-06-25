package medium

import (
	"reflect"
	"testing"
)

func TestCountSubarraysWithMajorityElementI(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "example 1",
			nums:   []int{1, 2, 2, 3},
			target: 2,
			want:   5,
		},
		{
			name:   "example 2",
			nums:   []int{1, 1, 1, 1},
			target: 1,
			want:   10,
		},
		{
			name:   "example 3",
			nums:   []int{1, 2, 3},
			target: 4,
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := countMajoritySubarrays(tt.nums, tt.target)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
