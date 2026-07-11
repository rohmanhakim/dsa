package lc104

import (
	"reflect"
	"testing"

	"github.com/rohmanhakim/leetcode/lib"
)

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name string
		root *lib.TreeNode
		want int
	}{
		{
			name: "example 1",
			root: lib.SliceToBinaryTreeNodes([]*int{lib.Ptr(3), lib.Ptr(9), lib.Ptr(20), nil, nil, lib.Ptr(15), lib.Ptr(7)}),
			want: 3,
		},
		{
			name: "example 2",
			root: lib.SliceToBinaryTreeNodes([]*int{lib.Ptr(1), nil, lib.Ptr(2)}),
			want: 2,
		},
		{
			name: "example 3",
			root: lib.SliceToBinaryTreeNodes([]*int{lib.Ptr(1), lib.Ptr(2)}),
			want: 2,
		},
		{
			name: "example 4",
			root: lib.SliceToBinaryTreeNodes([]*int{lib.Ptr(1), lib.Ptr(2), lib.Ptr(3), lib.Ptr(4), lib.Ptr(5)}),
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := maxDepth(tt.root)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
