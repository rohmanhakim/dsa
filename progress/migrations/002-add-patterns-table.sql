-- Migration 002: Add patterns table + session_patterns junction table
-- Date: 2026-07-08
-- Purpose: Support multi-pattern tracking and per-pattern mastery.

-- 1. Create patterns table (source of truth for pattern definitions)
CREATE TABLE IF NOT EXISTS patterns (
    name TEXT PRIMARY KEY,          -- canonical name from patterns-registry.md
    description TEXT,               -- what it does, when to use
    category TEXT                   -- e.g., "Two Pointers", "Binary Search", "Graph"
);

-- 2. Populate patterns table from registry
INSERT INTO patterns (name, description, category) VALUES
('Two Pointers', 'General two-pointer technique on arrays/strings', 'Two Pointers'),
('Two Pointers (Backward Merge)', 'Fill from end to avoid O(n) shifting', 'Two Pointers'),
('Two Pointers (Read/Write)', 'Separate read and write pointers', 'Two Pointers'),
('Sliding Window', 'Expand right, shrink left to maintain window property', 'Two Pointers'),
('Binary Search', 'Standard search on sorted data or monotonic predicate', 'Binary Search'),
('Binary Search on Answer', 'Binary search on answer space + feasibility check', 'Binary Search'),
('Greedy (Intervals)', 'Interval scheduling, activity selection, merge intervals', 'Greedy'),
('1D Dynamic Programming', 'Single-dimensional DP table', 'Dynamic Programming'),
('2D Dynamic Programming', 'Two-dimensional DP table', 'Dynamic Programming'),
('BFS', 'Breadth-first search, connected components', 'Graph'),
('DFS', 'Depth-first search, backtracking', 'Graph'),
('Topological Sort', 'DAG ordering', 'Graph'),
('Hash Map', 'Frequency counting, lookup tables, grouping', 'Other'),
('Monotonic Stack', 'Next greater/smaller element', 'Other'),
('Trie', 'Prefix tree', 'Other'),
('Digit Manipulation', 'Extracting/processing individual digits', 'Other'),
('Two-Pass Propagation', 'Left→right + right→left passes', 'Other'),
('Prefix Sum', 'Cumulative sums for range queries', 'Other');

-- 3. Create session_patterns junction table
CREATE TABLE IF NOT EXISTS session_patterns (
    session_id INTEGER NOT NULL,
    pattern_name TEXT NOT NULL,
    role TEXT NOT NULL DEFAULT 'primary',  -- 'primary' or 'secondary'
    PRIMARY KEY (session_id, pattern_name),
    FOREIGN KEY (session_id) REFERENCES sessions(id),
    FOREIGN KEY (pattern_name) REFERENCES patterns(name)
);

-- 4. Migrate existing sessions.pattern data into session_patterns
--    All existing patterns are treated as 'primary'
INSERT INTO session_patterns (session_id, pattern_name, role)
SELECT id, pattern, 'primary'
FROM sessions
WHERE pattern IS NOT NULL;

-- 5. Add indexes for common queries
CREATE INDEX IF NOT EXISTS idx_session_patterns_session_id ON session_patterns(session_id);
CREATE INDEX IF NOT EXISTS idx_session_patterns_pattern_name ON session_patterns(pattern_name);
CREATE INDEX IF NOT EXISTS idx_patterns_category ON patterns(category);

-- 6. Drop the pattern column from sessions
DROP INDEX IF EXISTS idx_sessions_pattern;
ALTER TABLE sessions DROP COLUMN pattern;
