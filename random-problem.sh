#!/bin/bash

if [ -z "$1" ]; then
  echo "Usage: $0 <dir>"
  echo "Example: $0 lc/easy"
  exit 1
fi

dir="problems/$1"

if [ ! -d "$dir" ]; then
  echo "Error: '$dir' is not a valid directory."
  exit 1
fi

file=$(find "$dir" -maxdepth 1 -type f | shuf -n 1)

if [ -z "$file" ]; then
  echo "No files found in '$dir'."
  exit 1
fi

echo "$file"