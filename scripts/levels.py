#!/usr/bin/env python3
"""levels.py - resolve difficulty tiers and severity-0 time goals.

progress/levels.md is the single source of truth for platform leveling
(lc/hr/cf/ac). This module parses its two tables so the test harnesses
(gengen/hrgen/acgen), acdiff.py and random-problem.py can resolve a goal
on demand instead of hard-coding numbers.

CLI:
    python3 levels.py easy                  # tier -> seconds
    python3 levels.py cf 800                # platform + numeric key -> tier + seconds
    python3 levels.py ac -418               # negative difficulty is valid
    python3 levels.py dir problems/lc/medium/lc72   # problem dir -> tier + seconds

Importable:
    from levels import lookup_dir, tier_seconds, band_tier, print_goal, Level
"""

import os
import re
import sys
from pathlib import Path

LEVELS_FILE = Path(__file__).resolve().parent.parent / "progress" / "levels.md"
PLATFORMS = ("lc", "hr", "cf", "ac")

USAGE = """usage:
  levels.py <tier>                 e.g. levels.py easy
  levels.py <platform> <key>       e.g. levels.py cf 800
                                   e.g. levels.py ac -418
                                   e.g. levels.py lc hard
  levels.py dir <problem-path>     e.g. levels.py dir problems/ac/easy/abc086_a"""

_CACHE = None


class Level:
    """A resolved goal: platform + key + tier + seconds."""

    def __init__(self, platform, key, tier, max_seconds):
        self.platform = platform
        self.key = key          # tier str (lc/hr/ac) or rating int (cf)
        self.tier = tier
        self.max_seconds = max_seconds

    @property
    def goal(self):
        return f"{self.max_seconds}s ({self.max_seconds // 60} min)"

    def __str__(self):
        if self.key == self.tier:
            head = f"{self.platform} ({self.key})"
        else:
            head = f"{self.platform} {self.key}"
        return f"{head} -> tier {self.tier} -> {self.goal}"


def _die(msg):
    """Parser/consistency errors are fatal: levels.md must stay machine-readable."""
    print(f"levels.py: {msg}", file=sys.stderr)
    sys.exit(1)


def _cells(line):
    """Split a markdown table row into stripped cells; None if not a table row."""
    line = line.strip()
    if not line.startswith("|") or not line.endswith("|"):
        return None
    return [c.strip() for c in line[1:-1].split("|")]


def _parse_section(text, title_prefix, expected_headers):
    """Parse the first table under the '## <title prefix>' heading.

    Returns data rows as list[list[str]]. Fails loudly when the section or the
    header row does not match expected_headers exactly.
    """
    lines = text.splitlines()
    start = None
    for i, line in enumerate(lines):
        if line.startswith("## ") and title_prefix in line:
            start = i + 1
            break
    if start is None:
        _die(f"section '## {title_prefix}...' not found in {LEVELS_FILE}")

    run = []
    for line in lines[start:]:
        if line.lstrip().startswith("|"):
            run.append(line)
        elif run:
            break
    if not run:
        _die(f"no table found under '## {title_prefix}' in {LEVELS_FILE}")

    header = _cells(run[0])
    if header != expected_headers:
        _die(f"unexpected header under '## {title_prefix}': "
             f"{header!r} != {expected_headers!r}")

    rows = []
    for line in run[1:]:
        cells = _cells(line)
        if cells and all(re.fullmatch(r":?-+:?", c) for c in cells):
            continue  # separator row
        if cells is None or len(cells) != len(expected_headers):
            _die(f"malformed table row in {LEVELS_FILE}: {line!r}")
        rows.append(cells)
    if not rows:
        _die(f"table under '## {title_prefix}' has no data rows")
    return rows


def _load():
    global _CACHE
    if _CACHE is not None:
        return _CACHE
    if not LEVELS_FILE.is_file():
        _die(f"{LEVELS_FILE} not found")
    text = LEVELS_FILE.read_text(encoding="utf-8")

    tiers = {}
    for tier, secs in _parse_section(
            text, "Severity-0 time goals", ["tier", "max_seconds"]):
        if not re.fullmatch(r"[a-z]+", tier):
            _die(f"bad tier name: {tier!r}")
        if tier in tiers:
            _die(f"duplicate tier row: {tier!r}")
        try:
            tiers[tier] = int(secs)
        except ValueError:
            _die(f"max_seconds for tier {tier!r} is not an integer: {secs!r}")
        if tiers[tier] <= 0:
            _die(f"max_seconds for tier {tier!r} must be positive")

    bands = []
    for plat, lo, hi, tier in _parse_section(
            text, "Numeric bands", ["platform", "min", "max", "tier"]):
        if plat not in PLATFORMS:
            _die(f"unknown platform in band table: {plat!r}")
        if tier not in tiers:
            _die(f"band tier {tier!r} not defined in the time-goals table")
        bounds = []
        for label, cell in (("min", lo), ("max", hi)):
            if cell == "":
                bounds.append(None)
            else:
                try:
                    bounds.append(int(cell))
                except ValueError:
                    _die(f"band {label} for {plat} is not an integer: {cell!r}")
        lo_i, hi_i = bounds
        if lo_i is not None and hi_i is not None and lo_i > hi_i:
            _die(f"band for {plat} has min > max: {lo_i} > {hi_i}")
        bands.append((plat, lo_i, hi_i, tier))

    _CACHE = {"tiers": tiers, "bands": bands}
    return _CACHE


