package contracts

import (
	"github.com/KristiyanGK/cloudcooking/models"
)

// IRecipeStore is interface for a recipe store
type IRecipeStore interface {
	GetAllRecipes() []models.Recipe
	AddRecipe(recipe models.Recipe) models.Recipe
	GetRecipeByID(id uint) (models.Recipe, error)
	DeleteRecipeByID(id uint) error
	UpdateRecipeByID(id uint, newRecipe models.Recipe) error
}
