#!/usr/bin/env python3
"""Convert seconds duration to HH:mm:ss string."""

import sys

USAGE = "Usage: seconds_to_hhmmss.py <seconds>\n\nExample: seconds_to_hhmmss.py 4530"


def seconds_to_hhmmss(seconds: int) -> str:
    if seconds < 0:
        raise ValueError(f"expected non-negative integer, got {seconds}")

    h = seconds // 3600
    m = (seconds % 3600) // 60
    s = seconds % 60

    return f"{h:02d}:{m:02d}:{s:02d}"


def main() -> None:
    if len(sys.argv) != 2 or sys.argv[1] in ("-h", "--help"):
        print(USAGE)
        sys.exit(0 if sys.argv[1:] in (["-h"], ["--help"]) else 1)

    try:
        seconds = int(sys.argv[1])
    except ValueError:
        print(f"Error: not an integer: {sys.argv[1]}\n", file=sys.stderr)
        print(USAGE, file=sys.stderr)
        sys.exit(1)

    try:
        result = seconds_to_hhmmss(seconds)
    except ValueError as e:
        print(f"Error: {e}\n", file=sys.stderr)
        print(USAGE, file=sys.stderr)
        sys.exit(1)

    print(result)


if __name__ == "__main__":
    main()