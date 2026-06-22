package easy

import (
	"reflect"
	"testing"
)

func TestMaximumNumberOfBaloons(t *testing.T) {
	tests := []struct {
		name string
		text string
		want int
	}{
		{
			name: "example 1",
			text: "nlaebolko",
			want: 1,
		},
		{
			name: "example 2",
			text: "loonbalxballpoon",
			want: 2,
		},
		{
			name: "example 3",
			text: "leetcode",
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := maxNumberOfBalloons(tt.text)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
