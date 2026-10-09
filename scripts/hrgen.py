#!/usr/bin/env python3
"""
hrgen.py - Generate Go test files for HackerRank problems.

Unlike LeetCode problems (see gengen.py), HackerRank problems expose an
exported `Implementation()` function that reads test input from os.Stdin
and writes results to the file named by the OUTPUT_PATH environment
variable. The input format is therefore not declared anywhere -- it is
implicit in the read statements inside Implementation().

This script does NOT parse that format. It drives Implementation()
end-to-end, exactly like HackerRank itself does:
    1. the pasted input text is embedded verbatim in the test file
    2. the test swaps os.Stdin for a temp file holding that input
    3. OUTPUT_PATH is pointed at a temp file via t.Setenv
    4. Implementation() is called
    5. the output file is compared (whitespace-normalized) against the
       pasted expected output

Usage:
    1. Navigate to the solution file's directory
    2. Run: python3 hrgen.py
    3. Paste input/output blocks (Ctrl+D to finish)
    4. The script generates a _test.go file

Input format (block markers; multiple pairs = multiple test cases):
    === input ===
    <raw stdin lines>
    === output ===
    <raw expected stdout lines>

Options:
    -i, --input FILE    Read input from file instead of stdin
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


def parse_solution(filepath):
    """Extract package name and inner solution function name from an HR stub.

    Verification of `func Implementation` is the key difference from
    gengen.py: it confirms the file is a HackerRank-style stub that reads
    os.Stdin and writes OUTPUT_PATH.
    """
    content = filepath.read_text()

    pkg_match = re.search(r"^package\s+(\w+)", content, re.MULTILINE)
    if pkg_match is None:
        die(f"No package declaration found in {filepath}")
    assert pkg_match is not None  # die() exits above
    package = pkg_match.group(1)

    if not re.search(r"^func\s+Implementation\s*\(", content, re.MULTILINE):
        die(f"No 'func Implementation' found in {filepath}.\n"
            "hrgen.py expects a HackerRank stub where Implementation() "
            "reads os.Stdin and writes to OUTPUT_PATH.")

    # Best-effort: name of the function Implementation() delegates to.
    # Display only -- the generated test always calls Implementation().
    inner_match = re.search(r"result\s*:=\s*(\w+)\s*\(", content)
    inner_func = inner_match.group(1) if inner_match else None

    return package, inner_func



def parse_blocks(text):
    """Parse '=== input ===' / '=== output ===' blocks into (in, out) pairs.

    Content is kept verbatim (no stripping) so the embedded input matches
    what HackerRank would feed to the program. Block marker lines are
    matched on their stripped form.
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


def generate_test(package, inner_func, pairs):
    """Generate the end-to-end test file for Implementation()."""
    comment = f"via {inner_func}()" if inner_func else "end-to-end"
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
    lines.append(f"// TestImplementation drives Implementation() {comment}:")
    lines.append("// stdin is redirected to a temp input file, OUTPUT_PATH to")
    lines.append("// a temp output file, then the output is compared.")
    lines.append("func TestImplementation(t *testing.T) {")
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
        lines.append(f'\t\t\tname:  "example {i}",')
        lines.append(f"\t\t\tinput: {go_string_literal(inp)},")
        lines.append(f"\t\t\twant:  {go_string_literal(out)},")
        lines.append("\t\t},")

    lines.append("\t}")
    lines.append("")
    lines.append("\tfor _, tt := range tests {")
    lines.append("\t\tt.Run(tt.name, func(t *testing.T) {")
    lines.append("\t\t\t// Redirect os.Stdin to a temp file holding the input.")
    lines.append('\t\t\tinFile, err := os.CreateTemp(t.TempDir(), "input")')
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
    lines.append("\t\t\toldStdin := os.Stdin")
    lines.append("\t\t\tos.Stdin = inFile")
    lines.append("\t\t\tdefer func() {")
    lines.append("\t\t\t\tos.Stdin = oldStdin")
    lines.append("\t\t\t\tinFile.Close()")
    lines.append("\t\t\t}()")
    lines.append("")
    lines.append('\t\t\toutPath := filepath.Join(t.TempDir(), "output")')
    lines.append('\t\t\tt.Setenv("OUTPUT_PATH", outPath)')
    lines.append("")
    lines.append("\t\t\tImplementation()")
    lines.append("")
    lines.append("\t\t\tgot, err := os.ReadFile(outPath)")
    lines.append("\t\t\tif err != nil {")
    lines.append('\t\t\t\tt.Fatalf("read output file: %v", err)')
    lines.append("\t\t\t}")
    lines.append("")
    lines.append("\t\t\tif hrNormalize(string(got)) != hrNormalize(tt.want) {")
    lines.append('\t\t\t\tt.Fatalf("output mismatch:\\n'
                 'got:  %q\\nwant: %q", string(got), tt.want)')
    lines.append("\t\t\t}")
    lines.append("\t\t})")
    lines.append("\t}")
    lines.append("}")
    lines.append("")
    lines.append("// hrNormalize trims trailing whitespace per line and overall")
    lines.append("// so cosmetic newline differences do not fail the test.")
    lines.append("func hrNormalize(s string) string {")
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
        description="Generate Go test files for HackerRank problems"
    )
    parser.add_argument(
        "-i", "--input",
        help="File containing '=== input ===' / '=== output ===' blocks"
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

    package, inner_func = parse_solution(solution_path)

    print(f"Solution:    {solution_path}", file=sys.stderr)
    print(f"Package:     {package}", file=sys.stderr)
    if inner_func:
        print(f"Solution fn: {inner_func}()", file=sys.stderr)
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
    content = generate_test(package, inner_func, pairs)

    if args.dry_run:
        sys.stdout.write(content)
    else:
        test_file.write_text(content)
        print(f"\nGenerated: {test_file}", file=sys.stderr)


if __name__ == "__main__":
    main()
