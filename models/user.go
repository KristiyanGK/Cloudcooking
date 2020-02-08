package models

// User is application user
type User struct {
	BaseModel
	Username string
	Email string
	Password string
	FirstName string
	LastName string
	Picture string
	RoleID uint
	Recipes []Recipe
	Comments []Comment
}
