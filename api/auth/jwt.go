package auth

import (
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/KristiyanGK/cloudcooking/models"
	"os"
)

/*
	ID uint
	Username string
	Email string
	Picture string
	Role string
	jwt.StandardClaims
*/

var secret = os.Getenv("API_SECRET")

// GenerateToken generates a jwt token and returns it
func GenerateToken(user models.User) string {
	expirationTime := time.Now().Add(24 * time.Hour)

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

	tokenString, _ := token.SignedString(secret)

	return tokenString
}

func IsTokenValid() {

}
