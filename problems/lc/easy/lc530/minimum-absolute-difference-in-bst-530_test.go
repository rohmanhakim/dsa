package lc530

import (
	"reflect"
	"testing"

	"github.com/rohmanhakim/leetcode/lib"
)

func TestGetMinimumDifference(t *testing.T) {
	tests := []struct {
		name string
		root *lib.TreeNode
		want int
	}{
		{
			name: "example 1",
			root: lib.SliceToBinaryTreeNodes([]*int{lib.Ptr(4), lib.Ptr(2), lib.Ptr(6), lib.Ptr(1), lib.Ptr(3)}),
			want: 1,
		},
		{
			name: "example 2",
			root: lib.SliceToBinaryTreeNodes([]*int{lib.Ptr(1), lib.Ptr(0), lib.Ptr(48), nil, nil, lib.Ptr(12), lib.Ptr(49)}),
			want: 1,
		},
		{
			name: "example 3",
			root: lib.SliceToBinaryTreeNodes([]*int{lib.Ptr(236), lib.Ptr(104), lib.Ptr(701), nil, lib.Ptr(227), nil, lib.Ptr(911)}),
			want: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := getMinimumDifference(tt.root)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
