package utils

import (
	"time"

	"github.com/Djancyp/go-rest/pkg/config"
	"github.com/golang-jwt/jwt/v4"
)

// TODO: get key from .env
var jwtKey = []byte(config.GetConfig("GOLANG_JWT_KEY"))
var expireMinute = 1

type Claims struct {
	Role  interface{} `json:"role"`
	Email string      `json:"username"`
	jwt.StandardClaims
}

func CreateJwtString() string {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, _ := token.SignedString(jwtKey)
	return tokenString
}

func CreateJwtWithClaim(email string, roles interface{}) (string, time.Time, error) {
	var expirationTime = time.Now().Add(1 * time.Minute)
	claims := &Claims{
		Role:  roles,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, expirationTime, err
}

func RefreshJwt(oldToken string) (string, time.Time, error) {
	claims := &Claims{}
	expirationTime := time.Now().Add(1 * time.Minute)
	tkn, err := jwt.ParseWithClaims(oldToken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", expirationTime, err
		}
		return "", expirationTime, err
	}
	if !tkn.Valid {
		return "", expirationTime, err
	}
	claims.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, expirationTime, err
}

func ValidateJwt(token string) (*jwt.Token, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return tkn, err
}
func GetJwtClaims(token string) (*Claims, error) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return claims, err
}
