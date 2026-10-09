#!/usr/bin/env python3
"""acdiff.py - Resolve AtCoder problem difficulty -> tier (placement helper).

Fetches kenkoooo's problem-models.json (the estimated-difficulty dataset
behind AtCoder Problems' symbol), looks up one problem, and reports the
numeric-band tier from progress/levels.md plus the suggested directory.

Usage:
    python3 scripts/acdiff.py abc086_a        # kenkoooo id / repo dir name
    python3 scripts/acdiff.py abc086a         # compact form also accepted
    python3 scripts/acdiff.py practice_a

Matching ignores case, underscores and a legacy 'ac' prefix, so several
spellings of the same problem id all resolve to it.
A full path may be given; only the last component is used.

When the dataset has no estimate, the letter heuristic from progress/levels.md
is used and the result is marked provisional.

Network: one fetch per run (no caching) - kenkoooo asks clients to keep
request rates low.
"""

import json
import re
import sys
import urllib.error
import urllib.request
from pathlib import Path

import levels

MODELS_URL = "https://kenkoooo.com/atcoder/resources/problem-models.json"
AC_ROOT = Path(__file__).resolve().parent.parent / "problems" / "ac"


def die(msg) -> None:
    print(f"Error: {msg}", file=sys.stderr)
    sys.exit(1)  # type: ignore[misc]


def canon(s):
    """Canonical matching form: lowercase, alphanumeric only."""
    return re.sub(r"[^a-z0-9]", "", s.lower())


def canonical_task_id(pid):
    """Canonical AtCoder task id: lowercase, underscore before the task letter.

    Compact or 'ac'-prefixed spellings -> 'abc086_a'; already-canonical ids
    pass through. Assumes a trailing task letter (atcoder.jp URL slug /
    kenkoooo key form).
    """
    s = pid[2:] if pid.lower().startswith("ac") else pid
    s = s.lower()
    if "_" in s:
        return s
    m = re.fullmatch(r"([a-z]+?\d*)_?([a-z])", s)
    if m:
        return f"{m.group(1)}_{m.group(2)}"
    return s


def letter_heuristic(pid):
    """Provisional (tier, why) from contest prefix + task letter.

    Mirrors the letter heuristic documented in progress/levels.md - used only
    when the dataset has no difficulty estimate.
    """
    s = pid.lower()
    if s.startswith("ac"):
        s = s[2:]
    m = re.match(r"([a-z]+?)(?:\d+)?_?([a-z]?)$", s)
    kind, letter = (m.group(1), m.group(2)) if m else ("", "")
    if kind == "abc" and letter:
        tier = ("easy" if letter in "abcd"
                else "medium" if letter == "e" else "hard")
    elif kind == "arc" and letter:
        tier = "medium" if letter in "ab" else "hard"
    elif kind == "agc" and letter:
        tier = "medium" if letter == "a" else "hard"
    else:
        tier = "easy"
    why = (f"letter heuristic ({kind or pid}"
           f"{letter.upper() if letter else ''}) -> {tier} (provisional)")
    return tier, why


def main(argv):
    if len(argv) != 1:
        die("usage: acdiff.py <problem-id>\n"
            "  e.g. acdiff.py abc086_a | abc086a | practice_a")
    pid = Path(argv[0]).name if "/" in argv[0] else argv[0]

    try:
        req = urllib.request.Request(
            MODELS_URL,
            headers={"User-Agent": "acdiff.py (personal DSA archive)"})
        with urllib.request.urlopen(req, timeout=30) as resp:
            models = json.load(resp)
    except (urllib.error.URLError, OSError, ValueError) as exc:
        die(f"fetch of {MODELS_URL} failed: {exc}\n"
            "no network? resolve the tier manually with the letter heuristic "
            "in progress/levels.md.")

    index = {canon(k): k for k in models}
    candidates = [pid] + ([pid[2:]] if pid.lower().startswith("ac") else [])
    hit = next((index[canon(c)] for c in candidates if canon(c) in index), None)

    # the kenkoooo key is already the canonical id; reconstruct it otherwise
    suggested = hit or canonical_task_id(pid)

    entry = models.get(hit) or {}
    difficulty = entry.get("difficulty")
    if hit and difficulty is not None:
        try:
            tier = levels.band_tier("ac", difficulty)
        except ValueError as exc:
            die(str(exc))
        source = (f"difficulty {difficulty} (kenkoooo: {hit}; "
                  f"experimental: {entry.get('is_experimental')})")
    else:
        tier, source = letter_heuristic(pid)
        source += ("; not found in problem-models.json" if not hit
                   else "; no difficulty estimate in problem-models.json")

    secs = levels.tier_seconds(tier)

    print(f"problem:  {pid}" + (f" (kenkoooo: {hit})" if hit else ""))
    print(f"source:   {source}")
    print(f"tier:     {tier} -> {secs}s ({secs // 60} min)  "
          f"[progress/levels.md]")
    print(f"dir:      problems/ac/{tier}/{suggested}")

    # cross-check an already-placed directory
    for tdir in ("easy", "medium", "hard"):
        base = AC_ROOT / tdir
        if not base.is_dir():
            continue
        for d in base.iterdir():
            if d.is_dir() and canon(d.name) == canon(suggested):
                status = ("ok" if tdir == tier
                          else f"MISMATCH (tier says {tier})")
                print(f"existing: problems/ac/{tdir}/{d.name} [{status}]")
    return 0


if __name__ == "__main__":
    sys.exit(main(sys.argv[1:]))
