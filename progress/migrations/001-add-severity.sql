-- Migration 001: Add severity column + backfill + index
-- Date: 2026-07-08
-- Purpose: Support weighted regression model by tracking failure severity per session.

-- 1. Add severity column (nullable for backfill compatibility)
ALTER TABLE sessions ADD COLUMN severity INTEGER;

-- 2. Backfill: guided sessions → severity 3 (hint-needed)
UPDATE sessions SET severity = 3 WHERE status = 'guided';

-- 3. Backfill: independent sessions within time limit → severity 0 (pass)
--    Time limits: easy ≤ 900s (15min), medium ≤ 1800s (30min), hard ≤ 3600s (1hr)
UPDATE sessions SET severity = 0
WHERE status = 'independent'
  AND severity IS NULL
  AND (
    (problem_id IN (SELECT problem_id FROM problems WHERE difficulty = 'easy') AND time_to_accepted <= 900)
    OR (problem_id IN (SELECT problem_id FROM problems WHERE difficulty = 'medium') AND time_to_accepted <= 1800)
    OR (problem_id IN (SELECT problem_id FROM problems WHERE difficulty = 'hard') AND time_to_accepted <= 3600)
  );

-- 4. Backfill: remaining independent sessions (over time limit) → severity 1 (slow pass)
UPDATE sessions SET severity = 1
WHERE status = 'independent'
  AND severity IS NULL;

-- 5. Add pattern index for LIKE queries used in mastery/regression checks
CREATE INDEX IF NOT EXISTS idx_sessions_pattern ON sessions(pattern);