package recipes

import (
	"time"
	"github.com/KristiyanGK/cloudcooking/models"
)

// RecipesListVm is a container for recipes with count
type RecipesListVm struct {
	Recipes []RecipesListItemVm `json:"recipes"`
	Count int `json:"count"`
}

// RecipesListItemVm is the individual recipe item in RecipesListVm
type RecipesListItemVm struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Picture string `json:"picture"`
	Category models.Category `json:"category"`
	CookingTime int `json:"cookingTime"`
	UsedProducts string `json:"usedProducts"`
	User string `json:"user"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewRecipesListItemVm(recipe models.Recipe) RecipesListItemVm{
	recipeVM := RecipesListItemVm {
		ID: string(recipe.ID),
		Title: recipe.Title,
		Description: recipe.Description,
		Picture: recipe.Picture,
		Category: recipe.Category,
		CookingTime: recipe.CookingTime,
		UsedProducts: recipe.UsedProducts,
		User: recipe.User.Username,
		CreatedAt: recipe.CreatedAt,
	}

	return recipeVM
}