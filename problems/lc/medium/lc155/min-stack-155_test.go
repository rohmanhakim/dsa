package lc155

import (
	"reflect"
	"testing"

	"github.com/rohmanhakim/leetcode/lib"
)

func TestMinStack(t *testing.T) {
	tests := []struct {
		name   string
		inputs []string
		values [][]int
		want   []*int
	}{
		{
			name:   "example 1",
			inputs: []string{"MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"},
			values: [][]int{{}, {-2}, {0}, {-3}, {}, {}, {}, {}},
			want:   []*int{nil, nil, nil, nil, lib.Ptr(-3), nil, lib.Ptr(0), lib.Ptr(-2)},
		},
		{
			name:   "example 2",
			inputs: []string{"MinStack", "push", "push", "push", "push", "push", "push", "getMin", "pop", "pop", "pop", "getMin", "pop", "pop", "getMin"},
			values: [][]int{{}, {7}, {2}, {5}, {1}, {3}, {4}, {}, {}, {}, {}, {}, {}, {}, {}},
			want:   []*int{nil, nil, nil, nil, nil, nil, nil, lib.Ptr(1), nil, nil, nil, lib.Ptr(2), nil, nil, lib.Ptr(7)},
		},
		{
			name:   "example 3",
			inputs: []string{"MinStack", "push", "push", "push", "getMin", "pop", "getMin", "pop", "getMin", "pop", "push", "push", "push", "getMin", "pop", "top", "getMin", "pop", "getMin", "pop"},
			values: [][]int{{}, {0}, {1}, {0}, {}, {}, {}, {}, {}, {}, {-2}, {-1}, {-2}, {}, {}, {}, {}, {}, {}, {}},
			want:   []*int{nil, nil, nil, nil, lib.Ptr(0), nil, lib.Ptr(0), nil, lib.Ptr(0), nil, nil, nil, nil, lib.Ptr(-2), nil, lib.Ptr(-1), lib.Ptr(-2), nil, lib.Ptr(-2), nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ms MinStack
			outputs := make([]*int, len(tt.values))
			for i, op := range tt.inputs {
				switch op {
				case "MinStack":
					ms = Constructor()
					outputs[i] = nil
				case "push":
					ms.Push(tt.values[i][0])
					outputs[i] = nil
				case "pop":
					ms.Pop()
					outputs[i] = nil
				case "top":
					outputs[i] = lib.Ptr(ms.Top())
				case "getMin":
					outputs[i] = lib.Ptr(ms.GetMin())
				}
			}

			if !reflect.DeepEqual(tt.want, outputs) {
				t.Fatalf("got %v want %v", outputs, tt.want)
			}
		})
	}
}
