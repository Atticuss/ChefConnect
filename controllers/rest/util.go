package rest

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"

	"github.com/atticuss/chefconnect/models"
	"github.com/atticuss/chefconnect/services"
)

var statusCodeMap = [...]int{
	services.Unhandled:      http.StatusBadRequest,
	services.NotImplemented: http.StatusInternalServerError,
	services.NotFound:       http.StatusNotFound,
	services.NotAuthorized:  http.StatusUnauthorized,
	services.ResourceInUse:  http.StatusUnauthorized,
}

// Unfortunately, we can't directly cast from the map[string]struct object pulled
// from `jwt.ExtractClaims()` into a models.User struct. This is because there's
// no json marshalling tags associated with structs within the `models` package.
// Instead, we have to unmarshall into a jwtClaims struct and then copy it over
// into a models.User struct. Not as clean as I'd like, but the `models` package
// should remain agnostic about things like json marshalling.
func getUserFromContext(c *gin.Context) (*models.User, error) {
	user := models.User{}
	claimsStruct := jwtClaims{}

	claims := jwt.ExtractClaims(c)
	jsonbody, err := json.Marshal(claims)
	if err != nil {
		return &user, err
	}

	json.Unmarshal(jsonbody, &claimsStruct)
	copier.Copy(&user, &claimsStruct)

	return &user, nil
}

func resolveFieldToTag(s interface{}, field string) string {
	t := reflect.TypeOf(s)
	f, _ := t.FieldByName(field)
	v, _ := f.Tag.Lookup("json")
	return v
}

func respondWithValidationError(c *gin.Context, err error, model interface{}) {
	validationErrors := err.(validator.ValidationErrors)
	missingFields := make([]string, len(validationErrors))
	for idx, err := range validationErrors {
		missingFields[idx] = resolveFieldToTag(model, err.Field())
	}

	errorMsg := "Required fields are missing: " + strings.Join(missingFields, ", ")
	c.Error(err)
	respondWithError(c, http.StatusBadRequest, errorMsg)
}

func respondWithServiceError(c *gin.Context, sErr *services.ServiceError) {
	c.Error(sErr.Error)
	respondWithError(c, statusCodeMap[sErr.ErrorCode], sErr.Error.Error())
}

func respondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"error": message})
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

// shamelessly stolen from: https://gist.github.com/dopey/c69559607800d2f2f90b1b1ed4e550fb
func assertAvailablePRNG() error {
	// Assert that a cryptographically secure PRNG is available.
	// Panic otherwise.
	buf := make([]byte, 1)

	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return err
	}

	return nil
}

func generateRandomBytes(n int) ([]byte, error) {
	if err := assertAvailablePRNG(); err != nil {
		return nil, fmt.Errorf("crypto/rand is unavailable: Read() failed with %#v", err)
	}

	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}
