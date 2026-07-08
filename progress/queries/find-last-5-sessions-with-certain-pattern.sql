SELECT 
    s.date,
    s.problem_id,
    p.difficulty,
    s.status,
    s.severity,
    s.pattern,
    s.time_to_accepted
FROM sessions s
JOIN problems p ON s.problem_id = p.problem_id
WHERE p.platform = 'lc'
  AND (s.pattern LIKE '%Binary Search%')
ORDER BY s.date DESC
LIMIT 5;
