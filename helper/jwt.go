package helper

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWT_KEY = Getenv("JWT_KEY", "secret for jwt key")

func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{}

	claims["authorized"] = true
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // token valid for 24 hours

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(JWT_KEY))
	return result, err
}

func ExtractToken(bearerToken string) (jwt.MapClaims, error) {
	hmacSecret := []byte(JWT_KEY)
	token, err := jwt.Parse(bearerToken, func(t *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func ExtractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(header, " ")

	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorized header")
	}

	return jwtToken[1], nil
}
