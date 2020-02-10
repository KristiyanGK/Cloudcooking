package models

// Recipe defines a cooking recipe
type Recipe struct {
	BaseModel
	Title string `gorm:"unique;not null"`
	Description string 
	UsedProducts string 
	Picture string
	CookingTime int
	UserID uint
	CategoryID uint
	Category Category
}