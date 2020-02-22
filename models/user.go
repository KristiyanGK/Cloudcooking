package models

import (
	"github.com/dgrijalva/jwt-go"
)

// User is application user
type User struct {
	BaseModel
	Username string `gorm:"unique;not null" json:"username"`
	Email string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Picture string `json:"picture"`
	RoleID ModelID `json:"roleid"`
	Role Role `json:"role"`
	Recipes []Recipe `json:"recipes"`
	Comments []Comment `json:"comments"`
}

type UserResult struct {
	Username string `json:"username"`
	Picture string `json:"picture"`
	Role string `json:"userRole"`
	Token string `json:"token"`
}

//UserToken is user info in token
type UserToken struct {
	UserID string `json:"userid"`
	Username string `json:"username"`
	Picture string `json:"picture"`
	Role string `json:"userRole"`
	jwt.StandardClaims
}
