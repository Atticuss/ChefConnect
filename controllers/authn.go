package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

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

// ConfigureMiddleware generates
func ConfigureMiddleware(controllerCtx *ControllerCtx) (*jwt.GinJWTMiddleware, error) {
	secretKey, err := generateRandomBytes(100)
	if err != nil {
		return &jwt.GinJWTMiddleware{}, err
	}

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "chefconnect",
		Key:         secretKey,
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "uid",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			fmt.Println("inside payloadfunc")
			if v, ok := data.(models.JwtUser); ok {
				// this logic is for converting the JwtUser struct to a map[string]interface{}
				// https://stackoverflow.com/a/42849112/13203635
				var claims jwt.MapClaims
				jsonbody, _ := json.Marshal(v)
				json.Unmarshal(jsonbody, &claims)

				return claims
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)

			// convert map[string]interface{} back into a JwtUser struct
			jwtUser := models.JwtUser{}
			jsonbody, _ := json.Marshal(claims)
			json.Unmarshal(jsonbody, &jwtUser)

			return jwtUser
		},
		Authenticator: controllerCtx.Login,
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true //push authorization off to the services layer
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{"error": message})
		},
		TokenLookup:   "header: Authorization, cookie: jwt",
		TokenHeadName: "Token",
		TimeFunc:      time.Now,
	})

	return authMiddleware, err
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
