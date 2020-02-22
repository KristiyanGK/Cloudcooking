package contracts

import (
	"github.com/KristiyanGK/cloudcooking/models"
)

// IRecipeStore is interface for a recipe store
type IRecipeStore interface {
	GetRecipes(limit, offset int, category string) ([]models.Recipe, int)
	AddRecipe(recipe models.Recipe) (models.Recipe, error)
	GetRecipeByID(id models.ModelID) (models.Recipe, error)
	DeleteRecipeByID(id models.ModelID) error
	UpdateRecipeByID(id models.ModelID, newRecipe models.Recipe) error
}
