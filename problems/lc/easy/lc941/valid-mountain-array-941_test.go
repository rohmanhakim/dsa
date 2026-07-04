package lc941

import (
	"reflect"
	"testing"
)

func TestValidMountainArray(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		{
			name: "example 1",
			arr:  []int{2, 1},
			want: false,
		},
		{
			name: "example 2",
			arr:  []int{3, 5, 5},
			want: false,
		},
		{
			name: "example 3",
			arr:  []int{0, 3, 2, 1},
			want: true,
		},
		{
			name: "example 4",
			arr:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: false,
		},
		{
			name: "example 5",
			arr:  []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			want: false,
		},
		{
			name: "example 6",
			arr:  []int{6, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			want: true,
		},
		{
			name: "example 6",
			arr:  []int{6, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := validMountainArray(tt.arr)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
