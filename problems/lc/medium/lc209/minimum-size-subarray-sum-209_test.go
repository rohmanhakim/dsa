package lc209

import (
	"reflect"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		name   string
		target int
		nums   []int
		want   int
	}{
		{
			name:   "example 1",
			target: 7,
			nums:   []int{2, 3, 1, 2, 4, 3},
			want:   2,
		},

		{
			name:   "example 2",
			target: 4,
			nums:   []int{1, 4, 4},
			want:   1,
		},

		{
			name:   "example 3",
			target: 11,
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := minSubArrayLen(tt.target, tt.nums)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
