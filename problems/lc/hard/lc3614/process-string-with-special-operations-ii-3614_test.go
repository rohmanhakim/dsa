package lc3614

import (
	"reflect"
	"testing"
)

func TestProcessStringWithSpecialOperationsII(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want byte
	}{
		{
			name: "test 1",
			s:    "a#b%*",
			k:    1,
			want: byte('a'),
		},
		{
			name: "test 2",
			s:    "cd%#*#",
			k:    3,
			want: byte('d'),
		},
		{
			name: "test 3",
			s:    "z*#",
			k:    0,
			want: byte('.'),
		},
		{
			name: "test 4",
			s:    "*mi#",
			k:    442611269896338,
			want: byte('.'),
		},
		{
			name: "test 5",
			s:    "%%pd#jw#%c*vi##nm%e#a#oz####k##rlow%*w%f#u%%h#f*#f#g##%#cnq#",
			k:    3580,
			want: byte('w'),
		},
		{
			name: "test 6",
			s:    "%#*gm#xib",
			k:    2,
			want: byte('g'),
		},
	}

	for _, tt := range tests {
		out := processStr(tt.s, int64(tt.k))
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
