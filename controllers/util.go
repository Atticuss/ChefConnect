package controllers

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"

	"github.com/dgraph-io/dgo/v2"
	"github.com/go-playground/validator/v10"
)

// ControllerCtx holds all contextual data for the controller package. Fields are configured
// by the `main` package when the app is initialized. Currently, all it holds is the dgraph
// client object, passed to `models` functions as a param. Could probably move this struct
// to the `models` package to avoid passing via params, but more contextual data may be
// added here later.
type ControllerCtx struct {
	DgraphClient *dgo.Dgraph
	Validator    *validator.Validate
}

// Unfinshed (obviously). The goal of this function is to provide a general purpose conversion
// between ManyFoo{} and ManyFooResponse{} types. For example, in the controllers.GetAllUsers()
// function, the following must be performed:
//   cleanResp := models.ManyUsersResponse{}
//   for _, user := range resp.Users {
//   	cleanResp.Users = append(cleanResp.Users, models.UserResponse(user))
//   }
// This logic could be replaced with a general purpose utility function leveraging reflection
// to dynamically pull and convert the fields between the two structs.

//func convertParentStruct(src interface{}, dst interface{}) {
//	srcVal := reflect.Indirect(reflect.ValueOf(src))
//	srcRootField := srcVal.Type().Field(0).Name
//	srcFieldValues := srcVal.FieldByName(srcRootField)
//}

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

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
