package lc438

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		name string
		s    string
		p    string
		want []int
	}{

		{
			name: "example 1",
			s:    "cbaebabacd",
			p:    "abc",
			want: []int{0, 6},
		},

		{
			name: "example 2",
			s:    "abab",
			p:    "ab",
			want: []int{0, 1, 2},
		},
		{
			name: "example 3",
			s:    "baa",
			p:    "aa",
			want: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := findAnagrams(tt.s, tt.p)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
