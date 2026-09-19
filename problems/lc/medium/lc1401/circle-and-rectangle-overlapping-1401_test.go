package lc1401

import (
	"reflect"
	"testing"
)

func TestCheckOverlap(t *testing.T) {
	tests := []struct {
		name    string
		radius  int
		xCenter int
		yCenter int
		x1      int
		y1      int
		x2      int
		y2      int
		want    bool
	}{

		{
			name:    "example 1",
			radius:  1,
			xCenter: 0,
			yCenter: 0,
			x1:      1,
			y1:      -1,
			x2:      3,
			y2:      1,
			want:    true,
		},

		{
			name:    "example 2",
			radius:  1,
			xCenter: 1,
			yCenter: 1,
			x1:      1,
			y1:      -3,
			x2:      2,
			y2:      -1,
			want:    false,
		},

		{
			name:    "example 3",
			radius:  1,
			xCenter: 0,
			yCenter: 0,
			x1:      -1,
			y1:      0,
			x2:      0,
			y2:      1,
			want:    true,
		},

		{
			name:    "example 4",
			radius:  1415,
			xCenter: 807,
			yCenter: -704,
			x1:      -733,
			y1:      623,
			x2:      -533,
			y2:      1005,
			want:    false,
		},

		{
			name:    "example 5",
			radius:  2,
			xCenter: 102,
			yCenter: 50,
			x1:      0,
			y1:      0,
			x2:      100,
			y2:      100,
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := checkOverlap(tt.radius, tt.xCenter, tt.yCenter, tt.x1, tt.y1, tt.x2, tt.y2)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
