package controllers

import (
	"log"
	"net/http"

	"github.com/atticuss/chefconnect/models"
	"github.com/gin-gonic/gin"
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

// ValidateJwt is the middleware responsible for ensuring a JWT, if present, is valid
func ValidateJwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing ValidateJwt()")
		next.ServeHTTP(w, r)
		log.Println("Executing ValidateJwt() again")
	})
}

// Login handles the POST /login req during authentication
func (ctx *ControllerCtx) Login(c *gin.Context) {
	// swagger:route POST /login authn login
	// Authenticate against the app
	// responses:
	//   200: AuthnResponse

	var authnReq models.AuthnRequest
	if err := c.ShouldBindJSON(&authnReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if resp, sErr := ctx.Service.Login(authnReq); sErr.Error != nil {
		respondWithServiceErrorGin(c, sErr)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}
