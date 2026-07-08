package lc3

import (
	"reflect"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "example 1",
			s:    "abcabcbb",
			want: 3,
		},
		{
			name: "example 2",
			s:    "bbbbb",
			want: 1,
		},
		{
			name: "example 3",
			s:    "pwwkew",
			want: 3,
		},
		{
			name: "example 4",
			s:    "dvdf",
			want: 3,
		},
		{
			name: "example 5",
			s:    "abba",
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := lengthOfLongestSubstring(tt.s)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
