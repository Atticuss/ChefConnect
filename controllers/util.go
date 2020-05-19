package controllers

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/atticuss/chefconnect/services"
)

// ControllerCtx holds all contextual data for the controller package. Fields are configured
// by the `main` package when the app is initialized. Currently, all it holds is the dgraph
// client object, passed to `models` functions as a param. Could probably move this struct
// to the `models` package to avoid passing via params, but more contextual data may be
// added here later.
type ControllerCtx struct {
	Service services.Service
}

var statusCodeMap = [...]int{
	services.Unhandled: http.StatusBadRequest,
	services.NotFound:  http.StatusNotFound,
}

func resolveFieldToTag(s interface{}, field string) string {
	t := reflect.TypeOf(s)
	f, _ := t.FieldByName(field)
	v, _ := f.Tag.Lookup("json")
	return v
}

func respondWithValidationError(w http.ResponseWriter, err error, model interface{}) {
	validationErrors := err.(validator.ValidationErrors)
	missingFields := make([]string, len(validationErrors))
	for idx, err := range validationErrors {
		missingFields[idx] = resolveFieldToTag(model, err.Field())
	}

	errorMsg := "Required fields are missing: " + strings.Join(missingFields, ", ")
	respondWithError(w, http.StatusBadRequest, errorMsg)
}

func respondWithServiceError(w http.ResponseWriter, sErr services.ServiceError) {
	respondWithError(w, statusCodeMap[sErr.ErrorCode], sErr.Error.Error())
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {

	response := []byte("{}")
	if payload != nil {
		response, _ = json.Marshal(payload)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
