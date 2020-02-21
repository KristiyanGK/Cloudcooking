package contracts

import (
	"github.com/KristiyanGK/cloudcooking/models"
	rvm "github.com/KristiyanGK/cloudcooking/api/viewmodels/recipes"
)

// IRecipeStore is interface for a recipe store
type IRecipeStore interface {
	GetAllRecipes() []rvm.RecipeListVm
	AddRecipe(recipe models.Recipe) (models.Recipe, error)
	GetRecipeByID(id models.ModelID) (rvm.RecipeDetailsVm, error)
	DeleteRecipeByID(id models.ModelID) error
	UpdateRecipeByID(id models.ModelID, newRecipe models.Recipe) error
}
