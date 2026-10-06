# DSA

Personal learning archive for data structures & algorithms, built around an LLM-guided practice loop and a structured progress tracker.

**What this is:**
- A record of problems solved — independently or with LLM guidance — and the journey around each one.
- A pattern-level mastery tracker (SQLite-backed) that detects stagnation and regression.
- A working set of mentor instructions (`AGENTS.md`) that turn every solving session into a coaching session.

**What this is not:**
- A solutions reference. Many solutions here are suboptimal — sometimes deliberately kept that way as a snapshot of how I solved the problem at that time, sometimes because I didn't know better yet. Read them as a learning journal, not as reference implementations.

Currently focused on **LeetCode** (a few early Codeforces / AtCoder problems remain as artifacts). Solutions are written in **Go**.

## Repository Layout

```
problems/
  lc/{easy,medium,hard}/lc<id>/   # LeetCode, one directory per problem
  cf/d800/                        # Codeforces (early problems)
  ac/abs/                         # AtCoder Beginner Contest (early problems)
progress/
  progress.db                     # SQLite tracker (gitignored — local only)
  schema.sql                      # DB schema: problems, patterns, sessions, session_patterns
  migrations/                     # Incremental schema changes
  patterns/                       # Per-pattern performance notes
  patterns-registry.md            # Canonical pattern names (source of truth for logging)
  queries/                        # Reusable SQL: mastery streaks, regression checks, recency
  overview.md                     # Living summary: mastered patterns, weaknesses, priorities
  practice-plan.md                # Current focus + problem queue
  weaknesses.md                   # Detailed weakness log
notes/patterns/                   # Concept notes per pattern (Obsidian-friendly markdown)
scripts/                          # Helpers: random-problem.py, hhmmss_to_seconds.py, ...
AGENTS.md                         # Mentor instructions (mode discipline, severity model, workflow)
```

### Problem Directory Convention

```
problems/lc/medium/lc72/
  edit-distance-72.go        # solution
  edit-distance-72_test.go   # Go tests (local verification)
  edit-distance-72.md        # problem writeup / official editorial (where present)
```

`main.go` at the repo root is a scratchpad for running a single problem's entry point.

## The Tracking System

Every session is logged into `progress.db` with two axes: **how I solved it** and **how well**.

**Status:** `independent` (no help) or `guided` (needed hints/LLM help).

**Severity:**

| Severity | Meaning |
|----------|---------|
| 0 | Pass — right pattern, within time |
| 1 | Slow pass — right pattern, over time |
| 2 | Wrong pattern — pursued a bad approach before recovering |
| 3 | Hint-needed — couldn't identify the approach without help |

**Difficulty time limits** (for severity 0): LC easy → 15 min · LC medium → 30 min · LC hard → 60 min.

**Sessions can carry multiple patterns** via a `session_patterns` junction table (primary = the pattern deliberately practiced, secondary = supporting). Each pattern is judged on its own merits — a bad binary-search attempt doesn't retroactively taint a clean DP session.

### Mastery & Regression

- A pattern is promoted to **mastered** only after **5 consecutive sessions with severity ≤ 1**, spread across **≥ 3 distinct days**, each within its difficulty's time limit.
- **Regression** is flagged on a single severity-3 session on a mastered pattern, or on 2+ severity ≥ 2 sessions in the last 3 *properly sampled* (distinct-day, non-contaminated) sessions.
- After a regression, re-certification has a lighter bar: 3 consecutive sessions across 2 days.

All SQL for these checks lives in `progress/queries/`.

## Mentor Workflow

Solving sessions run under `AGENTS.md`, which enforces:

- **Mode discipline** — a strict *plan* (read-only) / *act* (execute) alternation. No writes, including DB inserts and logging, until act mode is confirmed.
- **Pre-session checks** — look up prior attempts in the DB and current mastery context before judging a solution.
- **Trace before advising** — verify any proposed approach against a small example before presenting it.
- **Session logging** — every session is recorded with approach notes, key learning, weaknesses, and complexity analysis, with markdown-formatted notes for later review.

## Running

```sh
go test ./...                     # run all problem tests
go test ./problems/lc/easy/lc125/ # a single problem
python3 scripts/random-problem.py # pull a practice problem from the queue
```

`progress/progress.db` is gitignored and intentionally local — copy `schema.sql` (and apply migrations in order) to bootstrap a fresh tracker.

## Snapshot

| | |
|---|---|
| Problems archived | ~303 (295 LC + 5 CF + 3 AC) |
| Sessions logged | 138 (83 independent, 55 guided) |
| Patterns tracked | 31 |
| Patterns mastered | Binary Search, Stack, Greedy (Intervals), Hash Map |

*Counts grow over time; see `progress/overview.md` for the current state.*
