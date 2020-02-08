package stores

import (
	"github.com/KristiyanGK/cloudcooking/models"
	"github.com/KristiyanGK/cloudcooking/persistence"
	"github.com/jinzhu/gorm"
)

// RecipeStore is a store for recipes
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

// CreateRecipe adds recipe to store
func (rs *RecipeStore) CreateRecipe(recipe models.Recipe) {
	rs.db.Create(&recipe)
}
