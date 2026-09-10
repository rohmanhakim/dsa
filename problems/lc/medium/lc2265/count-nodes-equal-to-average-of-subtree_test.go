package lc2265

import (
	"reflect"
	"testing"

	"github.com/rohmanhakim/leetcode/lib"
)

func TestAverageOfSubtree(t *testing.T) {
	tests := []struct {
		name string
		root *lib.TreeNode
		want int
	}{

		{
			name: "example 1",
			root: lib.SliceToBinaryTreeNodes([]*int{lib.Ptr(4), lib.Ptr(8), lib.Ptr(5), lib.Ptr(0), lib.Ptr(1), nil, lib.Ptr(6)}),
			want: 5,
		},

		{
			name: "example 2",
			root: lib.SliceToBinaryTreeNodes([]*int{lib.Ptr(1)}),
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := averageOfSubtree(tt.root)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
