package models

import (
	"github.com/dgrijalva/jwt-go"
)

// User is application user
type User struct {
	BaseModel
	Username string `gorm:"unique;not null"`
	Email string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Picture string
	RoleID ModelID
	Role Role
	Recipes []Recipe
	Comments []Comment
}

type UserResult struct {
	Username string `json:"username"`
	Picture string `json:"picture"`
	Role string `json:"userRole"`
	Token string `json:"token"`
}

//UserToken is user info in token
type UserToken struct {
	Username string `json:"username"`
	Picture string `json:"picture"`
	Role string `json:"userRole"`
	jwt.StandardClaims
}
