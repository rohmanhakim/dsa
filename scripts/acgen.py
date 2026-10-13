#!/usr/bin/env python3
"""acgen.py - Generate Go test files for AtCoder (and Codeforces) problems.

Unlike HackerRank (hrgen.py), AtCoder/Codeforces problems read standard input
and write standard output. The solution file exposes a single exported
function (e.g. `func Product()`) that the test drives end-to-end:
    1. the pasted input text is embedded verbatim in the test file
    2. os.Stdin is swapped for a temp file holding that input
    3. os.Stdout is redirected to a temp output file
    4. the exported solution function is called
    5. streams are restored and the output file is compared
       (whitespace-normalized) against the pasted expected output

Usage:
    1. Navigate to the solution file's directory
    2. Run: python3 ../../scripts/acgen.py
    3. Paste input/output blocks (Ctrl+D to finish)
    4. The script generates a _test.go file

Input format (block markers; multiple pairs = multiple test cases):
    === input ===
    <raw stdin lines>
    === output ===
    <raw expected stdout lines>

Options:
    -i, --input FILE    Read input from file instead of stdin
    -f, --func NAME     Override function name
    -s, --solution FILE Specify solution file path
    --dry-run           Print generated test file without writing
"""

import argparse
import re
import sys
from pathlib import Path

try:
    import levels
except ImportError:  # standalone copy without levels.py next to it
    levels = None


def die(msg) -> None:
    """Print error message and exit."""
    print(f"Error: {msg}", file=sys.stderr)
    sys.exit(1)  # type: ignore[misc]


def find_solution_file(directory):
    """Find the single .go solution file (non-test) in the directory."""
    go_files = [f for f in Path(directory).glob("*.go")
                if not f.name.endswith("_test.go")]
    if len(go_files) == 0:
        die("No .go solution file found in current directory.")
    if len(go_files) > 1:
        die(f"Multiple .go files found: {[f.name for f in go_files]}. "
            f"Use --solution to specify.")
    return go_files[0]


def parse_solution(filepath, func_override=None):
    """Extract package name and the exported solution function name.

    AtCoder/Codeforces stubs expose exactly one exported function (uppercase)
    that reads os.Stdin and writes os.Stdout. --func skips the
    single-function validation when helpers are exported too.
    """
    content = filepath.read_text()

    pkg_match = re.search(r"^package\s+(\w+)", content, re.MULTILINE)
    if pkg_match is None:
        die(f"No package declaration found in {filepath}")
    assert pkg_match is not None  # die() exits above
    package = pkg_match.group(1)

    exported = re.findall(r"^func\s+([A-Z]\w*)\s*\(",
                          content, re.MULTILINE)

    if func_override:
        if func_override not in exported:
            die(f"Function '{func_override}' not found among exported "
                f"functions: {exported or 'none'}.")
        return package, func_override

    if len(exported) == 0:
        die("No exported function found in solution file.\n"
            "acgen.py expects an AtCoder/Codeforces stub whose exported "
            "function (e.g. func Product()) reads os.Stdin and writes "
            "os.Stdout.")
    if len(exported) > 1:
        die(f"Multiple exported functions found: {exported}. Use --func to "
            f"pick the one that performs the I/O.")
    return package, exported[0]


def parse_blocks(text):
    """Parse '=== input ===' / '=== output ===' blocks into (in, out) pairs.

    Content is kept verbatim (no stripping) so the embedded input matches
    what the judge would feed the program. Block marker lines are matched on
    their stripped form.
    """
    pairs = []
    mode = None  # None | "input" | "output"
    buf = []
    current_input = None

    def flush_output():
        assert current_input is not None
        pairs.append((current_input, "\n".join(buf)))

    for raw_line in text.split("\n"):
        stripped = raw_line.strip()
        if stripped == "=== input ===":
            if mode == "input":
                die("Found '=== input ===' while an input block is open. "
                    "Missing '=== output ==='?")
            if mode == "output":
                flush_output()
            mode = "input"
            buf = []
            continue
        if stripped == "=== output ===":
            if mode is None:
                die("Found '=== output ===' before any '=== input ==='.")
            if mode == "output":
                die("Found '=== output ===' while an output block is open.")
            current_input = "\n".join(buf)
            mode = "output"
            buf = []
            continue
        if mode is not None:
            buf.append(raw_line)

    if mode == "input":
        die("Input block at end of paste is missing its '=== output ===' "
            "block.")
    if mode == "output":
        flush_output()

    if not pairs:
        die("No input/output blocks found.\n"
            "Expected format:\n"
            "  === input ===\n"
            "  <stdin lines>\n"
            "  === output ===\n"
            "  <expected stdout lines>")

    return pairs


