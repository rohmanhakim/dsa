package lc1143

import (
	"reflect"
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		name  string
		text1 string
		text2 string
		want  int
	}{
		{
			name:  "example 1",
			text1: "abcde",
			text2: "ace",
			want:  3,
		},
		{
			name:  "example 2",
			text1: "abc",
			text2: "abc",
			want:  3,
		},
		{
			name:  "example 3",
			text1: "abc",
			text2: "def",
			want:  0,
		},
		{
			name:  "example 4",
			text1: "bl",
			text2: "yby",
			want:  1,
		},
		{
			name:  "example 5",
			text1: "bsbininm",
			text2: "jmjkbkjkv",
			want:  1,
		},
		{
			name:  "example 6",
			text1: "abcba",
			text2: "abcbcba",
			want:  5,
		},
		{
			name:  "example 7",
			text1: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			text2: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			want:  210,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := longestCommonSubsequence(tt.text1, tt.text2)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
