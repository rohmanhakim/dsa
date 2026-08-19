package lc1386

import (
	"reflect"
	"testing"
)

func TestMaxNumberOfFamilies(t *testing.T) {
	tests := []struct {
		name          string
		n             int
		reservedSeats [][]int
		want          int
	}{

		{
			name:          "example 1",
			n:             3,
			reservedSeats: [][]int{{1, 2}, {1, 3}, {1, 8}, {2, 6}, {3, 1}, {3, 10}},
			want:          4,
		},

		{
			name:          "example 2",
			n:             2,
			reservedSeats: [][]int{{2, 1}, {1, 8}, {2, 6}},
			want:          2,
		},

		{
			name:          "example 3",
			n:             4,
			reservedSeats: [][]int{{4, 3}, {1, 4}, {4, 6}, {1, 7}},
			want:          4,
		},

		{
			name:          "example 4",
			n:             2,
			reservedSeats: [][]int{{1, 5}, {2, 8}, {2, 10}, {2, 2}, {1, 6}, {1, 10}, {1, 1}, {2, 5}, {1, 2}},
			want:          0,
		},

		{
			name:          "example 5",
			n:             2,
			reservedSeats: [][]int{{1, 6}, {1, 8}, {1, 3}, {2, 3}, {1, 10}, {1, 2}, {1, 5}, {2, 2}, {2, 4}, {2, 10}, {1, 7}, {2, 5}},
			want:          1,
		},

		{
			name:          "example 6",
			n:             3,
			reservedSeats: [][]int{{2, 3}},
			want:          5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := maxNumberOfFamilies(tt.n, tt.reservedSeats)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
