package v1

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/atticuss/chefconnect/models"
	"github.com/atticuss/chefconnect/services"
	"github.com/jinzhu/copier"
)

const randValues = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type jwtClaims struct {
	ID       string `json:"uid"`
	Name     string `json:"name"`
	Username string `json:"username"`

	Roles []*models.Role `json:"roles,omitempty"`

	jwt.StandardClaims
}

func compareHash(password, hash string) bool {
	if hash == "" {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateJwt(secretKey string, tokenExpiryPeriod int, user *models.User) (string, error) {
	claimDetails := jwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + int64(tokenExpiryPeriod),
			IssuedAt:  time.Now().Unix(),
			Audience:  "Authentication",
		},
	}
	copier.Copy(&claimDetails, user)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimDetails)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// source: https://gist.github.com/dopey/c69559607800d2f2f90b1b1ed4e550fb#gistcomment-3527095
func generateRefreshToken(tokenLength int) (string, error) {
	ret := make([]byte, tokenLength)
	for i := 0; i < tokenLength; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(randValues))))
		if err != nil {
			return "", err
		}
		ret[i] = randValues[num.Int64()]
	}

	return string(ret), nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// ValidateCredentials handles the business logic when a client passes in authn creds
func (s *v1Service) GenerateJwtTokens(userReq *models.User) (*models.User, *services.ServiceError) {
	user, err := s.Repository.GetUserByUsername(userReq.Username)
	if err != nil {
		// a simple attempt to stop timing attacks should someone discover a reliable way
		// to trigger an error on invalid usernames against the dgraph repo
		compareHash("a", "$2a$14$zR/r6hmGbPk1mh1G8fsvJOE/iKfhosK5YjVoiA51zgKmDnp6lETja")
		return user, &services.ServiceError{Error: err, ErrorCode: services.Unhandled}
	}

	if !compareHash(userReq.Password, user.Password) {
		return user, &services.ServiceError{
			Error:     errors.New("invalid credentials provided"),
			ErrorCode: services.NotAuthorized,
		}
	}

	jwtToken, _ := generateJwt(s.SecretKey, s.TokenExpiry, user)
	refreshToken, _ := generateRefreshToken(s.RefreshTokenLength)

	user.RefreshTokenIssuedAt = time.Now().Unix()
	user.RefreshToken = refreshToken
	user.AuthToken = jwtToken

	// save the refresh token to the data store
	s.Repository.UpdateUser(user)
	return user, &nilErr
}

func (s *v1Service) ExchangeRefreshToken(refreshToken string) (*models.User, *services.ServiceError) {
	user, err := s.Repository.GetUserByRefreshToken(refreshToken)
	if err != nil {
		return user, &services.ServiceError{Error: err}
	}

	// this is a dumb way to validate whether or not a user was returned. should
	// implement custom "NotFound" error within the repo layer and handle
	// here. example because i'm going to forget about this pattern:
	// 	switch t := err.(type) {
	//		default:
	//			fmt.Println("not a model missing error")
	//		case *ModelMissingError:
	//			fmt.Println("ModelMissingError", t)
	//	}
	if user.Username != "" {
		jwtToken, _ := generateJwt(s.SecretKey, s.TokenExpiry, user)
		user.AuthToken = jwtToken
	}

	return user, &nilErr
}

func (s *v1Service) DeserializeJwt(jwtToken string) (*models.User, *services.ServiceError) {
	user := &models.User{}

	token, err := jwt.ParseWithClaims(jwtToken, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return user, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(s.SecretKey), nil
	})

	if err != nil {
		fmt.Printf("err1: %+v\n", err)
		return user, &services.ServiceError{
			Error:     err,
			ErrorCode: services.NotAuthorized,
		}
	}

	claims, ok := token.Claims.(*jwtClaims)
	if ok && token.Valid {
		copier.Copy(user, claims)
	} else {
		fmt.Printf("ok: %+v\n", ok)
		fmt.Printf("token: %+v\n", token)
		return user, &services.ServiceError{
			Error:     errors.New("invalid JWT token provided"),
			ErrorCode: services.NotAuthorized,
		}
	}

	return user, &nilErr
}
