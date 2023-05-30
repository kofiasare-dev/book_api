package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const ()

func Hash(plainText string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)

	return string(bytes), err
}

func Compare(digest, plaintext string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(digest), []byte(plaintext))

	return err == nil
}

func GenerateToken(userID string) (string, error) {

	expirationTime := time.Now().Add(1 * time.Hour)

	claims := jwt.MapClaims{
		"exp":        expirationTime.Unix(),
		"authorized": true,
		"user":       userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func VerifyToken(tokenStr string) (token *jwt.Token, err error) {
	token, err = jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	return token, err

}
