package lc2

import (
	"reflect"
	"testing"

	"github.com/rohmanhakim/leetcode/lib"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   *lib.ListNode
		l2   *lib.ListNode
		want *lib.ListNode
	}{

		{
			name: "example 1",
			l1:   lib.ConstructListNode([]int{2, 4, 3}),
			l2:   lib.ConstructListNode([]int{5, 6, 4}),
			want: lib.ConstructListNode([]int{7, 0, 8}),
		},

		{
			name: "example 2",
			l1:   lib.ConstructListNode([]int{0}),
			l2:   lib.ConstructListNode([]int{0}),
			want: lib.ConstructListNode([]int{0}),
		},

		{
			name: "example 3",
			l1:   lib.ConstructListNode([]int{9, 9, 9, 9, 9, 9, 9}),
			l2:   lib.ConstructListNode([]int{9, 9, 9, 9}),
			want: lib.ConstructListNode([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := addTwoNumbers(tt.l1, tt.l2)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
