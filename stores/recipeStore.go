package stores

import (
	"fmt"
	"github.com/KristiyanGK/cloudcooking/models"
	"github.com/KristiyanGK/cloudcooking/persistence"
	"github.com/jinzhu/gorm"
)

// RecipeStore is a store for recipes
// Implements contracts/IRecipeStore
type RecipeStore struct {
	db *gorm.DB
}

// NewRecipeStore creates a new RecipeStore
func NewRecipeStore() *RecipeStore {
	return &RecipeStore{persistence.GetDb()}
}

// GetAllRecipes returns all recipes from store
func (rs *RecipeStore) GetAllRecipes() []models.Recipe {
	var result []models.Recipe

	rs.db.Find(&result)

	return result
}

// AddRecipe adds recipe to store
func (rs *RecipeStore) AddRecipe(recipe models.Recipe) models.Recipe {
	rs.db.Create(&recipe)

	return recipe
}

// GetRecipeByID finds recipe by given id and returns it.
// Returns error if not found
func (rs *RecipeStore) GetRecipeByID(id models.ModelID) (models.Recipe, error) {
	var recipeResult models.Recipe

	if err := rs.db.Where("id = ?", id).First(&recipeResult).Error; err != nil {
		return models.Recipe{}, fmt.Errorf("Recipe with id: %s not found", id)
	}

	return recipeResult, nil
}

//DeleteRecipeByID deletes recipe by given id from store
//Returns error if recipe is not found
func (rs *RecipeStore) DeleteRecipeByID(id models.ModelID) error {
	if rowsAffected := rs.db.Delete(&models.Recipe{BaseModel: models.BaseModel{ID: id}}).RowsAffected; rowsAffected <= 0 {
		return fmt.Errorf("Recipe with id: %s not found", id)
	}

	return  nil
}

// UpdateRecipeByID updates a given recipe by id with new fields
// Return error if recipe not found
func (rs *RecipeStore) UpdateRecipeByID(id models.ModelID, newRecipe models.Recipe) error {

	var oldRecipe models.Recipe

	if err := rs.db.Where("id = ?", id).First(&oldRecipe).Error; err != nil {
		return fmt.Errorf("Recipe with id: %s not found", id)
	}

	oldRecipe.CookingTime = newRecipe.CookingTime
	oldRecipe.Description = newRecipe.Description
	oldRecipe.Title = newRecipe.Title
	oldRecipe.UsedProducts = newRecipe.UsedProducts

	rs.db.Save(&oldRecipe)

	return nil
}
