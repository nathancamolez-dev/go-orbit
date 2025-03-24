// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: goals.sql

package pgstore

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const completeGoal = `-- name: CompleteGoal :exec
INSERT INTO goalsCompletions(goalId) VALUES ($1)
`

func (q *Queries) CompleteGoal(ctx context.Context, goalid uuid.UUID) error {
	_, err := q.db.Exec(ctx, completeGoal, goalid)
	return err
}

const createGoal = `-- name: CreateGoal :one
INSERT INTO goals(title, desiredWeeklyFrequency) values ($1, $2) RETURNING id
`

type CreateGoalParams struct {
	Title                  string `json:"title"`
	Desiredweeklyfrequency int32  `json:"desiredweeklyfrequency"`
}

func (q *Queries) CreateGoal(ctx context.Context, arg CreateGoalParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, createGoal, arg.Title, arg.Desiredweeklyfrequency)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const getGoalsCreatedThisWeekAndPending = `-- name: GetGoalsCreatedThisWeekAndPending :many
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
    cc.completion_count < cc.desiredWeeklyFrequency
`

type GetGoalsCreatedThisWeekAndPendingRow struct {
	ID                     uuid.UUID   `json:"id"`
	Title                  string      `json:"title"`
	Desiredweeklyfrequency int32       `json:"desiredweeklyfrequency"`
	Createdat              time.Time   `json:"createdat"`
	CompletionCount        interface{} `json:"completion_count"`
}

func (q *Queries) GetGoalsCreatedThisWeekAndPending(ctx context.Context) ([]GetGoalsCreatedThisWeekAndPendingRow, error) {
	rows, err := q.db.Query(ctx, getGoalsCreatedThisWeekAndPending)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetGoalsCreatedThisWeekAndPendingRow
	for rows.Next() {
		var i GetGoalsCreatedThisWeekAndPendingRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Desiredweeklyfrequency,
			&i.Createdat,
			&i.CompletionCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getWeekSummary = `-- name: GetWeekSummary :many
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
    g.title
`

type GetWeekSummaryRow struct {
	Title                  string `json:"title"`
	Desiredweeklyfrequency int32  `json:"desiredweeklyfrequency"`
	CompletionCount        int64  `json:"completion_count"`
}

func (q *Queries) GetWeekSummary(ctx context.Context) ([]GetWeekSummaryRow, error) {
	rows, err := q.db.Query(ctx, getWeekSummary)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetWeekSummaryRow
	for rows.Next() {
		var i GetWeekSummaryRow
		if err := rows.Scan(&i.Title, &i.Desiredweeklyfrequency, &i.CompletionCount); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
