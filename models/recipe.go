package models

// Recipe defines a cooking recipe
type Recipe struct {
	BaseModel
	Title string
	Description string
	UsedProducts string
	Picture string
	CookingTime int
	UserID uint
	CategoryID uint
}