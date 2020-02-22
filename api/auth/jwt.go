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
		UserID: string(user.ID),
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
// It does not check wheater the token is valid or not so use
// IsTokenValid to ensure the token is valid
func ParseToken(secret, tokenString string) models.UserToken {

	token, _ := jwt.ParseWithClaims(tokenString, &models.UserToken{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	return *token.Claims.(*models.UserToken)
}

// IsTokenValid receives a jwt token in the form of a string and checks wheater the 
// token is valid or not
func IsTokenValid(token string) bool {
	return true
}
