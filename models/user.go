package models

// User is application user
type User struct {
	BaseModel
	Username string `gorm:"unique;not null"`
	Email string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	FirstName string
	LastName string
	Picture string
	RoleID uint
	Role Role
	Recipes []Recipe
	Comments []Comment
}
