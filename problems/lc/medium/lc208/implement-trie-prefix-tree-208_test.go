package lc208

import (
	"reflect"
	"testing"
)

func TestImplementTriePrefixTree(t *testing.T) {
	tests := []struct {
		name     string
		commands []string
		param    [][]string
		want     []bool
	}{
		{
			name: "example 1",
			commands: []string{
				"Trie", "insert", "search", "search", "startsWith", "insert", "search",
			},
			param: [][]string{{}, {"apple"}, {"apple"}, {"app"}, {"app"}, {"app"}, {"app"}},
			want:  []bool{true, false, true, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comms := tt.commands
			var trie Trie
			out := []bool{}
			for i, comm := range comms {
				switch comm {
				case "Trie":
					trie = Constructor()
				case "insert":
					trie.Insert(tt.param[i][0])
				case "search":
					res := trie.Search(tt.param[i][0])
					out = append(out, res)
				case "startsWith":
					res := trie.StartsWith(tt.param[i][0])
					out = append(out, res)
				}
			}

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
