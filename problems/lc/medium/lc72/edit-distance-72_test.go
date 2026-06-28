package lc72

import (
	"reflect"
	"testing"
)

func TestEditDistance(t *testing.T) {
	tests := []struct {
		name  string
		word1 string
		word2 string
		want  int
	}{
		{
			name:  "example 1",
			word1: "horse",
			word2: "ros",
			want:  3,
		},
		{
			name:  "example 2",
			word1: "intention",
			word2: "execution",
			want:  5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := minDistance(tt.word1, tt.word2)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
