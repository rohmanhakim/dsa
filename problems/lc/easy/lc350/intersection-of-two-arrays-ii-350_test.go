package lc350

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{

		{
			name:  "example 1",
			nums1: []int{1, 2, 2, 1},
			nums2: []int{2, 2},
			want:  []int{2, 2},
		},

		{
			name:  "example 2",
			nums1: []int{4, 9, 5},
			nums2: []int{9, 4, 9, 8, 4},
			want:  []int{4, 9},
		},

		{
			name:  "example 3",
			nums1: []int{1, 2, 2, 1},
			nums2: []int{2},
			want:  []int{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := intersect(tt.nums1, tt.nums2)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}

func TestIntersectAlternative(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{

		{
			name:  "example 1",
			nums1: []int{1, 2, 2, 1},
			nums2: []int{2, 2},
			want:  []int{2, 2},
		},

		{
			name:  "example 2",
			nums1: []int{4, 9, 5},
			nums2: []int{9, 4, 9, 8, 4},
			want:  []int{4, 9},
		},

		{
			name:  "example 3",
			nums1: []int{1, 2, 2, 1},
			nums2: []int{2},
			want:  []int{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := intersectAlternative(tt.nums1, tt.nums2)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
