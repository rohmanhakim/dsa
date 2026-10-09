#!/usr/bin/env python3
"""
gengen.py - Generate Go test files for LeetCode problems from copied input/output.

Usage:
    1. Navigate to the solution file's directory
    2. Run: python3 gengen.py
    3. Paste LeetCode input/output (Ctrl+D to finish)
    4. The script generates a _test.go file

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
    """Extract package name, function name, parameters, and return type.

    If func_override is provided, use that function name instead of
    auto-detecting (skips the single-function validation).
    """
    content = filepath.read_text()

    # Package declaration
    pkg_match = re.search(r"^package\s+(\w+)", content, re.MULTILINE)
    if pkg_match is None:
        die(f"No package declaration found in {filepath}")
    assert pkg_match is not None  # die() exits above
    package = pkg_match.group(1)

    # Find unexported functions (start with lowercase)
    func_starts = list(re.finditer(
        r"^func\s+([a-z]\w*)\s*\(", content, re.MULTILINE
    ))

    functions = {}
    for match in func_starts:
        func_name = match.group(1)
        start = match.end()  # position after the opening '('

        # Find matching closing ')'
        depth = 1
        i = start
        while i < len(content) and depth > 0:
            if content[i] == "(":
                depth += 1
            elif content[i] == ")":
                depth -= 1
            i += 1

        if depth != 0:
            continue

        params_str = content[start:i - 1].strip()

        # After the closing ')', find the return type
        rest = content[i:].strip()
        ret_match = re.match(r"^(\([^)]*\)|\S+)", rest)
        return_type = ret_match.group(1) if ret_match else ""

        # Parse params: "name type, name type" -> [("name", "type"), ...]
        params = []
        if params_str:
            for p in params_str.split(","):
                p = p.strip()
                if not p:
                    continue
                parts = p.rsplit(None, 1)
                if len(parts) == 2:
                    params.append((parts[0], parts[1]))

        functions[func_name] = (params, return_type)

    if len(functions) == 0:
        die("No unexported functions found in solution file.")

    if func_override:
        if func_override not in functions:
            names = list(functions.keys())
            die(f"Function '{func_override}' not found in solution file. "
                f"Available: {names}")
        params, return_type = functions[func_override]
        return package, func_override, params, return_type

    if len(functions) > 1:
        names = list(functions.keys())
        die(f"Multiple unexported functions found: {names}. "
            f"Use --func to specify which one to test.")

    name = list(functions.keys())[0]
    params, return_type = functions[name]
    return package, name, params, return_type


def smart_split(text):
    """Split by comma, respecting nested brackets and quoted strings."""
    parts = []
    depth = 0
    in_string = False
    in_rune = False
    current = []
    i = 0
    while i < len(text):
        c = text[i]
        # Handle escape characters inside strings/runes
        if c == "\\" and (in_string or in_rune) and i + 1 < len(text):
            current.append(c)
            current.append(text[i + 1])
            i += 2
            continue
        if c == '"':
            in_string = not in_string
            current.append(c)
        elif c == "'":
            in_rune = not in_rune
            current.append(c)
        elif not in_string and not in_rune:
            if c == "[":
                depth += 1
                current.append(c)
            elif c == "]":
                depth -= 1
                current.append(c)
            elif c == "," and depth == 0:
                parts.append("".join(current).strip())
                current = []
            else:
                current.append(c)
        else:
            current.append(c)
        i += 1
    if current:
        parts.append("".join(current).strip())
    return parts


def parse_value(text):
    """Parse a JSON-like value string into a Python object."""
    text = text.strip()
    if text in ("null", "None", "nil"):
        return None
    if text == "true":
        return True
    if text == "false":
        return False
    if text.startswith('"') and text.endswith('"'):
        return text[1:-1]
    # Handle single-quoted rune literals: 'a' -> "a"
    if text.startswith("'") and text.endswith("'") and len(text) >= 2:
        return text[1:-1]
    if text.startswith("[") and text.endswith("]"):
        inner = text[1:-1].strip()
        if not inner:
            return []
        return [parse_value(e) for e in smart_split(inner)]
    try:
        return int(text)
    except ValueError:
        pass
    try:
        return float(text)
    except ValueError:
        pass
    return text


def infer_go_type(value):
    """Infer Go type from a parsed Python value."""
    if value is None:
        return "any"
    if isinstance(value, bool):
        return "bool"
    if isinstance(value, int):
        return "int"
    if isinstance(value, float):
        return "float64"
    if isinstance(value, str):
        return "string"
    if isinstance(value, list):
        if not value:
            return "[]any"
        # Check if any element is None -> use pointer type
        inner_values = [v for v in value if v is not None]
        if not inner_values:
            return "[]*int"  # all nulls, default to pointer int
        inner_type = infer_go_type(inner_values[0])
        if inner_type.startswith("[]"):
            return "[]" + inner_type
        return "[]" + inner_type
    return "any"


def _rune_literal(ch):
    """Convert a single character to a Go rune literal like 'a'."""
    escape_map = {
        "\\": "\\\\", "'": "\\'", "\n": "\\n", "\r": "\\r",
        "\t": "\\t", "\0": "\\0",
    }
    escaped = escape_map.get(ch, ch)
    return f"'{escaped}'"


def to_go_literal(value, go_type, include_type=True):
    """Convert a parsed Python value to a Go literal string."""
    if value is None:
        return "nil"
    if isinstance(value, bool):
        return "true" if value else "false"
    if isinstance(value, int):
        return str(value)
    if isinstance(value, float):
        return str(value)
    if isinstance(value, str):
        # Rune type: use single-quoted character literal
        if go_type == "rune":
            return _rune_literal(value)
        escaped = value.replace("\\", "\\\\").replace('"', '\\"')
        return f'"{escaped}"'
    if isinstance(value, list):
        prefix = go_type if include_type else ""
        if go_type.startswith("[][]"):
            inner_type = go_type[2:]
            elements = [to_go_literal(v, inner_type, include_type=False)
                        for v in value]
            return f"{prefix}{{{', '.join(elements)}}}"
        elif go_type.startswith("[]"):
            inner_type = go_type[2:]
            elements = [to_go_literal(v, inner_type, include_type=False)
                        for v in value]
            return f"{prefix}{{{', '.join(elements)}}}"
    return str(value)


def parse_input_line(text):
    """Parse 'Input: var = val, var = val' into list of raw value strings."""
    text = text.strip()
    text = re.sub(r"^Input\s*:\s*", "", text)

    assignments = smart_split(text)
    values = []
    for a in assignments:
        if "=" in a:
            _, val = a.split("=", 1)
            values.append(val.strip())
        else:
            values.append(a.strip())
    return values


def parse_output_line(text):
    """Parse 'Output: value' into raw value string."""
    text = text.strip()
    text = re.sub(r"^Output\s*:\s*", "", text)
    return text.strip()


def parse_test_cases(input_text):
    """Parse multiple test cases from LeetCode input/output text.

    Handles formats:
        Example 1:
        Input: ...
        Output: ...

    Or plain:
        Input: ...
        Output: ...
    """
    test_cases = []

    # Split by "Example N:" markers
    blocks = re.split(r"Example\s+\d+\s*:", input_text)
    if len(blocks) <= 1:
        blocks = [input_text]
    else:
        blocks = [b for b in blocks if b.strip()]

    for block in blocks:
        lines = block.strip().split("\n")
        current_input = None

        for line in lines:
            ls = line.strip()
            if ls.startswith("Input:"):
                current_input = ls
            elif ls.startswith("Output:") and current_input is not None:
                input_values = parse_input_line(current_input)
                output_value = parse_output_line(ls)
                test_cases.append((input_values, output_value))
                current_input = None

    return test_cases


def pascal_case(s):
    """Convert camelCase/snake_case function name to PascalCase."""
    # Handle snake_case
    if "_" in s:
        return "".join(word.capitalize() for word in s.split("_"))
    # Handle camelCase: capitalize first char
    return s[0].upper() + s[1:]


def generate_test(package, func_name, params, return_type, test_cases):
    """Generate the Go test file content using string building."""
    test_func_name = "Test" + pascal_case(func_name)

    # Determine want type: prefer the function's return type over inference
    want_type = "any"
    if return_type and return_type not in ("", "void"):
        want_type = return_type
    elif test_cases:
        first_output = parse_value(test_cases[0][1])
        want_type = infer_go_type(first_output)

    lines = []
    lines.append(f"package {package}")
    lines.append("")
    lines.append('import (')
    lines.append('\t"reflect"')
    lines.append('\t"testing"')
    lines.append(')')
    lines.append("")
    lines.append(f"func {test_func_name}(t *testing.T) {{")
    lines.append("\ttests := []struct {")
    lines.append("\t\tname string")

    for pname, ptype in params:
        lines.append(f"\t\t{pname} {ptype}")

    lines.append(f"\t\twant {want_type}")
    lines.append("\t}{")
    lines.append("")

    # Test case blocks
    for i, (input_values, output_raw) in enumerate(test_cases):
        output_val = parse_value(output_raw)
        want_literal = to_go_literal(output_val, want_type)

        lines.append("\t\t{")
        lines.append(f'\t\t\tname: "example {i + 1}",')

        for j, (pname, ptype) in enumerate(params):
            if j < len(input_values):
                val = parse_value(input_values[j])
                val_literal = to_go_literal(val, ptype)
                lines.append(f"\t\t\t{pname}: {val_literal},")

        lines.append(f"\t\t\twant: {want_literal},")
        lines.append("\t\t},")
        lines.append("")

    lines.append("\t}")
    lines.append("")

    # Test loop
    lines.append("\tfor _, tt := range tests {")
    lines.append("\t\tt.Run(tt.name, func(t *testing.T) {")

    call_args = ", ".join(f"tt.{pname}" for pname, _ in params)
    lines.append(f"\t\t\tout := {func_name}({call_args})")
    lines.append("")
    lines.append("\t\t\tif !reflect.DeepEqual(out, tt.want) {")
    lines.append('\t\t\t\tt.Fatalf("got %v want %v", out, tt.want)')
    lines.append("\t\t\t}")
    lines.append("\t\t})")
    lines.append("\t}")
    lines.append("}")
    lines.append("")

    return "\n".join(lines)


def main():
    parser = argparse.ArgumentParser(
        description="Generate Go test files for LeetCode problems"
    )
    parser.add_argument(
        "-i", "--input",
        help="File containing LeetCode input/output"
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

    # Parse solution (--func override passed to skip validation if specified)
    package, func_name, params, return_type = parse_solution(
        solution_path, func_override=args.func
    )

    print(f"Solution:  {solution_path}", file=sys.stderr)
    print(f"Package:   {package}", file=sys.stderr)
    params_str = ", ".join(f"{n} {t}" for n, t in params)
    print(f"Function:  {func_name}({params_str}) {return_type}",
          file=sys.stderr)
    if levels:
        levels.print_goal(solution_path)

    # Check for existing test file
    test_file = solution_path.with_name(solution_path.stem + "_test.go")
    if test_file.exists():
        die(f"Test file already exists: {test_file}\n"
            f"Please delete or rename it first.")

    # Get input
    if args.input:
        input_text = Path(args.input).read_text()
    else:
        print("\nPaste LeetCode input/output (Ctrl+D to finish):",
              file=sys.stderr)
        input_text = sys.stdin.read()

    # Parse test cases
    test_cases = parse_test_cases(input_text)

    if not test_cases:
        die("No test cases found in input.\n"
            "Expected format:\n"
            "  Input: var1 = val1, var2 = val2\n"
            "  Output: result")

    print(f"Cases:     {len(test_cases)}", file=sys.stderr)

    # Generate
    content = generate_test(package, func_name, params, return_type,
                            test_cases)

    if args.dry_run:
        print(content)
    else:
        test_file.write_text(content)
        print(f"\nGenerated: {test_file}", file=sys.stderr)


if __name__ == "__main__":
    main()