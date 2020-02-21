package models

// Recipe defines a cooking recipe
type Recipe struct {
	BaseModel
	Title string `gorm:"unique;not null"`
	Description string 
	UsedProducts string `validate:"required"`
	Picture string
	CookingTime int `validate:"required"`
	UserID ModelID
	CategoryID ModelID
	Category Category
	Comments []Comment
}