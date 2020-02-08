package models

// Category is used to differentiate recipes 
type Category struct {
	BaseModel
	Name string
	Description string
	Recipes []Recipe
}
