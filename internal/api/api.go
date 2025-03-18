package api

import (
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"

	"github.com/nathancamolez-dev/go-orbit/internal/functions"
)

type API struct {
	Router        *chi.Mux
	GoalFunctions functions.GoalFunctions
	Sessions      *scs.SessionManager
}
