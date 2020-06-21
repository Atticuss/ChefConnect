package v1

import (
	"crypto/rand"
	"fmt"
	"io"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/copier"

	"github.com/atticuss/chefconnect/models"
	"github.com/atticuss/chefconnect/services"
)

var hmacSecret []byte

type jwtClaims struct {
	Roles []models.NestedRole `json:"roles"`
	jwt.StandardClaims
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func compareHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// shamelessly stolen from: https://gist.github.com/dopey/c69559607800d2f2f90b1b1ed4e550fb
func assertAvailablePRNG() error {
	// Assert that a cryptographically secure PRNG is available.
	// Panic otherwise.
	buf := make([]byte, 1)

	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return err
	}

	return nil
}

func generateRandomBytes(n int) ([]byte, error) {
	if err := assertAvailablePRNG(); err != nil {
		return nil, fmt.Errorf("crypto/rand is unavailable: Read() failed with %#v", err)
	}

	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func generateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	bytes, err := generateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes), nil
}

func generateJWT(user *models.User, authnResp *models.AuthnResponse) error {
	var err error
	if len(hmacSecret) == 0 {
		hmacSecret, err = generateRandomBytes(100)

		if err != nil {
			return err
		}
	}

	// migrate to a NestedRole for JSON serialization within the JWT, otherwise we end up with a "User"
	// field within each serialized Role struct
	apiUser := models.APIUser{}
	copier.Copy(&apiUser, &user)

	claims := jwtClaims{
		Roles: apiUser.Roles,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "chefconnect",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		return err
	}

	authnResp.JWT = tokenString
	return nil
}

// https://godoc.org/github.com/dgrijalva/jwt-go#example-Parse--Hmac
func parseJwt(tokenString string) (jwt.MapClaims, error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSecret, nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}

// GetAllCategories handles the business logic when a client requests all categories
func (s *v1Service) Login(authReq models.AuthnRequest) (models.AuthnResponse, services.ServiceError) {
	authnResp := models.AuthnResponse{}

	user, err := s.UserRepository.GetByUsername(authReq.Username)
	if err != nil {
		return authnResp, services.ServiceError{Error: err}
	}

	if compareHash(authReq.Password, user.Password) {
		if err := generateJWT(user, &authnResp); err != nil {
			return authnResp, services.ServiceError{Error: err}
		}
	}

	_, err = parseJwt(authnResp.JWT)
	if err != nil {
		fmt.Printf("error ocurred: %+v", err)
	}

	return authnResp, nilErr
}
