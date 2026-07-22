package lc3501

import (
	"reflect"
	"testing"
)

func TestMaxActiveSectionsAfterTrade(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		queries [][]int
		want    []int
	}{

		{
			name:    "example 1",
			s:       "01",
			queries: [][]int{{0, 1}},
			want:    []int{1},
		},

		{
			name:    "example 2",
			s:       "0100",
			queries: [][]int{{0, 3}, {0, 2}, {1, 3}, {2, 3}},
			want:    []int{4, 3, 1, 1},
		},

		{
			name:    "example 3",
			s:       "1000100",
			queries: [][]int{{1, 5}, {0, 6}, {0, 4}},
			want:    []int{6, 7, 2},
		},

		{
			name:    "example 4",
			s:       "01010",
			queries: [][]int{{0, 3}, {1, 4}, {1, 3}},
			want:    []int{4, 4, 2},
		},

		{
			name:    "example 5",
			s:       "1",
			queries: [][]int{{0, 0}},
			want:    []int{1},
		},

		{
			name:    "example 6",
			s:       "0",
			queries: [][]int{{0, 0}},
			want:    []int{0},
		},

		{
			name:    "example 7",
			s:       "11",
			queries: [][]int{{0, 0}, {0, 1}},
			want:    []int{2, 2},
		},

		{
			name:    "example 8",
			s:       "10001110011111000010000011110000",
			queries: [][]int{{3, 29}, {4, 6}, {5, 10}, {8, 15}},
			want:    []int{23, 14, 14, 17},
		},

		{
			name:    "example 9",
			s:       "0101110001101",
			queries: [][]int{{0, 7}, {0, 12}},
			want:    []int{10, 11},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := maxActiveSectionsAfterTrade(tt.s, tt.queries)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
