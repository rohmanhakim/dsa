package lc21

import (
	"reflect"
	"testing"

	"github.com/rohmanhakim/leetcode/lib"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name  string
		list1 *lib.ListNode
		list2 *lib.ListNode
		want  *lib.ListNode
	}{

		{
			name:  "example 1",
			list1: lib.ConstructListNode([]int{1, 2, 4}),
			list2: lib.ConstructListNode([]int{1, 3, 4}),
			want:  lib.ConstructListNode([]int{1, 1, 2, 3, 4, 4}),
		},

		{
			name:  "example 2",
			list1: lib.ConstructListNode([]int{}),
			list2: lib.ConstructListNode([]int{}),
			want:  nil,
		},

		{
			name:  "example 3",
			list1: lib.ConstructListNode([]int{}),
			list2: lib.ConstructListNode([]int{0}),
			want:  lib.ConstructListNode([]int{0}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := mergeTwoLists(tt.list1, tt.list2)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
