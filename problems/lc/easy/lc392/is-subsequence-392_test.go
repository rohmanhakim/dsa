package lc392

import (
	"reflect"
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{

		{
			name: "example 1",
			s:    "abc",
			t:    "ahbgdc",
			want: true,
		},

		{
			name: "example 2",
			s:    "axc",
			t:    "ahbgdc",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := isSubsequence(tt.s, tt.t)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}

func TestIsSubsequenceFollowUp(t *testing.T) {
	tests := []struct {
		name   string
		stream []string
		t      string
		want   []bool
	}{

		{
			name:   "example 1",
			stream: []string{"abc", "axc"},
			t:      "ahbgdc",
			want:   []bool{true, false},
		},

		{
			name:   "example 2",
			stream: []string{"axc", "abc", "agd"},
			t:      "ahbgdc",
			want:   []bool{false, true, true},
		},
		{
			name:   "example 3",
			stream: []string{"axc", "abc", "aaa"},
			t:      "abacada",
			want:   []bool{false, true, true},
		},
		{
			name:   "example 4",
			stream: []string{"ba", "aa"},
			t:      "aab",
			want:   []bool{false, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := isSubsequenceFollowUp(tt.stream, tt.t)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
