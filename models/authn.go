package models

// AuthnRequest is a struct that represents a single auth request. It is used exclusively
// for interaction with clients.
type AuthnRequest struct {
	Username string `json:"username,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}

// AuthnResponse is a struct that represents a single auth response.
type AuthnResponse struct {
	Token  string `json:"token,omitempty"`
	Code   int    `json:"code,omitempty"`
	Expire string `json:"expire,omitempty"`
}
