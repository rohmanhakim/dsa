package lc115

import (
	"reflect"
	"testing"
)

func TestNumDistinct(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want int
	}{

		{
			name: "example 1",
			s:    "rabbbit",
			t:    "rabbit",
			want: 3,
		},

		{
			name: "example 2",
			s:    "babgbag",
			t:    "bag",
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := numDistinct(tt.s, tt.t)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
