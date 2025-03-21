package api

import (
	"errors"
	"net/http"

	"go.uber.org/zap"

	"github.com/nathancamolez-dev/go-orbit/internal/functions"
	"github.com/nathancamolez-dev/go-orbit/internal/jsonutils"
	goal "github.com/nathancamolez-dev/go-orbit/internal/usecases/goal"
)

func (api *API) handleCreateGoal(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJson[goal.CreateGoalReq](r)
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, map[string]any{
			"error": err.Error(),
		})
		if problems != nil {
			jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, map[string]any{
				"problems": problems,
			})
		}
		return
	}

	if err := api.GoalFunctions.CreateGoal(r.Context(), data.Title, data.DesiredWeeklyFrequency); err != nil {
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "Internal Server Error",
		})
		zap.Error(err)
		return
	}

	jsonutils.EncodeJson(w, r, http.StatusCreated, map[string]any{
		"Message": "Goal create successfully",
	})
	return

}

func (api *API) handleCompleteGoal(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJson[goal.CompleteGoalReq](r)
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, map[string]any{
			"error": err.Error(),
		})
		if problems != nil {
			jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, map[string]any{
				"problems": problems,
			})
		}
		return
	}

	if err := api.GoalFunctions.CompleteGoal(r.Context(), data.GoalID); err != nil {
		if errors.Is(err, functions.ErrNoGoal) {
			jsonutils.EncodeJson(w, r, http.StatusNotFound, map[string]any{
				"error": err.Error(),
			})
			return

		}
		zap.Error(err)
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "Internal Server Error",
		})
		return
	}

	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{
		"Message": "Goal completed successfully",
	})
	return
}

func (api *API) handleGetWeekSummary(w http.ResponseWriter, r *http.Request) {
	weekSummary, err := api.GoalFunctions.GetWeekSummary(r.Context())
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"Error": "Internal server error",
		})
		return
	}
	jsonutils.EncodeJson(w, r, http.StatusOK, weekSummary)
	return
}
