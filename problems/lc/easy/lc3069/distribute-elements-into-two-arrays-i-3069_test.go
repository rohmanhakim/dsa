package lc3069

import (
	"reflect"
	"testing"
)

func TestResultArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{

		{
			name: "example 1",
			nums: []int{2, 1, 3},
			want: []int{2, 3, 1},
		},

		{
			name: "example 2",
			nums: []int{5, 4, 3, 8},
			want: []int{5, 3, 4, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := resultArray(tt.nums)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
