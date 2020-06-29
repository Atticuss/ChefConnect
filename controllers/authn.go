package controllers

import (
	"log"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"

	"github.com/atticuss/chefconnect/models"
	"github.com/atticuss/chefconnect/services"
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
func (ctx *ControllerCtx) Login(c *gin.Context) (interface{}, error) {
	// swagger:route POST /login authn login
	// Authenticate against the app
	// responses:
	//   200: AuthnResponse

	var authnReq models.AuthnRequest
	if err := c.ShouldBindJSON(&authnReq); err != nil {
		return nil, jwt.ErrMissingLoginValues
	}

	user, sErr := ctx.Service.Login(authnReq)
	if sErr.Error != nil {
		if sErr.ErrorCode == services.NotAuthorized {
			return nil, jwt.ErrFailedAuthentication
		}
		return nil, sErr.Error
	}

	return user, nil
}
