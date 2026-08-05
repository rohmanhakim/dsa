package lc669

import (
	"reflect"
	"testing"

	"github.com/rohmanhakim/leetcode/lib"
)

func TestTrimBST(t *testing.T) {
	tests := []struct {
		name string
		root *lib.TreeNode
		low  int
		high int
		want *lib.TreeNode
	}{

		{
			name: "example 1",
			root: lib.SliceToBinaryTreeNodes([]*int{lib.Ptr(1), lib.Ptr(0), lib.Ptr(2)}),
			low:  1,
			high: 2,
			want: lib.SliceToBinaryTreeNodes([]*int{lib.Ptr(1), nil, lib.Ptr(2)}),
		},

		{
			name: "example 2",
			root: lib.SliceToBinaryTreeNodes([]*int{lib.Ptr(3), lib.Ptr(0), lib.Ptr(4), nil, lib.Ptr(2), nil, nil, lib.Ptr(1)}),
			low:  1,
			high: 3,
			want: lib.SliceToBinaryTreeNodes([]*int{lib.Ptr(3), lib.Ptr(2), nil, lib.Ptr(1)}),
		},

		{
			name: "example 3",
			root: lib.SliceToBinaryTreeNodes([]*int{lib.Ptr(1), nil, lib.Ptr(2)}),
			low:  2,
			high: 4,
			want: lib.SliceToBinaryTreeNodes([]*int{lib.Ptr(2)}),
		},

		{
			name: "example 4",
			root: lib.SliceToBinaryTreeNodes([]*int{lib.Ptr(3)}),
			low:  2,
			high: 2,
			want: lib.SliceToBinaryTreeNodes([]*int{nil}),
		},

		{
			name: "example 5",
			root: lib.SliceToBinaryTreeNodes([]*int{lib.Ptr(3), lib.Ptr(1), lib.Ptr(4), nil, lib.Ptr(2)}),
			low:  3,
			high: 4,
			want: lib.SliceToBinaryTreeNodes([]*int{lib.Ptr(3), nil, lib.Ptr(4)}),
		},

		{
			name: "example 6",
			root: lib.SliceToBinaryTreeNodes([]*int{lib.Ptr(3), lib.Ptr(2), lib.Ptr(4), lib.Ptr(1)}),
			low:  1,
			high: 1,
			want: lib.SliceToBinaryTreeNodes([]*int{lib.Ptr(1)}),
		},

		{
			name: "example 7",
			root: lib.SliceToBinaryTreeNodes([]*int{lib.Ptr(3), lib.Ptr(2), lib.Ptr(4), lib.Ptr(1)}),
			low:  1,
			high: 1,
			want: lib.SliceToBinaryTreeNodes([]*int{lib.Ptr(1)}),
		},

		{
			name: "example 8",
			root: lib.SliceToBinaryTreeNodes([]*int{
				lib.Ptr(2),
				lib.Ptr(0),
				lib.Ptr(33),
				nil,
				lib.Ptr(1),
				lib.Ptr(25),
				lib.Ptr(40),
				nil,
				nil,
				lib.Ptr(11),
				lib.Ptr(31),
				lib.Ptr(34),
				lib.Ptr(45),
				lib.Ptr(10),
				lib.Ptr(18),
				lib.Ptr(29),
				lib.Ptr(32),
				nil,
				lib.Ptr(36),
				lib.Ptr(43),
				lib.Ptr(46),
				lib.Ptr(4),
				nil,
				lib.Ptr(12),
				lib.Ptr(24),
				lib.Ptr(26),
				lib.Ptr(30),
				nil,
				nil,
				lib.Ptr(35),
				lib.Ptr(39),
				lib.Ptr(42),
				lib.Ptr(44),
				nil,
				lib.Ptr(48),
				lib.Ptr(3),
				lib.Ptr(9),
				nil,
				lib.Ptr(14),
				lib.Ptr(22),
				nil,
				nil,
				lib.Ptr(27),
				nil,
				nil,
				nil,
				nil,
				lib.Ptr(38),
				nil,
				lib.Ptr(41),
				nil,
				nil,
				nil,
				lib.Ptr(47),
				lib.Ptr(49),
				nil,
				nil,
				lib.Ptr(5),
				nil,
				lib.Ptr(13),
				lib.Ptr(15),
				lib.Ptr(21),
				lib.Ptr(23),
				nil,
				lib.Ptr(28),
				lib.Ptr(37),
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				lib.Ptr(8),
				nil,
				nil,
				nil,
				lib.Ptr(17),
				lib.Ptr(19),
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				lib.Ptr(7),
				nil,
				lib.Ptr(16),
				nil,
				nil,
				lib.Ptr(20),
				lib.Ptr(6),
			}),
			low:  25,
			high: 26,
			want: lib.SliceToBinaryTreeNodes([]*int{lib.Ptr(25), nil, lib.Ptr(26)}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := trimBST(tt.root, tt.low, tt.high)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
