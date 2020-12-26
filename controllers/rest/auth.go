package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	"github.com/atticuss/chefconnect/models"
	"github.com/atticuss/chefconnect/services"
)

// body comment
// swagger:parameters login
type swaggerAuthnRequest struct {
	// in:body
	Body authnRequest
}

// swagger:response AuthnResponse
type swaggerAuthnResponse struct {
	// in:body
	Body authnResponse
}

type authnRequest struct {
	Username     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
}

type authnResponse struct {
	AuthToken    string `json:"authToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
}

// both middleware funcs based on code provided here:
// https://sosedoff.com/2014/12/21/gin-middleware.html
func (restCtrl *restController) jwtDeserializationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken := c.Request.Header.Get(restCtrl.Config.AuthTokenHeaderName)
		if jwtToken == "" {
			c.Set("callingUser", &models.User{})
			c.Next()
			return
		}

		// TODO: unexpected error, log before return
		callingUser, sErr := restCtrl.Service.DeserializeJwt(jwtToken)
		if sErr.Error != nil && sErr.ErrorCode != services.NotAuthorized {
			respondWithServiceError(c, sErr)
			c.Abort()
		}

		c.Set("callingUser", callingUser)
		c.Next()
	}
}

func (restCtrl *restController) LoginHandler(c *gin.Context) {
	// swagger:route POST /login authn login
	// Authenticate against the app
	// responses:
	//   200: AuthnResponse

	var authnReq authnRequest
	if err := c.ShouldBindJSON(&authnReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{}
	copier.Copy(user, &authnReq)

	user, sErr := restCtrl.Service.GenerateJwtTokens(user)
	if sErr.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect username or password"})
		if sErr.ErrorCode == services.NotAuthorized {
			return
		}

		// unexpected error received. should log something.
		// TODO: when logging is implemented

		return
	}

	resp := &authnResponse{
		AuthToken:    user.AuthToken,
		RefreshToken: user.RefreshToken,
	}

	c.JSON(http.StatusOK, resp)
}

func (restCtrl *restController) RefreshHandler(c *gin.Context) {
	// swagger:route POST /refresh-token authn refresh
	// Authenticate against the app
	// responses:
	//   200: AuthnResponse

	var authnReq authnRequest
	if err := c.ShouldBindJSON(&authnReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, sErr := restCtrl.Service.ExchangeRefreshToken(authnReq.RefreshToken)
	if sErr.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})

		// unexpected error received. should log something.
		// TODO: when logging is implemented
		if sErr.ErrorCode != services.NotAuthorized {
		}

		return
	}

	resp := &authnResponse{
		AuthToken:    user.AuthToken,
		RefreshToken: user.RefreshToken,
	}

	c.JSON(http.StatusOK, resp)
}
