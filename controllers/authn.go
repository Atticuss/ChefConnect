package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/atticuss/chefconnect/models"
)

// body comment
// swagger:parameters login
type authnRequest struct {
	// in:body
	models.AuthnRequest
}

// swagger:response AuthnResponse
type authn struct {
	// in:body
	Body models.AuthnResponse
}

func (ctx *ControllerCtx) Login(w http.ResponseWriter, r *http.Request) {
	// swagger:route POST /login authn login
	// Authenticate against the app
	// responses:
	//   200: AuthnResponse

	var authReq models.AuthnRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&authReq); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if resp, sErr := ctx.Service.Login(authReq); sErr.Error != nil {
		respondWithServiceError(w, sErr)
	} else {
		respondWithJSON(w, http.StatusOK, resp)
	}
}
