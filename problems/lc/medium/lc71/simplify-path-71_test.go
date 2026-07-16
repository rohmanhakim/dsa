package lc71

import (
	"reflect"
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{

		{
			name: "example 1",
			path: "/home/",
			want: "/home",
		},

		{
			name: "example 2",
			path: "/home//foo/",
			want: "/home/foo",
		},

		{
			name: "example 3",
			path: "/home/user/Documents/../Pictures",
			want: "/home/user/Pictures",
		},

		{
			name: "example 4",
			path: "/../",
			want: "/",
		},

		{
			name: "example 5",
			path: "/.../a/../b/c/../d/./",
			want: "/.../b/d",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := simplifyPath(tt.path)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
