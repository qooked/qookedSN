package main

import (
	"log"
	"strings"
	"time"
	"unicode"

	"github.com/golang-jwt/jwt"
)

func checkPassword(password string) bool {
	hasLower := false
	hasUpper := false
	hasDigit := false
	hasSpecial := false

	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasLower && hasUpper && hasDigit && hasSpecial && len(password) >= 8
}

func checkEmail(email string) bool {
	atIndex := strings.Index(email, "@")
	dotIndex := strings.LastIndex(email, ".")
	if atIndex < 0 || dotIndex < 0 {
		return false
	}
	if atIndex >= dotIndex {
		return false
	}
	if len(email)-dotIndex <= 1 {
		return false
	}
	return true
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
