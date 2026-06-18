package medium

import (
	"reflect"
	"testing"
)

func TestAngleBetweenHandsOfAClock(t *testing.T) {
	tests := []struct {
		name    string
		hour    int
		minutes int
		want    float64
	}{
		{
			name:    "tests 1",
			hour:    12,
			minutes: 30,
			want:    165,
		},
		{
			name:    "tests 2",
			hour:    3,
			minutes: 30,
			want:    75,
		},
		{
			name:    "tests 2",
			hour:    3,
			minutes: 15,
			want:    7.5,
		},
		{
			name:    "tests 3",
			hour:    1,
			minutes: 57,
			want:    76.5,
		},
		{
			name:    "tests 4",
			hour:    4,
			minutes: 50,
			want:    155,
		},
		{
			name:    "tests 5",
			hour:    12,
			minutes: 0,
			want:    0,
		},
		{
			name:    "tests 6",
			hour:    8,
			minutes: 52,
			want:    46,
		},
		{
			name:    "tests 7",
			hour:    9,
			minutes: 41,
			want:    44.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := angleClock(tt.hour, tt.minutes)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
