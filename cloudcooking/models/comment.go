package models

// Comment defines a recipe comment
type Comment struct {
	BaseModel
	Content string
	UserID uint
	RecipeID uint
}