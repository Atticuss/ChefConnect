package integration_rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/atticuss/chefconnect/models"
	"github.com/stretchr/testify/assert"
)

func TestSuccesfulLogin(t *testing.T) {
	assert := assert.New(t)

	loginBody, _ := json.Marshal(map[string]string{
		"username": "jay.sea",
		"password": "Password1",
	})

	w := performRequest("POST", "/auth/login", "", loginBody)
	assert.Equal(http.StatusOK, w.Code, "Valid login test failed -- unexpected HTTP response code")

	var resp models.AuthnResponse
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		panic(err)
	}

	assert.NotEmptyf(resp, "Valid login test failed -- empty response body")
}

func TestFailedLogin(t *testing.T) {
	assert := assert.New(t)

	loginBody, _ := json.Marshal(map[string]string{
		"username": "jay.sea",
		"password": "1",
	})

	w := performRequest("POST", "/auth/login", "", loginBody)
	assert.Equal(http.StatusUnauthorized, w.Code, "Invalid login test failed -- unexpected HTTP response code")

	var resp models.AuthnResponse
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		panic(err)
	}

	assert.Emptyf(resp, "Valid login test failed -- non-empty response body")
}

func TestRefreshToken(t *testing.T) {
	assert := assert.New(t)

	loginBody, _ := json.Marshal(map[string]string{
		"username": "jay.sea",
		"password": "Password1",
	})

	var resp models.AuthnResponse
	w := performRequest("POST", "/auth/login", "", loginBody)
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		panic(err)
	}

	w = performRequest("GET", "/auth/refresh-token", resp.Token, nil)
	assert.Equal(http.StatusOK, w.Code, "Refresh token test failed -- unexpected HTTP response code")

	fmt.Println("resp body: ", w.Body.String())

	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		panic(err)
	}

	assert.NotEmptyf(resp, "Refresh token test failed -- empty response body")
}
