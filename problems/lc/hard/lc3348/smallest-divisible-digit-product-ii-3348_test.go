package lc3348

import (
	"reflect"
	"testing"
)

func TestSmallestNumber(t *testing.T) {
	tests := []struct {
		name string
		num  string
		t    int64
		want string
	}{

		{
			name: "example 1",
			num:  "1234",
			t:    256,
			want: "1488",
		},

		{
			name: "example 2",
			num:  "12355",
			t:    50,
			want: "12355",
		},

		{
			name: "example 3",
			num:  "11111",
			t:    26,
			want: "-1",
		},

		{
			name: "example 4",
			num:  "10",
			t:    320,
			want: "588",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := smallestNumber(tt.num, tt.t)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