def go_string_literal(s):
    """Convert a Python string to a Go interpreted string literal."""
    out = []
    for ch in s:
        if ch == "\\":
            out.append("\\\\")
        elif ch == '"':
            out.append('\\"')
        elif ch == "\n":
            out.append("\\n")
        elif ch == "\t":
            out.append("\\t")
        elif ch == "\r":
            out.append("\\r")
        elif ord(ch) < 0x20:
            out.append(f"\\x{ord(ch):02x}")
        else:
            out.append(ch)
    return '"' + "".join(out) + '"'


def generate_test(package, func_name, pairs):
    """Generate the end-to-end test file for the exported solution function."""
    lines = []
    lines.append(f"package {package}")
    lines.append("")
    lines.append("import (")
    lines.append('\t"io"')
    lines.append('\t"os"')
    lines.append('\t"path/filepath"')
    lines.append('\t"strings"')
    lines.append('\t"testing"')
    lines.append(")")
    lines.append("")
    lines.append(f"// Test{func_name} drives {func_name}() end-to-end like the")
    lines.append("// online judge: stdin comes from a temp input file, stdout is")
    lines.append("// captured from a temp output file, then compared.")
    lines.append(f"func Test{func_name}(t *testing.T) {{")
    lines.append("\ttests := []struct {")
    lines.append("\t\tname  string")
    lines.append("\t\tinput string")
    lines.append("\t\twant  string")
    lines.append("\t}{")
    lines.append("")

    for i, (inp, out) in enumerate(pairs, 1):
        if i > 1:
            lines.append("")
        lines.append("\t\t{")
        lines.append(f'\t\t\tname:  "sample {i}",')
        lines.append(f"\t\t\tinput: {go_string_literal(inp)},")
        lines.append(f"\t\t\twant:  {go_string_literal(out)},")
        lines.append("\t\t},")

    lines.append("\t}")
    lines.append("")
    lines.append("\tfor _, tt := range tests {")
    lines.append("\t\tt.Run(tt.name, func(t *testing.T) {")
    lines.append("\t\t\tdir := t.TempDir()")
    lines.append("")
    lines.append('\t\t\tinFile, err := os.Create(filepath.Join(dir, "input"))')
    lines.append("\t\t\tif err != nil {")
    lines.append('\t\t\t\tt.Fatalf("create input file: %v", err)')
    lines.append("\t\t\t}")
    lines.append("\t\t\tif _, err := io.WriteString(inFile, tt.input); err != nil {")
    lines.append('\t\t\t\tt.Fatalf("write input file: %v", err)')
    lines.append("\t\t\t}")
    lines.append("\t\t\tif _, err := inFile.Seek(0, io.SeekStart); err != nil {")
    lines.append('\t\t\t\tt.Fatalf("rewind input file: %v", err)')
    lines.append("\t\t\t}")
    lines.append("")
    lines.append('\t\t\toutFile, err := os.Create(filepath.Join(dir, "output"))')
    lines.append("\t\t\tif err != nil {")
    lines.append('\t\t\t\tt.Fatalf("create output file: %v", err)')
    lines.append("\t\t\t}")
    lines.append("")
    lines.append("\t\t\toldStdin, oldStdout := os.Stdin, os.Stdout")
    lines.append("\t\t\tos.Stdin, os.Stdout = inFile, outFile")
    lines.append("\t\t\tdefer func() {")
    lines.append("\t\t\t\tos.Stdin, os.Stdout = oldStdin, oldStdout")
    lines.append("\t\t\t\tinFile.Close()")
    lines.append("\t\t\t\toutFile.Close()")
    lines.append("\t\t\t}()")
    lines.append("")
    lines.append(f"\t\t\t{func_name}()")
    lines.append("")
    lines.append("\t\t\t// restore before reading so all output reached the file")
    lines.append("\t\t\tos.Stdin, os.Stdout = oldStdin, oldStdout")
    lines.append("\t\t\tif err := outFile.Close(); err != nil {")
    lines.append('\t\t\t\tt.Fatalf("close output file: %v", err)')
    lines.append("\t\t\t}")
    lines.append("")
    lines.append("\t\t\tgot, err := os.ReadFile(outFile.Name())")
    lines.append("\t\t\tif err != nil {")
    lines.append('\t\t\t\tt.Fatalf("read output file: %v", err)')
    lines.append("\t\t\t}")
    lines.append("")
    lines.append("\t\t\tif acNormalize(string(got)) != acNormalize(tt.want) {")
    lines.append('\t\t\t\tt.Fatalf("output mismatch:\\n'
                 'got:  %q\\nwant: %q", string(got), tt.want)')
    lines.append("\t\t\t}")
    lines.append("\t\t})")
    lines.append("\t}")
    lines.append("}")
    lines.append("")
    lines.append("// acNormalize trims trailing whitespace per line and overall")
    lines.append("// so cosmetic newline differences do not fail the test.")
    lines.append("func acNormalize(s string) string {")
    lines.append('\tlines := strings.Split(s, "\\n")')
    lines.append("\tfor i, l := range lines {")
    lines.append('\t\tlines[i] = strings.TrimRight(l, " \\t\\r")')
    lines.append("\t}")
    lines.append('\treturn strings.TrimSpace(strings.Join(lines, "\\n"))')
    lines.append("}")
    lines.append("")

    return "\n".join(lines)


