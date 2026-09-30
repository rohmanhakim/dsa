package lc1807

import (
	"reflect"
	"testing"
)

func TestEvaluate(t *testing.T) {
	tests := []struct {
		name      string
		s         string
		knowledge [][]string
		want      string
	}{

		{
			name:      "example 1",
			s:         "(name)is(age)yearsold",
			knowledge: [][]string{{"name", "bob"}, {"age", "two"}},
			want:      "bobistwoyearsold",
		},

		{
			name:      "example 2",
			s:         "hi(name)",
			knowledge: [][]string{{"a", "b"}},
			want:      "hi?",
		},

		{
			name:      "example 3",
			s:         "(a)(a)(a)aaa",
			knowledge: [][]string{{"a", "yes"}},
			want:      "yesyesyesaaa",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := evaluate(tt.s, tt.knowledge)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
