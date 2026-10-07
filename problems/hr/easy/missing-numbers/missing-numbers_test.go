package missingnumbers

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestImplementation drives Implementation() via missingNumbers():
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
			input: "10\n203 204 205 206 207 208 203 204 205 206\n13\n203 204 204 205 206 207 205 208 203 206 205 206 204",
			want:  "204 205 206\n",
		},

		{
			name:  "example 2",
			input: "10\n11 4 11 7 13 4 12 11 10 14\n15\n11 4 11 7 3 7 10 13 4 8 12 11 10 14 12",
			want:  "3 7 8 10 12\n",
		},

		{
			name:  "example 3",
			input: "1\n100\n3\n100 200 200",
			want:  "200\n",
		},

		{
			name:  "example 4",
			input: "1\n100\n4\n100 200 200 200",
			want:  "200\n",
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
