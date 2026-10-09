package cf231a

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestTeam drives Team() end-to-end like the
// online judge: stdin comes from a temp input file, stdout is
// captured from a temp output file, then compared.
func TestTeam(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{

		{
			name:  "sample 1",
			input: "4\n1 1 0\n0 0 1\n1 1 1\n1 0 0",
			want:  "2",
		},

		{
			name:  "sample 2",
			input: "1\n1 1 1",
			want:  "1\n",
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

			Team()

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
