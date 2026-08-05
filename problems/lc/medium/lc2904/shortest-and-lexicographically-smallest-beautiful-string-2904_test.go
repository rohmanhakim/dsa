package lc2904

import (
	"reflect"
	"testing"
)

func TestShortestBeautifulSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want string
	}{

		{
			name: "example 1",
			s:    "100011001",
			k:    3,
			want: "11001",
		},

		{
			name: "example 2",
			s:    "1011",
			k:    2,
			want: "11",
		},

		{
			name: "example 3",
			s:    "000",
			k:    1,
			want: "",
		},

		{
			name: "example 4",
			s:    "001110101101101111",
			k:    10,
			want: "10101101101111",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := shortestBeautifulSubstring(tt.s, tt.k)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
