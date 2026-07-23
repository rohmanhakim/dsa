package lc2487

import (
	"reflect"
	"testing"

	"github.com/rohmanhakim/leetcode/lib"
)

func TestRemoveNodes(t *testing.T) {
	tests := []struct {
		name string
		head *lib.ListNode
		want *lib.ListNode
	}{

		{
			name: "example 1",
			head: lib.ConstructListNode([]int{5, 2, 13, 3, 8}),
			want: lib.ConstructListNode([]int{13, 8}),
		},

		{
			name: "example 2",
			head: lib.ConstructListNode([]int{1, 1, 1, 1}),
			want: lib.ConstructListNode([]int{1, 1, 1, 1}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := removeNodes(tt.head)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
