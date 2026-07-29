package lc3518

import (
	"reflect"
	"testing"
)

func TestSmallestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want string
	}{

		{
			name: "example 1",
			s:    "abba",
			k:    2,
			want: "baab",
		},

		{
			name: "example 2",
			s:    "aa",
			k:    2,
			want: "",
		},

		{
			name: "example 3",
			s:    "bacab",
			k:    1,
			want: "abcba",
		},

		{
			name: "example 4",
			s:    "bcaacb",
			k:    6,
			want: "cbaabc",
		},

		{
			name: "example 5",
			s:    "hcch",
			k:    2,
			want: "hcch",
		},

		{
			name: "example 6",
			s:    "xxnfnxx",
			k:    3,
			want: "xxnfnxx",
		},

		{
			name: "example 7",
			s:    "qttq",
			k:    2,
			want: "tqqt",
		},

		{
			name: "example 8",
			s:    "ycnypoidimbrvasctjkkwwkkjtcsavrbmidiopyncy",
			k:    32993,
			want: "abccdiijkkmnpwovytyrssrytyvowpnmkkjiidccba",
		},

		{
			name: "example 9",
			s:    "aaaaaaaabbbbbbbbccccccccdddddddccccccccbbbbbbbbaaaaaaaa",
			k:    750000,
			want: "aaaaaaaabbcbbbdcbcdcbcccdcbdbcdcccbcdcbcdbbbcbbaaaaaaaa",
		},

		{
			name: "example 10",
			s:    "kxk",
			k:    2,
			want: "",
		},

		{
			name: "example 11",
			s:    "o",
			k:    1,
			want: "o",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := smallestPalindrome(tt.s, tt.k)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
