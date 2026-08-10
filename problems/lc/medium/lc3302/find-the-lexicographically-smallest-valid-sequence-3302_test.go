package lc3302

import (
	"reflect"
	"testing"
)

func TestValidSequence(t *testing.T) {
	tests := []struct {
		name  string
		word1 string
		word2 string
		want  []int
	}{

		{
			name:  "example 0",
			word1: "plane",
			word2: "plank",
			want:  []int{0, 1, 2, 3, 4},
		},
		{
			name:  "example 1",
			word1: "vbcca",
			word2: "abc",
			want:  []int{0, 1, 2},
		},

		{
			name:  "example 2",
			word1: "bacdc",
			word2: "abc",
			want:  []int{1, 2, 4},
		},

		{
			name:  "example 3",
			word1: "aaaaaa",
			word2: "aaabc",
			want:  []int{},
		},

		{
			name:  "example 4",
			word1: "abc",
			word2: "ab",
			want:  []int{0, 1},
		},

		{
			name:  "example 5",
			word1: "ccbccccbcc",
			word2: "b",
			want:  []int{0},
		},

		{
			name:  "example 6",
			word1: "bbeigiibhjafjig",
			word2: "iihhj",
			want:  []int{3, 5, 6, 8, 9},
		},
		{
			name:  "example 7",
			word1: "cbbccc",
			word2: "bb",
			want:  []int{0, 1},
		},
		{
			name:  "example 8",
			word1: "cbcbccbbcbcbbbb",
			word2: "bb",
			want:  []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := validSequence(tt.word1, tt.word2)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
