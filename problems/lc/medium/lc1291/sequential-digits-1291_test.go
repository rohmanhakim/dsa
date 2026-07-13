package lc1291

import (
	"reflect"
	"testing"
)

func TestSequentialDigits(t *testing.T) {
	tests := []struct {
		name string
		low  int
		high int
		want []int
	}{
		{
			name: "example 1",
			low:  100,
			high: 300,
			want: []int{123, 234},
		},
		{
			name: "example 2",
			low:  1000,
			high: 13000,
			want: []int{1234, 2345, 3456, 4567, 5678, 6789, 12345},
		},
		{
			name: "example 3",
			low:  10,
			high: 1000000000,
			want: []int{12, 23, 34, 45, 56, 67, 78, 89, 123, 234, 345, 456, 567, 678, 789, 1234, 2345, 3456, 4567, 5678, 6789, 12345, 23456, 34567, 45678, 56789, 123456, 234567, 345678, 456789, 1234567, 2345678, 3456789, 12345678, 23456789, 123456789},
		},
		{
			name: "example 4",
			low:  58,
			high: 155,
			want: []int{67, 78, 89, 123},
		},
		{
			name: "example 5",
			low:  89,
			high: 234,
			want: []int{89, 123, 234},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := sequentialDigits(tt.low, tt.high)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
