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

-- Sessions table: one row per attempt/session
CREATE TABLE IF NOT EXISTS sessions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    problem_id TEXT NOT NULL,          -- FK → problems.problem_id
    date TEXT NOT NULL,                -- YYYY-MM-DD
    status TEXT NOT NULL,              -- "independent" or "guided"
    time_to_accepted INTEGER,          -- minutes, 0 if guided
    pattern TEXT,                      -- e.g., "2D Dynamic Programming"
    approach_notes TEXT,               -- approach notes (observed by mentor)
    key_learning TEXT,                 -- key learning from the session
    weaknesses TEXT,                   -- weaknesses identified
    complexity_time TEXT,              -- time complexity
    complexity_space TEXT,             -- space complexity
    created_at TEXT DEFAULT (datetime('now')),
    FOREIGN KEY (problem_id) REFERENCES problems(problem_id)
);

-- Indexes for common queries
CREATE INDEX IF NOT EXISTS idx_sessions_problem_id ON sessions(problem_id);
CREATE INDEX IF NOT EXISTS idx_sessions_date ON sessions(date);
CREATE INDEX IF NOT EXISTS idx_sessions_status ON sessions(status);
CREATE INDEX IF NOT EXISTS idx_problems_platform ON problems(platform);
CREATE INDEX IF NOT EXISTS idx_problems_difficulty ON problems(difficulty);