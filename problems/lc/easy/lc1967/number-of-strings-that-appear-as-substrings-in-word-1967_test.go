package lc1967

import (
	"reflect"
	"testing"
)

func TestNumberofStringsThatAppearasSubstringsinWord(t *testing.T) {
	tests := []struct {
		name     string
		patterns []string
		word     string
		want     int
	}{
		{
			name:     "example 1",
			patterns: []string{"a", "abc", "bc", "d"},
			word:     "abc",
			want:     3,
		},
		{
			name:     "example 2",
			patterns: []string{"a", "b", "c"},
			word:     "aaaaabbbbb",
			want:     2,
		},
		{
			name:     "example 3",
			patterns: []string{"a", "a", "a"},
			word:     "ab",
			want:     3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := numOfStrings(tt.patterns, tt.word)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
