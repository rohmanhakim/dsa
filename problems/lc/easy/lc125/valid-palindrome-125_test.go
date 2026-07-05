package lc125

import (
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "example 1",
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: "example 2",
			s:    "Race a car",
			want: false,
		},
		{
			name: "example 3",
			s:    " ",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := isPalindrome(tt.s)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
