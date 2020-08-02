package rest

import (
	"encoding/json"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	"github.com/atticuss/chefconnect/models"
	"github.com/atticuss/chefconnect/services"
)

var hmacSecret []byte

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

type jwtClaims struct {
	ID       string `json:"uid"`
	Name     string `json:"name"`
	Username string `json:"username"`

	Roles []nestedRole `json:"roles,omitempty"`
}

func (restCtrl *restController) configureMiddleware() (*jwt.GinJWTMiddleware, error) {
	secretKey, err := generateRandomBytes(100)
	if err != nil {
		return &jwt.GinJWTMiddleware{}, err
	}

	missingTokenMsg := "token not found"

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:         "chefconnect",
		Key:           secretKey,
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		IdentityKey:   "uid",
		DisabledAbort: true,
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Token",
		TimeFunc:      time.Now,
		Authenticator: restCtrl.login,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(jwtClaims); ok {
				// this logic is for converting the jwtUser struct to a map[string]interface{}
				// https://stackoverflow.com/a/42849112/13203635

				var claims jwt.MapClaims
				jsonbody, err := json.Marshal(v)
				if err != nil {
					return claims
				}

				json.Unmarshal(jsonbody, &claims)
				return claims
			}

			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)

			// convert map[string]interface{} back into a jwtUser struct
			user := jwtClaims{}
			jsonbody, err := json.Marshal(claims)
			if err != nil {
				return user
			}

			json.Unmarshal(jsonbody, &user)
			return user
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true //push authorization off to the services layer
		},
		// standardize the error message returned if a token is not found, regardless of where it is searched for.
		// unfortunately, can't rely on the raw error interface as the HTTPStatusMessageFunc returns a string.
		HTTPStatusMessageFunc: func(err error, c *gin.Context) string {
			missingJwtSlice := []error{jwt.ErrEmptyAuthHeader, jwt.ErrEmptyQueryToken, jwt.ErrEmptyCookieToken}
			for _, e := range missingJwtSlice {
				if e == err {
					return missingTokenMsg
				}
			}
			return err.Error()
		},
		// using the standardized error message on missing token, allow the next middleware to execute if this
		// func is being called due to a missing token. this is because most resources are accessible regardless
		// if a user is authenticated, but some resources will return different data depending on authn/z status.
		Unauthorized: func(c *gin.Context, code int, message string) {
			if message == missingTokenMsg {
				c.Next()
				return
			}

			c.JSON(code, gin.H{"error": message})
		},
	})

	return authMiddleware, err
}

func (restCtrl *restController) login(c *gin.Context) (interface{}, error) {
	// swagger:route POST /login authn login
	// Authenticate against the app
	// responses:
	//   200: AuthnResponse

	var authnReq authnRequest
	if err := c.ShouldBindJSON(&authnReq); err != nil {
		return nil, jwt.ErrMissingLoginValues
	}

	user := &models.User{}
	copier.Copy(user, &authnReq)

	user, sErr := restCtrl.Service.ValidateCredentials(user)
	if sErr.Error != nil {
		if sErr.ErrorCode == services.NotAuthorized {
			return nil, jwt.ErrFailedAuthentication
		}
		return nil, sErr.Error
	}

	claimDetails := jwtClaims{}
	copier.Copy(&claimDetails, &user)

	return claimDetails, nil
}
