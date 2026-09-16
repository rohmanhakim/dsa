package lc1621

import (
	"reflect"
	"testing"
)

func TestNumberOfSets(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want int
	}{

		{
			name: "example 1",
			n:    4,
			k:    2,
			want: 5,
		},

		{
			name: "example 2",
			n:    3,
			k:    1,
			want: 3,
		},

		{
			name: "example 3",
			n:    30,
			k:    7,
			want: 796297179,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := numberOfSetsONSpace(tt.n, tt.k)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
