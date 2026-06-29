package lc14

import (
	"reflect"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
		{
			name: "example 1",
			strs: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			name: "example 2",
			strs: []string{"dog", "racecar", "car"},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := longestCommonPrefix(tt.strs)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
