#!/usr/bin/env python3
"""Convert HH:mm:ss duration string to seconds."""

import sys

USAGE = "Usage: hhmmss_to_seconds.py HH:mm:ss\n\nExample: hhmmss_to_seconds.py 01:15:30"


def hhmmss_to_seconds(duration: str) -> int:
    parts = duration.split(":")
    if len(parts) != 3:
        raise ValueError(f"expected 3 colon-separated parts, got {len(parts)}")

    h, m, s = (int(p) for p in parts)

    if h < 0 or m < 0 or m >= 60 or s < 0 or s >= 60:
        raise ValueError(f"out-of-range values: {duration}")

    return h * 3600 + m * 60 + s


def main() -> None:
    if len(sys.argv) != 2 or sys.argv[1] in ("-h", "--help"):
        print(USAGE)
        sys.exit(0 if sys.argv[1:] in (["-h"], ["--help"]) else 1)

    try:
        seconds = hhmmss_to_seconds(sys.argv[1])
    except ValueError as e:
        print(f"Error: {e}\n", file=sys.stderr)
        print(USAGE, file=sys.stderr)
        sys.exit(1)

    print(seconds)


if __name__ == "__main__":
    main()