def main():
    parser = argparse.ArgumentParser(
        description="Generate Go test files for AtCoder problems"
    )
    parser.add_argument(
        "-i", "--input",
        help="File containing '=== input ===' / '=== output ===' blocks"
    )
    parser.add_argument(
        "-f", "--func",
        help="Function name to test (overrides auto-detection)"
    )
    parser.add_argument(
        "-s", "--solution",
        help="Path to solution .go file"
    )
    parser.add_argument(
        "--dry-run",
        action="store_true",
        help="Print generated test file without writing"
    )
    args = parser.parse_args()

    # Find solution file
    if args.solution:
        solution_path = Path(args.solution)
    else:
        solution_path = find_solution_file(".")

    if not solution_path.exists():
        die(f"Solution file not found: {solution_path}")

    # Parse solution
    package, func_name = parse_solution(solution_path, func_override=args.func)

    print(f"Solution:    {solution_path}", file=sys.stderr)
    print(f"Package:     {package}", file=sys.stderr)
    print(f"Solution fn: {func_name}()", file=sys.stderr)
    if levels:
        levels.print_goal(solution_path)

    # Check for existing test file
    test_file = solution_path.with_name(solution_path.stem + "_test.go")
    if test_file.exists():
        die(f"Test file already exists: {test_file}\n"
            f"Please delete or rename it first.")

    # Get input
    if args.input:
        text = Path(args.input).read_text()
    else:
        print("\nPaste input/output blocks (Ctrl+D to finish):",
              file=sys.stderr)
        text = sys.stdin.read()

    # Parse blocks
    pairs = parse_blocks(text)
    print(f"Cases:       {len(pairs)}", file=sys.stderr)

    # Generate
    content = generate_test(package, func_name, pairs)

    if args.dry_run:
        sys.stdout.write(content)
    else:
        test_file.write_text(content)
        print(f"\nGenerated: {test_file}", file=sys.stderr)


if __name__ == "__main__":
    main()



