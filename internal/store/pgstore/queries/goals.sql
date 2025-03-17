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
PendingGoals AS (
    SELECT 
        g.id, 
        g.title, 
        g.desiredWeeklyFrequency, 
        g.createdAt, 
        COALESCE(COUNT(gc.goalsId), 0) AS completion_count
    FROM 
        goals g
    LEFT JOIN 
        goalsCompletions gc ON g.id = gc.goalsId
    GROUP BY 
        g.id, g.title, g.desiredWeeklyFrequency, g.createdAt
    HAVING 
        COALESCE(COUNT(gc.goalsId), 0) < g.desiredWeeklyFrequency
)
SELECT 
    gtw.id, 
    gtw.title, 
    gtw.desiredWeeklyFrequency, 
    gtw.createdAt
FROM 
    GoalsCreatedThisWeek gtw
JOIN 
    PendingGoals pg ON gtw.id = pg.id;
