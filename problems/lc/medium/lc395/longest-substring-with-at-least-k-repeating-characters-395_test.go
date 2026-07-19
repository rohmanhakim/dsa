package lc395

import (
	"reflect"
	"testing"
)

func TestLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{

		{
			name: "example 1",
			s:    "aaabb",
			k:    3,
			want: 3,
		},

		{
			name: "example 2",
			s:    "ababbc",
			k:    2,
			want: 5,
		},

		{
			name: "example 3",
			s:    "aaabbcccdddd",
			k:    3,
			want: 7,
		},

		{
			name: "example 4",
			s:    "ababbcddddeee",
			k:    2,
			want: 7,
		},
		{
			name: "example 5",
			s:    "ababbcddddeeec",
			k:    2,
			want: 14,
		},
		{
			name: "example 6",
			s:    "ababbcddddae",
			k:    2,
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := longestSubstring(tt.s, tt.k)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
