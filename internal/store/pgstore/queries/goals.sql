-- name: CreateGoal :one
INSERT INTO goals(title, desiredWeeklyFrequency) values ($1, $2) RETURNING id;

-- name: CompleteGoal :exec
INSERT INTO goalsCompletions(goalId) VALUES ($1);

-- name: GetGoalsCreatedThisWeekAndPending :many
WITH GoalsCreatedThisWeek AS (
    SELECT 
        id, 
        title, 
        desiredWeeklyFrequency, 
        createdAt 
    FROM 
        goals 
    WHERE 
        createdAt >= date_trunc('week', CURRENT_DATE) 
        AND createdAt < date_trunc('week', CURRENT_DATE) + interval '1 week'
),
CompletionCounts AS (
    SELECT 
        g.id, 
        g.title, 
        g.desiredWeeklyFrequency, 
        g.createdAt, 
        COALESCE(COUNT(gc.goalId), 0) AS completion_count
    FROM 
        GoalsCreatedThisWeek g
    LEFT JOIN 
        goalsCompletions gc ON g.id = gc.goalId
    GROUP BY 
        g.id, g.title, g.desiredWeeklyFrequency, g.createdAt
)
SELECT 
    cc.id, 
    cc.title, 
    cc.desiredWeeklyFrequency, 
    cc.createdAt,
    cc.completion_count
FROM 
    CompletionCounts cc
WHERE 
    cc.completion_count < cc.desiredWeeklyFrequency;


-- name: GetWeekSummary :many
SELECT 
    g.title,
    g.desiredWeeklyFrequency,
    COUNT(gc.id) AS completion_count
FROM 
    goals g
LEFT JOIN 
    goalsCompletions gc ON g.id = gc.goalId
WHERE 
    g.createdAt >= date_trunc('week', current_date) 
    AND g.createdAt < date_trunc('week', current_date) + interval '1 week'
GROUP BY 
    g.id, g.title, g.desiredWeeklyFrequency
ORDER BY 
    g.title;
