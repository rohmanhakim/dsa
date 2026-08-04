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
		{
			name:     "example 5",
			s:        "aaaaaaa",
			wordDict: []string{"aaaa", "aaa"},
			want:     true,
		},
		{
			name:     "example 6",
			s:        "aaaaaaaa",
			wordDict: []string{"aaaa", "aaa", "aa"},
			want:     true,
		},
		{
			name:     "example 7",
			s:        "ab",
			wordDict: []string{"a", "b"},
			want:     true,
		},

		{
			name:     "example 8",
			s:        "bb",
			wordDict: []string{"a", "b", "bbb", "bbbb"},
			want:     true,
		},

		{
			name:     "example 9",
			s:        "cbca",
			wordDict: []string{"bc", "ca"},
			want:     false,
		},

		{
			name:     "example 10",
			s:        "goalspecial",
			wordDict: []string{"go", "goal", "goals", "special"},
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
