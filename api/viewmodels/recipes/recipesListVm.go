package recipes

import (
	"github.com/KristiyanGK/cloudcooking/models"
)

type RecipesListVm struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Picture string `json:"picture"`
	Category models.Category `json:"category"`
	CookingTime int `json:"cookingTime"`
	UsedProducts string `json:"usedProducts"`
	User string `json:"user"`
}