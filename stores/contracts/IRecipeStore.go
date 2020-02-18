package contracts

import (
	"github.com/KristiyanGK/cloudcooking/models"
	rvm "github.com/KristiyanGK/cloudcooking/api/viewmodels/recipes"
)

// IRecipeStore is interface for a recipe store
type IRecipeStore interface {
	GetAllRecipes() []rvm.RecipeListVm
	AddRecipe(recipe models.Recipe) models.Recipe
	GetRecipeByID(id models.ModelID) (models.Recipe, error)
	DeleteRecipeByID(id models.ModelID) error
	UpdateRecipeByID(id models.ModelID, newRecipe models.Recipe) error
}
