package models

import (
	"time"
)

// Comment defines a recipe comment
type Comment struct {
	BaseModel
	Content string
	UserID ModelID
	User User
	RecipeID ModelID
	Recipe Recipe
	CreatedAt time.Time
}