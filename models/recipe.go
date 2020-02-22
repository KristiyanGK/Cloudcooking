package models

// Recipe defines a cooking recipe
type Recipe struct {
	BaseModel
	Title string `gorm:"unique;not null" json:"title"`
	Description string `json:"description"`
	UsedProducts string `validate:"required" json:"usedProducts"`
	Picture string `json:"picture"`
	CookingTime int `validate:"required" json:"cookingTime"`
	UserID ModelID `json:"userid"`
	User User `json:"user"`
	CategoryID ModelID `json:"categoryid"`
	Category Category `json:"category"`
	Comments []Comment `json:"comments"`
}