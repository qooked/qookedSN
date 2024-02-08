package main

import (
	"log"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt"
)

func checkPassword(password string) bool {
	regularExpression := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`)
	return regularExpression.MatchString(password)
}

func checkEmail(email string) bool {
	regularExpression := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return regularExpression.MatchString(email)
}

func GenerateAccessToken(username string, uuid string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Header["kid"] = "signin_1"
	tokenString, err := token.SignedString([]byte(CFG.JWTkey))
	if err != nil {
		log.Panicln("Error generating access token: ", err)
	}
	return tokenString
}

func GenerateTokenPair(username string, uuid string) (string, string, error) {

	tokenString := GenerateAccessToken(username, uuid)

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshToken.Header["kid"] = "signin_2"

	expirationTimeRefreshToken := time.Now().Add(24 * time.Hour).Unix()

	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = uuid
	rtClaims["exp"] = expirationTimeRefreshToken

	refreshTokenString, err := refreshToken.SignedString([]byte(CFG.JWTkey))
	if err != nil {
		log.Panicln("Error generating refresh token: ", err)
		return "", "", err
	}

	return tokenString, refreshTokenString, nil
}
