package lc383

import (
	"reflect"
	"testing"
)

func TestCanConstruct(t *testing.T) {
	tests := []struct {
		name       string
		ransomNote string
		magazine   string
		want       bool
	}{
		{
			name:       "example 1",
			ransomNote: "a",
			magazine:   "b",
			want:       false,
		},
		{
			name:       "example 2",
			ransomNote: "aa",
			magazine:   "ab",
			want:       false,
		},
		{
			name:       "example 3",
			ransomNote: "aa",
			magazine:   "aab",
			want:       true,
		},
		{
			name:       "example 4",
			ransomNote: "ac",
			magazine:   "aab",
			want:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := canConstruct(tt.ransomNote, tt.magazine)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
