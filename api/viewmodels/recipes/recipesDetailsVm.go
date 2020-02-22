package recipes

import (
	"github.com/KristiyanGK/cloudcooking/models"
)

type RecipesDetailsVm struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Picture string `json:"picture"`
	CookingTime int `json:"cookingTime"`
	Category models.Category `json:"category"`
	UsedProducts string `json:"usedProducts"`
	User string `json:"user"`
}

func NewRecipesDetailsVm(recipe models.Recipe) *RecipesDetailsVm {
	return &RecipesDetailsVm {
		ID: string(recipe.ID),
		Title: recipe.Title,
		Description: recipe.Description,
		Picture: recipe.Picture,
		CookingTime: recipe.CookingTime,
		Category: recipe.Category,
		UsedProducts: recipe.UsedProducts,
		User: recipe.User.Username,
	}
}