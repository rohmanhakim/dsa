-- DSA Progress Tracking Database Schema
-- SQLite database for tracking problem-solving sessions

-- Problems table: one row per unique problem
CREATE TABLE IF NOT EXISTS problems (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    problem_id TEXT UNIQUE NOT NULL,   -- e.g., "lc72", "lc1268", "cf231a"
    problem_name TEXT NOT NULL,        -- e.g., "Edit Distance"
    difficulty TEXT,                   -- "easy", "medium", "hard"
    platform TEXT                      -- "lc", "cf", "ac"
);

-- Patterns table: one row per unique pattern
CREATE TABLE IF NOT EXISTS patterns (
    name TEXT PRIMARY KEY,             -- canonical name from patterns-registry.md
    description TEXT,                  -- what it does, when to use
    category TEXT                      -- e.g., "Two Pointers", "Binary Search", "Graph"
);

-- Sessions table: one row per attempt/session
CREATE TABLE IF NOT EXISTS sessions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    problem_id TEXT NOT NULL,          -- FK → problems.problem_id
    date TEXT NOT NULL,                -- YYYY-MM-DD
    status TEXT NOT NULL,              -- "independent" or "guided"
    time_to_accepted INTEGER,          -- seconds, 0 if guided
    severity INTEGER,                  -- 0=pass, 1=slow pass, 2=wrong pattern, 3=hint-needed. For guided sessions, always 3.
    approach_notes TEXT,               -- approach notes (observed by mentor)
    key_learning TEXT,                 -- key learning from the session
    weaknesses TEXT,                   -- weaknesses identified
    complexity_time TEXT,              -- time complexity
    complexity_space TEXT,             -- space complexity
    created_at TEXT DEFAULT (datetime('now')),
    FOREIGN KEY (problem_id) REFERENCES problems(problem_id)
);

-- Session-Patterns junction table: which patterns were used in each session
CREATE TABLE IF NOT EXISTS session_patterns (
    session_id INTEGER NOT NULL,
    pattern_name TEXT NOT NULL,
    role TEXT NOT NULL DEFAULT 'primary',  -- 'primary' or 'secondary'
    PRIMARY KEY (session_id, pattern_name),
    FOREIGN KEY (session_id) REFERENCES sessions(id),
    FOREIGN KEY (pattern_name) REFERENCES patterns(name)
);

-- Indexes for common queries
CREATE INDEX IF NOT EXISTS idx_sessions_problem_id ON sessions(problem_id);
CREATE INDEX IF NOT EXISTS idx_sessions_date ON sessions(date);
CREATE INDEX IF NOT EXISTS idx_sessions_status ON sessions(status);
CREATE INDEX IF NOT EXISTS idx_session_patterns_session_id ON session_patterns(session_id);
CREATE INDEX IF NOT EXISTS idx_session_patterns_pattern_name ON session_patterns(pattern_name);
CREATE INDEX IF NOT EXISTS idx_patterns_category ON patterns(category);
CREATE INDEX IF NOT EXISTS idx_problems_platform ON problems(platform);
CREATE INDEX IF NOT EXISTS idx_problems_difficulty ON problems(difficulty);
