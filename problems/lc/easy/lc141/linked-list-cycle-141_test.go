package lc141

import (
	"reflect"
	"testing"

	"github.com/rohmanhakim/leetcode/lib"
)

func TestHasCycle(t *testing.T) {
	tests := []struct {
		name string
		head *lib.ListNode
		pos  int
		want bool
	}{
		{
			name: "example 1",
			head: lib.ConstructListNode([]int{3, 2, 0, -4}),
			pos:  1,
			want: true,
		},
		{
			name: "example 2",
			head: lib.ConstructListNode([]int{1, 2}),
			pos:  0,
			want: true,
		},
		{
			name: "example 3",
			head: lib.ConstructListNode([]int{1}),
			pos:  -1,
			want: false,
		},
		{
			name: "example 4",
			head: lib.ConstructListNode([]int{-21, 10, 17, 8, 4, 26, 5, 35, 33, -7, -16, 27, -12, 6, 29, -12, 5, 9, 20, 14, 14, 2, 13, -24, 21, 23, -21, 5}),
			pos:  -1,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lib.LinkTail(tt.head, tt.pos)
			out := hasCycle(tt.head)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
