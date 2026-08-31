package lc2058

import (
	"reflect"
	"testing"

	"github.com/rohmanhakim/leetcode/lib"
)

func TestNodesBetweenCriticalPoints(t *testing.T) {
	tests := []struct {
		name string
		head *lib.ListNode
		want []int
	}{

		{
			name: "example 1",
			head: lib.ConstructListNode([]int{3, 1}),
			want: []int{-1, -1},
		},

		{
			name: "example 2",
			head: lib.ConstructListNode([]int{5, 3, 1, 2, 5, 1, 2}),
			want: []int{1, 3},
		},

		{
			name: "example 3",
			head: lib.ConstructListNode([]int{1, 3, 2, 2, 3, 2, 2, 2, 7}),
			want: []int{3, 3},
		},

		{
			name: "example 4",
			head: lib.ConstructListNode([]int{2, 2, 1, 3}),
			want: []int{-1, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := nodesBetweenCriticalPoints(tt.head)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