def tiers():
    """{tier: max_seconds}"""
    return dict(_load()["tiers"])


def tier_seconds(tier):
    """Seconds goal for a tier; ValueError if unknown."""
    known = _load()["tiers"]
    if tier not in known:
        raise ValueError(
            f"unknown tier {tier!r} (known: {', '.join(sorted(known))})")
    return known[tier]


def band_tier(platform, value):
    """Resolve a numeric key (cf rating / ac difficulty) to a tier."""
    value = int(value)
    for plat, lo, hi, tier in _load()["bands"]:
        if plat != platform:
            continue
        if (lo is None or value >= lo) and (hi is None or value <= hi):
            return tier
    raise ValueError(
        f"no band in progress/levels.md covers {platform} value {value}")


def _platform_from_path(path):
    """Extract (platform, difficulty-component) from a problem path."""
    parts = Path(path).parts
    idx = None
    for i, part in enumerate(parts):
        if part == "problems" and i + 1 < len(parts):
            idx = i + 1
            break
    if idx is None:
        idx = next((i for i, p in enumerate(parts) if p in PLATFORMS), None)
    if idx is None or parts[idx] not in PLATFORMS:
        raise ValueError(f"no platform component in path {path!r}")
    platform = parts[idx]
    if idx + 1 >= len(parts):
        raise ValueError(
            f"path {path!r} has no difficulty component after {platform!r}")
    return platform, parts[idx + 1]


def lookup_dir(path, *fallback_paths):
    """Resolve a problem path to a Level.

    lc/hr/ac paths carry the tier directly (problems/<plat>/<tier>/...);
    cf paths carry the rating (problems/cf/d800/...). fallback_paths (e.g.
    os.getcwd()) are tried when path itself has no platform component — the
    harnesses pass a bare solution filename while cwd is the problem dir.
    """
    failures = []
    for candidate in (path,) + tuple(fallback_paths):
        try:
            platform, comp = _platform_from_path(candidate)
        except ValueError as exc:
            failures.append(str(exc))
            continue
        try:
            if platform in ("lc", "hr", "ac"):
                seconds = tier_seconds(comp)
                return Level(platform, comp, comp, seconds)
            m = re.fullmatch(r"d(\d+)", comp)
            if not m:
                raise ValueError(
                    f"cf directory {comp!r} is not of the form d<rating>")
            rating = int(m.group(1))
            tier = band_tier(platform, rating)
            return Level(platform, rating, tier, tier_seconds(tier))
        except ValueError as exc:
            failures.append(str(exc))
            continue
    raise ValueError(
        f"cannot derive platform+tier from {path!r}"
        + (f" (tried {list(fallback_paths)})" if fallback_paths else "")
        + f" ({'; '.join(failures)})")


def print_goal(path, stream=None):
    """Print the severity-0 goal for a problem path; never raises."""
    stream = stream or sys.stderr
    try:
        level = lookup_dir(path, os.getcwd())
    except ValueError as exc:
        print(f"Goal:      unresolved ({exc})", file=stream)
        return None
    print(f"Goal:      tier {level.tier} -> {level.goal} "
          f"[see progress/levels.md]", file=stream)
    return level


def main(argv):
    try:
        if len(argv) == 2 and argv[0] == "dir":
            print(lookup_dir(argv[1], os.getcwd()))
        elif len(argv) == 1:
            tier = argv[0]
            secs = tier_seconds(tier)
            print(f"{tier} -> {secs}s ({secs // 60} min)")
        elif len(argv) == 2:
            platform, key = argv
            if platform not in PLATFORMS:
                raise ValueError(
                    f"unknown platform {platform!r} "
                    f"(known: {', '.join(PLATFORMS)})")
            if key in _load()["tiers"]:
                secs = tier_seconds(key)
                print(f"{platform} {key} -> tier {key} -> "
                      f"{secs}s ({secs // 60} min)")
            else:
                tier = band_tier(platform, key)
                secs = tier_seconds(tier)
                print(f"{platform} {key} -> tier {tier} -> "
                      f"{secs}s ({secs // 60} min)")
        else:
            print(USAGE, file=sys.stderr)
            return 2
    except ValueError as exc:
        print(f"levels.py: {exc}", file=sys.stderr)
        return 1
    return 0


if __name__ == "__main__":
    sys.exit(main(sys.argv[1:]))


