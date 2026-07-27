package lc567

import (
	"reflect"
	"testing"
)

func TestCheckInclusion(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{

		{
			name: "example 1",
			s1:   "ab",
			s2:   "eidbaooo",
			want: true,
		},

		{
			name: "example 2",
			s1:   "ab",
			s2:   "eidboaoo",
			want: false,
		},

		{
			name: "example 3",
			s1:   "adc",
			s2:   "dcda",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := checkInclusion(tt.s1, tt.s2)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
