#!/usr/bin/env python3
"""
Migration script to restructure problem files into per-problem packages.

Current structure:
  problems/{platform}/{difficulty}/problem-name-{id}.go

New structure:
  problems/{platform}/{difficulty}/{platform}{id}/problem-name-{id}.go

Each problem gets its own directory (package), solving Go's package naming collision issue.
"""

import os
import re
import shutil
from pathlib import Path


def parse_filename(filename: str):
    """
    Parse a filename like 'problem-name-123.go' or 'problem-name-123_test.go' or 'problem-name-123.md'

    Returns (problem_name, problem_id) or None if the file doesn't match.
    """
    # Remove extension to get base name
    # Handle .go, _test.go, .md
    if filename.endswith('_test.go'):
        base = filename[:-len('_test.go')]
        ext = '_test.go'
    elif filename.endswith('.go'):
        base = filename[:-len('.go')]
        ext = '.go'
    elif filename.endswith('.md'):
        base = filename[:-len('.md')]
        ext = '.md'
    else:
        return None

    # Split on last '-' to get ID and problem name
    # The ID is the last segment after the last '-'
    last_dash = base.rfind('-')
    if last_dash == -1:
        return None

    problem_name = base[:last_dash]
    problem_id = base[last_dash + 1:]

    return (problem_name, problem_id, ext)


def get_package_name(platform: str, problem_id: str) -> str:
    """Generate a valid Go package name from platform + id."""
    return f"{platform}{problem_id}"


def update_package_declaration(filepath: Path, new_package_name: str):
    """Update the package declaration in a .go file to match the new directory name."""
    content = filepath.read_text()
    # Replace the package declaration line
    # Match 'package <old_name>' at the start of a line
    new_content = re.sub(
        r'^package\s+\w+',
        f'package {new_package_name}',
        content,
        count=1,
        flags=re.MULTILINE
    )
    if new_content != content:
        filepath.write_text(new_content)
        print(f"  Updated package declaration in {filepath.name} -> package {new_package_name}")


def migrate(root_dir: str):
    """Main migration function."""
    problems_dir = Path(root_dir) / "problems"

    if not problems_dir.exists():
        print(f"Error: {problems_dir} does not exist")
        return

    # Collect all files to migrate
    files_to_migrate = []
    for dirpath, dirnames, filenames in os.walk(problems_dir):
        for filename in filenames:
            filepath = Path(dirpath) / filename
            result = parse_filename(filename)
            if result is None:
                continue

            # Extract platform, difficulty from the directory path
            # e.g., problems/lc/easy/ -> platform=lc, difficulty=easy
            rel_path = filepath.relative_to(problems_dir)
            parts = rel_path.parts  # e.g., ('lc', 'easy', 'can-place-flowers-605.go')

            if len(parts) < 3:
                # This shouldn't happen for properly structured files
                print(f"  Skipping {filepath} (unexpected directory structure)")
                continue

            platform = parts[0]
            difficulty = parts[1]
            problem_name, problem_id, ext = result

            files_to_migrate.append({
                'filepath': filepath,
                'platform': platform,
                'difficulty': difficulty,
                'problem_name': problem_name,
                'problem_id': problem_id,
                'ext': ext,
                'package_name': get_package_name(platform, problem_id),
            })

    print(f"Found {len(files_to_migrate)} files to migrate.\n")

    # Group files by (platform, difficulty, problem_id) to handle test files
    groups = {}
    for file_info in files_to_migrate:
        key = (file_info['platform'], file_info['difficulty'], file_info['problem_id'])
        if key not in groups:
            groups[key] = []
        groups[key].append(file_info)

    print(f"Found {len(groups)} unique problems to migrate.\n")

    # Perform migration
    for (platform, difficulty, problem_id), files in groups.items():
        package_name = get_package_name(platform, problem_id)
        new_dir = problems_dir / platform / difficulty / package_name

        print(f"Migrating {platform}/{difficulty}/{package_name}/ ({len(files)} file(s))")

        # Create the new directory
        new_dir.mkdir(parents=True, exist_ok=True)

        for file_info in files:
            old_path = file_info['filepath']
            new_path = new_dir / old_path.name

            # Move the file
            shutil.move(str(old_path), str(new_path))
            print(f"  Moved {old_path.name} -> {new_path}")

            # Update package declaration for .go and _test.go files
            if file_info['ext'] in ('.go', '_test.go'):
                update_package_declaration(new_path, package_name)

        print()

    print("Migration complete!")
    print("\nNote: main.go was intentionally not modified (it's always changing).")
    print("You may need to update the import path in main.go manually.")


if __name__ == "__main__":
    root = os.path.dirname(os.path.abspath(__file__))
    migrate(root)