package goal

import (
	"context"

	"github.com/nathancamolez-dev/go-orbit/internal/validator"
)

type CompleteGoalReq struct {
	GoalID string `json:"goalID"`
}

func (req CompleteGoalReq) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(
		validator.NotBlank(req.GoalID),
		"goalID",
		"Missing goalId for checking and completing",
	)

	return eval
}
