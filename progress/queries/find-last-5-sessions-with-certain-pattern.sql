SELECT 
    s.date,
    s.problem_id,
    p.difficulty,
    s.status,
    s.severity,
    sp.pattern_name,
    s.time_to_accepted
FROM sessions s
JOIN problems p ON s.problem_id = p.problem_id
JOIN session_patterns sp ON s.id = sp.session_id
WHERE p.platform = 'lc'
  AND sp.pattern_name = 'Binary Search'
ORDER BY s.date DESC
LIMIT 5;
