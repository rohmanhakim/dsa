package lc1927

import (
	"reflect"
	"testing"
)

func TestSumGame(t *testing.T) {
	tests := []struct {
		name string
		num  string
		want bool
	}{

		{
			name: "example 1",
			num:  "5023",
			want: false,
		},

		{
			name: "example 2",
			num:  "25??",
			want: true,
		},

		{
			name: "example 3",
			num:  "?3295???",
			want: false,
		},

		{
			name: "example 4",
			num:  "?3?32?",
			want: true,
		},

		{
			name: "example 5",
			num:  "7??12?",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := sumGame(tt.num)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
