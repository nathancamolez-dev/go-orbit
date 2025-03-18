package goal

import (
	"context"

	"github.com/nathancamolez-dev/go-orbit/internal/validator"
)

type CreateGoalReq struct {
	Title                  string `json:"title"`
	DesiredWeeklyFrequency int    `json:"desiredWeeklyfrequency"`
}

func (req CreateGoalReq) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(validator.NotBlank(req.Title), "title", "You must provide a title")
	eval.CheckField(
		validator.NotBlankNumber(req.DesiredWeeklyFrequency),
		"desiredWeeklyFrequency",
		"You must provide a desired weekly frequency",
	)

	eval.CheckField(
		validator.MinChar(req.Title, 5) && validator.MaxChar(req.Title, 100),
		"title",
		"Title must be between 5 and 100 characters",
	)

	eval.CheckField(
		validator.MinValue(req.DesiredWeeklyFrequency, 1) &&
			validator.MaxValue(req.DesiredWeeklyFrequency, 7),
		"desiredWeeklyFrequency",
		"Desired weekly frequency must be between 1 and 7",
	)

	return eval
}
