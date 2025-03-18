package jsonutils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nathancamolez-dev/go-orbit/internal/validator"
)

func EncodeJson[T any](w http.ResponseWriter, r *http.Request, statusCode int, data T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		return fmt.Errorf("error enconding json: %w", err)
	}
	return nil
}

func DecodeValidJson[T validator.Validator](r *http.Request) (T, map[string]string, error) {
	var data T
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, nil, fmt.Errorf("!!Failed to decode json: %w", err)
	}
	if problems := data.Valid(r.Context()); len(problems) > 0 {
		return data, problems, fmt.Errorf("!! Invalid %T: %d problems", data, len(problems))
	}

	return data, nil, nil
}
