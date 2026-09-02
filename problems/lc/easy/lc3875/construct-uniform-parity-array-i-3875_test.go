package lc3875

import (
	"reflect"
	"testing"
)

func TestUniformArray(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		want  bool
	}{

		{
			name:  "example 1",
			nums1: []int{2, 3},
			want:  true,
		},

		{
			name:  "example 2",
			nums1: []int{4, 6},
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := uniformArray(tt.nums1)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
