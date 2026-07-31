package lc3016

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
			word: "xyzxyzxyzxyz",
			want: 12,
		},

		{
			name: "example 3",
			word: "aremgjnptohswfdkyuzvicqxb",
			want: 52,
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
