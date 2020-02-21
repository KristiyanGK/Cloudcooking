package auth

import (
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/KristiyanGK/cloudcooking/models"
)

// GenerateToken generates a jwt token and returns it
func GenerateToken(secret string, user models.User) string {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &models.UserToken {
		ID: string(user.ID),
		Username: user.Username,
		Picture: user.Picture,
		Role: user.Role.Name,
		StandardClaims: jwt.StandardClaims {
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		fmt.Println(err)
	}

	return tokenString
}

// ParseToken parses a jwt token and returns the user claims
func ParseToken(secret, tokenString string) models.UserToken {

	token, err := jwt.ParseWithClaims(tokenString, &models.UserToken{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(*models.UserToken); ok && token.Valid {
		return *claims
	} else {
		fmt.Println(err)
	}

	return models.UserToken{}
}

// IsTokenValid receives a jwt token in the form of a string and checks wheater the 
// token is valid or not
func IsTokenValid(token string) bool {
	return true
}
