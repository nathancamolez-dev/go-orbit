package functions

import (
	"context"
	"fmt"

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
