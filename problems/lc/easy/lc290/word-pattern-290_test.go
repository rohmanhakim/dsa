package lc290

import (
	"reflect"
	"testing"
)

func TestWordPattern(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		s       string
		want    bool
	}{

		{
			name:    "example 1",
			pattern: "abba",
			s:       "dog cat cat dog",
			want:    true,
		},

		{
			name:    "example 2",
			pattern: "abba",
			s:       "dog cat cat fish",
			want:    false,
		},

		{
			name:    "example 3",
			pattern: "aaaa",
			s:       "dog cat cat dog",
			want:    false,
		},

		{
			name:    "example 4",
			pattern: "abba",
			s:       "dog dog dog dog",
			want:    false,
		},

		{
			name:    "example 5",
			pattern: "aaa",
			s:       "aa aa aa aa",
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := wordPattern(tt.pattern, tt.s)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
