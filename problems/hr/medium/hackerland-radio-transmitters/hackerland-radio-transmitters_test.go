package hackerlandradiotransmitters

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestImplementation drives Implementation() via hackerlandRadioTransmitters():
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
			input: "5 1\n1 2 3 5 9",
			want:  "3\n",
		},

		{
			name:  "example 2",
			input: "5 1\n1 2 3 4 5",
			want:  "2\n",
		},

		{
			name:  "example 3",
			input: "8 2\n7 2 4 6 5 9 12 11",
			want:  "3\n",
		},

		{
			name:  "example 4",
			input: "1 1\n100",
			want:  "1\n",
		},

		{
			name:  "example 5",
			input: "2 1\n1 100",
			want:  "2\n",
		},

		{
			name:  "example 6",
			input: "2 100\n1 100",
			want:  "1\n",
		},

		{
			name:  "example 7",
			input: "6 3\n1 4 7 10 13 16",
			want:  "2\n",
		},

		{
			name:  "example 8",
			input: "3 5\n1 2 3",
			want:  "1\n",
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
