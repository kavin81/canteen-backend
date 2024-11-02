package auth

import (
	"canteen-backend/config"

	"time"

	"github.com/golang-jwt/jwt/v5"

	"golang.org/x/crypto/bcrypt"
)

var jwtKey string = config.Config.Auth.JWT.Secret

type Claims struct {
	Username string `json:"username"`
	// ExpiresAt int64 `json:"exp"`
	jwt.RegisteredClaims
}

func GenerateToken(username string) (string, error) {

	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.Config.Auth.JWT.Expiry)),
			Issuer:    "PESU",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Audience:  jwt.ClaimStrings{"pes-canteen"},
			// TODO: if logged in add SRN
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString([]byte(jwtKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// TODO: verify token

func ValidatePassword(password string) bool {
	return len(password) >= 8 && len(password) <= 64
}

func HashPassword(password string) string {

	var salt string = config.Config.Auth.Salt
	var saltedPassword = password + salt

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func ComparePasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {
		return false
	}

	return true
}
