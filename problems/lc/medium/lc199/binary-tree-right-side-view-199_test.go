package lc199

import (
	"reflect"
	"testing"

	"github.com/rohmanhakim/leetcode/lib"
)

func TestRightSideView(t *testing.T) {
	tests := []struct {
		name string
		root *lib.TreeNode
		want []int
	}{

		{
			name: "example 1",
			root: lib.SliceToBinaryTreeNodes([]*int{lib.Ptr(1), lib.Ptr(2), lib.Ptr(3), nil, lib.Ptr(5), nil, lib.Ptr(4)}),
			want: []int{1, 3, 4},
		},

		{
			name: "example 2",
			root: lib.SliceToBinaryTreeNodes([]*int{lib.Ptr(1), lib.Ptr(2), lib.Ptr(3), lib.Ptr(4), nil, nil, nil, lib.Ptr(5)}),
			want: []int{1, 3, 4, 5},
		},

		{
			name: "example 3",
			root: lib.SliceToBinaryTreeNodes([]*int{lib.Ptr(1), nil, lib.Ptr(3)}),
			want: []int{1, 3},
		},
		{
			name: "example 4",
			root: lib.SliceToBinaryTreeNodes([]*int{}),
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := rightSideView(tt.root)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
