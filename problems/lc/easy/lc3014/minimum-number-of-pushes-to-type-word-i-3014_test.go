package lc3014

import (
	"reflect"
	"testing"
)

func TestMinimumPushes(t *testing.T) {
	tests := []struct {
		name string
		word string
		want int
	}{

		{
			name: "example 1",
			word: "abcde",
			want: 5,
		},

		{
			name: "example 2",
			word: "xycdefghij",
			want: 12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := minimumPushes(tt.word)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
