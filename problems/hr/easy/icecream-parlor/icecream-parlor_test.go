package icecreamparlor

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestImplementation drives Implementation() via icecreamParlor():
// stdin is redirected to a temp input file, OUTPUT_PATH to
// a temp output file, then the output is compared.
func TestImplementation(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{

		{
			name:  "example 1",
			input: "2\n4\n5\n1 4 5 3 2\n4\n4\n2 2 4 3",
			want:  "1 4\n1 2\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Redirect os.Stdin to a temp file holding the input.
			inFile, err := os.CreateTemp(t.TempDir(), "input")
			if err != nil {
				t.Fatalf("create input file: %v", err)
			}
			if _, err := io.WriteString(inFile, tt.input); err != nil {
				t.Fatalf("write input file: %v", err)
			}
			if _, err := inFile.Seek(0, io.SeekStart); err != nil {
				t.Fatalf("rewind input file: %v", err)
			}

			oldStdin := os.Stdin
			os.Stdin = inFile
			defer func() {
				os.Stdin = oldStdin
				inFile.Close()
			}()

			outPath := filepath.Join(t.TempDir(), "output")
			t.Setenv("OUTPUT_PATH", outPath)

			Implementation()

			got, err := os.ReadFile(outPath)
			if err != nil {
				t.Fatalf("read output file: %v", err)
			}

			if hrNormalize(string(got)) != hrNormalize(tt.want) {
				t.Fatalf("output mismatch:\ngot:  %q\nwant: %q", string(got), tt.want)
			}
		})
	}
}

// hrNormalize trims trailing whitespace per line and overall
// so cosmetic newline differences do not fail the test.
func hrNormalize(s string) string {
	lines := strings.Split(s, "\n")
	for i, l := range lines {
		lines[i] = strings.TrimRight(l, " \t\r")
	}
	return strings.TrimSpace(strings.Join(lines, "\n"))
}
