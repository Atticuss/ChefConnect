package integration_rest

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccesfulLogin(t *testing.T) {
	assert := assert.New(t)

	loginBody, _ := json.Marshal(map[string]string{
		"username": "jay.sea",
		"password": "Password1",
	})

	w := performRequest("POST", "/auth/login", nil, loginBody)
	assert.Equal(http.StatusOK, w.Code, "Login test failed")
}

func TestFailedLogin(t *testing.T) {
	assert := assert.New(t)

	loginBody, _ := json.Marshal(map[string]string{
		"username": "jay.sea",
		"password": "1",
	})

	w := performRequest("POST", "/auth/login", nil, loginBody)
	assert.Equal(http.StatusUnauthorized, w.Code, "Login test failed")
}
