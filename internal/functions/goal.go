package functions

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	"github.com/nathancamolez-dev/go-orbit/internal/store/pgstore"
)

type GoalFunctions struct {
	pool    *pgxpool.Pool
	queries *pgstore.Queries
}

func NewGoalFunctions(pool *pgxpool.Pool) GoalFunctions {
	return GoalFunctions{
		pool:    pool,
		queries: pgstore.New(pool),
	}
}

func (gf *GoalFunctions) CreateGoal(
	ctx context.Context,
	title string, desiredWeeklyfrequency int) error {

	args := pgstore.CreateGoalParams{
		Title:                  title,
		Desiredweeklyfrequency: int32(desiredWeeklyfrequency),
	}

	if _, err := gf.queries.CreateGoal(ctx, args); err != nil {
		var pgerr *pgconn.PgError
		zap.Error(err)
		return pgerr
	}
	fmt.Println("Goal created successfully")
	return nil

}

var ErrNoGoal = errors.New("This are not created this week, or is not pending")

func (gf *GoalFunctions) CompleteGoal(
	ctx context.Context,
	goalId string) error {
	pendingWeekGoals, err := gf.queries.GetGoalsCreatedThisWeekAndPending(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrNoGoal
		}
		zap.Error(err)
		return err
	}
	goalID, err := uuid.Parse(goalId)
	if err != nil {
		zap.Error(err)
		return err
	}

	for _, goal := range pendingWeekGoals {
		if goal.ID == goalID {
			if err := gf.queries.CompleteGoal(ctx, goalID); err != nil {
				zap.Error(err)
				return err
			}
			return nil
		}
	}

	return ErrNoGoal

}

func (gf *GoalFunctions) GetWeekSummary(
	ctx context.Context,
) (weekSummary []pgstore.GetWeekSummaryRow, err error) {
	weekSummary, err = gf.queries.GetWeekSummary(ctx)
	if err != nil {
		return nil, err
	}

	return weekSummary, nil
}

func (gf *GoalFunctions) GetWeekPendingGoals(
	ctx context.Context,
) (weekPendingGoals []pgstore.GetWeekSummaryRow, err error) {
	weekPendingGoals, err = gf.queries.GetWeekSummary(ctx)
	if err != nil {
		return nil, err
	}

	return weekPendingGoals, nil
}
