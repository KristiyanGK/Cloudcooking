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
	RoleID uint
	Role Role
	Recipes []Recipe
	Comments []Comment
}

//UserToken is user info in token
type UserToken struct {
	ID uint `json:"userid"`
	Username string `json:"username"`
	Email string `json:"email"`
	Picture string `json:"picture"`
	Role string `json:"userRole"`
	jwt.StandardClaims
}