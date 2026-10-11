package lc373

import (
	"reflect"
	"testing"
)

func TestKSmallestPairs(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		k     int
		want  [][]int
	}{

		{
			name:  "example 1",
			nums1: []int{1, 7, 11},
			nums2: []int{2, 4, 6},
			k:     3,
			want:  [][]int{{1, 2}, {1, 4}, {1, 6}},
		},

		{
			name:  "example 2",
			nums1: []int{1, 1, 2},
			nums2: []int{1, 2, 3},
			k:     2,
			want:  [][]int{{1, 1}, {1, 1}},
		},

		{
			name:  "example 3",
			nums1: []int{0, 0, 0, 0, 0},
			nums2: []int{-3, 22, 35, 56, 76},
			k:     22,
			want:  [][]int{{0, -3}, {0, -3}, {0, -3}, {0, -3}, {0, -3}, {0, 22}, {0, 22}, {0, 22}, {0, 22}, {0, 22}, {0, 35}, {0, 35}, {0, 35}, {0, 35}, {0, 35}, {0, 56}, {0, 56}, {0, 56}, {0, 56}, {0, 56}, {0, 76}, {0, 76}},
		},

		{
			name:  "example 4",
			nums1: []int{1, 2, 4, 5, 6},
			nums2: []int{3, 5, 7, 9},
			k:     20,
			want:  [][]int{{1, 3}, {2, 3}, {1, 5}, {2, 5}, {4, 3}, {1, 7}, {5, 3}, {2, 7}, {4, 5}, {6, 3}, {1, 9}, {5, 5}, {2, 9}, {4, 7}, {6, 5}, {5, 7}, {4, 9}, {6, 7}, {5, 9}, {6, 9}},
		},

		{
			name:  "example 5",
			nums1: []int{1},
			nums2: []int{3, 5, 6, 7, 8, 100},
			k:     4,
			want:  [][]int{{1, 3}, {1, 5}, {1, 6}, {1, 7}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := kSmallestPairs(tt.nums1, tt.nums2, tt.k)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
