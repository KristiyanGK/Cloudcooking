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
		ID: user.ID,
		Username: user.Username,
		Email: user.Email,
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

func IsTokenValid() {

}
