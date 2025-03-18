package api

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/nathancamolez-dev/go-orbit/internal/jsonutils"
	goal "github.com/nathancamolez-dev/go-orbit/internal/usecases"
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

	if err := api.GoalFunctions.CreateGoal(r.Context(), data.Title, data.DesiredWeeklyFequency); err != nil {
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "Internal Server Error",
		})
		zap.Error(err)
		return
	}
}
