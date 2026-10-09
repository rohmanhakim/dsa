# Difficulty Levels & Time Limits

Single source of truth for platform difficulty leveling and the severity-0 time
goal. Read this file **on demand** — AGENTS.md / .clinerules point here instead
of duplicating the tables, and `scripts/levels.py` parses the two tables below
so the test harnesses (gengen / hrgen / acgen), `scripts/acdiff.py` and
`scripts/random-problem.py` can resolve a goal programmatically.

Platforms covered: `lc` (LeetCode), `hr` (HackerRank), `cf` (Codeforces),
`ac` (AtCoder).

## Severity-0 time goals (per tier)

| tier | max_seconds |
|------|-------------|
| easy | 900 |
| medium | 1800 |
| hard | 3600 |

The goal for a problem is `max_seconds` of its tier (15 / 30 / 60 min): correct
pattern + solved within goal = severity 0; correct pattern but over goal =
severity 1 (severity scale itself lives in AGENTS.md).

## How the key is derived (per platform)

| platform | key source | directory example |
|----------|------------|-------------------|
| lc | tier directory; also stored in `problems.difficulty` | `problems/lc/medium/lc72` → medium |
| hr | tier directory | `problems/hr/easy/missing-numbers` → easy |
| cf | problem rating from the directory name (`d<rating>`) | `problems/cf/d800/cf231a` → 800 |
| ac | AtCoder Problems difficulty number (kenkoooo `problem-models.json`), **not** the problem letter; tier directory after placement | `abc081_a` → -418 → `problems/ac/easy/abc081_a` |

## Numeric bands → tier (cf, ac)

| platform | min | max | tier |
|----------|-----|-----|------|
| cf |  | 1199 | easy |
| cf | 1200 | 1499 | medium |
| cf | 1500 |  | hard |
| ac |  | 1199 | easy |
| ac | 1200 | 1499 | medium |
| ac | 1500 |  | hard |

Bands are inclusive; an empty bound means unbounded. Rows are ordered by
platform then min — the first matching band wins. `ac` easy deliberately has no
lower bound: AtCoder Problems difficulty is negative for trivial tasks.

## Notes / caveats

- the goal depends only on the **tier**, so the same tier means the same goal on
  every platform (lc / hr / ac encode the tier directly in the directory name;
  cf resolves its rating through the band table above).
- cf problem ratings run ~800–3500 (steps of 50/100); values below 800 are rare
  → easy.
- ac difficulty is an Elo estimate ("the rating at which 50% of solvers solve
  it") on the same scale as AtCoder user colors (grey <400 … red 2800+);
  negatives are normal (`abc081_a` ≈ -418).
- ac letter heuristic — **only** when no difficulty estimate exists (mark the
  tier provisional and revisit once an estimate appears; the number always wins
  over the letter):
  - `abc`: A–D easy · E medium · F and later hard
  - `arc`: A–B medium · C and later hard
  - `agc`: A medium · B and later hard
  - practice / tutorial / unknown contest: easy
- db convention: `problems.difficulty` stores the **tier** (`easy|medium|hard`)
  for all platforms; `problems.problem_id` for ac = the canonical task id.
  the raw cf rating lives in the directory name (`cf/d800`); the raw ac
  difficulty number is re-derivable with `scripts/acdiff.py`.
- ac naming (lc-style): directory = canonical task id (`abc086_a` — identical
  to the atcoder.jp URL slug and the kenkoooo key); package = id minus
  underscore (`abc086a`); file = `<kebab-title>-<task-id>.go` (+ `_test.go`,
  + `.md` for writeups), mirroring `lc/medium/lc72/edit-distance-72.go`.

## Lookup

```sh
python3 scripts/levels.py easy                 # tier → seconds
python3 scripts/levels.py cf 800               # platform + numeric key → tier + seconds
python3 scripts/levels.py ac -418              # negative difficulty is valid
python3 scripts/levels.py dir problems/lc/medium/lc72    # problem dir → tier + seconds
```

## Changelog

- **2026-10-09**: initial — extracted from AGENTS.md / .clinerules, which had
  drifted apart on CF bands (bullets said 800-900 / 1000-1100 / 1200+, tables
  said 800-1100 / 1200-1400 / 1500+, leaving 1101–1199 unmapped). Unified-model
  table won: 800–1199 / 1200–1499 / 1500+. Added `hr` (was missing everywhere)
  and `ac`.
