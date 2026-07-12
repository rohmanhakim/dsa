-- mastery-streak-check.sql
-- Counts consecutive qualifying sessions for a pattern, starting from the
-- most recent session and working backwards. Stops counting at the first
-- session with severity > 1 (i.e., severity 2 or 3 breaks the streak).
--
-- WHY THIS EXISTS:
-- A naive "grab last 5 sessions and check if all severity <= 1" approach
-- has a bug: it would count qualifying sessions from BEFORE a break.
-- Example: [sev1, sev1, sev3, sev1, sev1] (newest -> oldest)
--   - Naive approach: qualifying = 4 (ignores the sev3, counts the other 4)
--   - Correct answer: qualifying = 0 (sev3 at position 3 breaks the streak,
--     so the two sev1s before it are not counted either)
--
-- HOW IT WORKS:
-- 1. all_sessions: ranks primary-pattern sessions for the pattern by date DESC using ROW_NUMBER().
--    Only sessions where the pattern is the primary role are included — secondary patterns
--    in multi-pattern problems do not affect the mastery streak.
-- 2. streak CTE: starts with the most recent session (rn=1), then recursively
--    adds the next older session (rn+1) ONLY if the current session had severity <= 1.
--    The moment severity > 1 is encountered, recursion stops.
-- 3. Outer SELECT: counts qualifying sessions (severity <= 1 only) and distinct days.
--
-- NOTE: SQLite does not allow ORDER BY + LIMIT inside a recursive CTE anchor,
-- so we use ROW_NUMBER() in a separate CTE to assign ranks first.
--
-- USAGE:
-- Replace 'YOUR_PATTERN' with the canonical pattern name from patterns-registry.md.

WITH RECURSIVE all_sessions AS (
    SELECT s.id, s.severity, s.date,
           ROW_NUMBER() OVER (ORDER BY s.date DESC) as rn
    FROM sessions s
    JOIN session_patterns sp ON s.id = sp.session_id
    WHERE sp.pattern_name = 'YOUR_PATTERN'
      AND sp.role = 'primary'
),
streak AS (
    -- anchor: most recent session (rn = 1)
    SELECT id, severity, date, rn
    FROM all_sessions WHERE rn = 1
    UNION ALL
    -- recurse: next older row, ONLY if current session qualified
    SELECT a.id, a.severity, a.date, a.rn
    FROM all_sessions a
    JOIN streak st ON a.rn = st.rn + 1
    WHERE st.severity <= 1
)
SELECT
    COUNT(*) as qualifying_sessions,
    COUNT(DISTINCT date) as distinct_days,
    CASE
        WHEN COUNT(*) >= 5 AND COUNT(DISTINCT date) >= 3
        THEN 'READY FOR MASTERY'
        ELSE 'IN PROGRESS'
    END as mastery_status
FROM streak
WHERE severity <= 1;