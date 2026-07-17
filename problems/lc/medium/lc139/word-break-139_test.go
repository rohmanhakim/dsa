package lc139

import (
	"reflect"
	"testing"
)

func TestWordBreak(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		want     bool
	}{

		{
			name:     "example 1",
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			want:     true,
		},

		{
			name:     "example 2",
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
			want:     true,
		},

		{
			name:     "example 3",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			want:     false,
		},
		{
			name:     "example 4",
			s:        "cars",
			wordDict: []string{"car", "ca", "rs"},
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := wordBreak(tt.s, tt.wordDict)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
