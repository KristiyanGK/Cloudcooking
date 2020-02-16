package models

// Category is used to differentiate recipes 
type Category struct {
	BaseModel
	Name string `gorm:"unique;not null"`
	Description string
	Recipes []Recipe
}
