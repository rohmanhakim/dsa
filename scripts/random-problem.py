#!/usr/bin/env python3
"""Pick a random problem directory from a given category/difficulty path."""

import os
import sys
import random

SCRIPT_DIR = os.path.dirname(os.path.abspath(__file__))
PROBLEMS_DIR = os.path.join(SCRIPT_DIR, os.pardir, "problems")
EXCLUDED_FILE = os.path.join(SCRIPT_DIR, "excluded-problems.txt")


def main():
    if len(sys.argv) < 2:
        print(f"Usage: {sys.argv[0]} <dir>")
        print(f"Example: {sys.argv[0]} lc/easy")
        sys.exit(1)

    rel_path = sys.argv[1].rstrip("/")
    target_dir = os.path.join(PROBLEMS_DIR, rel_path)

    if not os.path.isdir(target_dir):
        print(f"Error: '{target_dir}' is not a valid directory.")
        sys.exit(1)

    # Collect subdirectories
    dirs = [d for d in os.listdir(target_dir) if os.path.isdir(os.path.join(target_dir, d))]

    if not dirs:
        print(f"No problem directories found in '{target_dir}'.")
        sys.exit(1)

    # Load excluded problems
    excluded = set()
    if os.path.isfile(EXCLUDED_FILE):
        with open(EXCLUDED_FILE) as f:
            for line in f:
                line = line.strip()
                if line:
                    excluded.add(line)

    # Filter out excluded problems
    candidates = [d for d in dirs if f"{rel_path}/{d}" not in excluded]

    if not candidates:
        print(f"All problems in '{rel_path}' are excluded.")
        sys.exit(1)

    chosen = random.choice(candidates)
    print(os.path.join("problems", rel_path, chosen))


if __name__ == "__main__":
    main()