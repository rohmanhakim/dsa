package easy

import (
	"reflect"
	"testing"
)

func TestWeightedWordMapping(t *testing.T) {
	tests := []struct {
		name    string
		words   []string
		weights []int
		want    string
	}{
		{
			name:    "test 1",
			words:   []string{"abcd", "def", "xyz"},
			weights: []int{5, 3, 12, 14, 1, 2, 3, 2, 10, 6, 6, 9, 7, 8, 7, 10, 8, 9, 6, 9, 9, 8, 3, 7, 7, 2},
			want:    "rij",
		},
		{
			name:    "test 1",
			words:   []string{"a", "b", "c"},
			weights: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			want:    "yyy",
		},
		{
			name:    "test 3",
			words:   []string{"abcd"},
			weights: []int{7, 5, 3, 4, 3, 5, 4, 9, 4, 2, 2, 7, 10, 2, 5, 10, 6, 1, 2, 2, 4, 1, 3, 4, 4, 5},
			want:    "g",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := mapWordWeights(tt.words, tt.weights)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
