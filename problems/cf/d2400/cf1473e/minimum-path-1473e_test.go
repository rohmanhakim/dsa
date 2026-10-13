package cf1473e

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestMinimumPath drives MinimumPath() end-to-end like the
// online judge: stdin comes from a temp input file, stdout is
// captured from a temp output file, then compared.
func TestMinimumPath(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{

		{
			name:  "sample 1",
			input: "5 4\n5 3 4\n2 1 1\n3 2 2\n2 4 2",
			want:  "1 2 2 4 \n",
		},

		{
			name:  "sample 2",
			input: "6 8\n3 1 1\n3 6 2\n5 4 2\n4 2 2\n6 1 1\n5 2 1\n3 2 3\n1 5 4",
			want:  "2 1 4 3 1 \n",
		},

		{
			name:  "sample 3",
			input: "7 10\n7 5 5\n2 3 3\n4 7 1\n5 3 6\n2 7 6\n6 2 6\n3 7 6\n4 2 1\n3 1 4\n1 7 4",
			want:  "3 4 2 7 7 3\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()

			inFile, err := os.Create(filepath.Join(dir, "input"))
			if err != nil {
				t.Fatalf("create input file: %v", err)
			}
			if _, err := io.WriteString(inFile, tt.input); err != nil {
				t.Fatalf("write input file: %v", err)
			}
			if _, err := inFile.Seek(0, io.SeekStart); err != nil {
				t.Fatalf("rewind input file: %v", err)
			}

			outFile, err := os.Create(filepath.Join(dir, "output"))
			if err != nil {
				t.Fatalf("create output file: %v", err)
			}

			oldStdin, oldStdout := os.Stdin, os.Stdout
			os.Stdin, os.Stdout = inFile, outFile
			defer func() {
				os.Stdin, os.Stdout = oldStdin, oldStdout
				inFile.Close()
				outFile.Close()
			}()

			MinimumPath()

			// restore before reading so all output reached the file
			os.Stdin, os.Stdout = oldStdin, oldStdout
			if err := outFile.Close(); err != nil {
				t.Fatalf("close output file: %v", err)
			}

			got, err := os.ReadFile(outFile.Name())
			if err != nil {
				t.Fatalf("read output file: %v", err)
			}

			if acNormalize(string(got)) != acNormalize(tt.want) {
				t.Fatalf("output mismatch:\ngot:  %q\nwant: %q", string(got), tt.want)
			}
		})
	}
}

// acNormalize trims trailing whitespace per line and overall
// so cosmetic newline differences do not fail the test.
func acNormalize(s string) string {
	lines := strings.Split(s, "\n")
	for i, l := range lines {
		lines[i] = strings.TrimRight(l, " \t\r")
	}
	return strings.TrimSpace(strings.Join(lines, "\n"))
}
