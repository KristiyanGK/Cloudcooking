package stores

import (
	"github.com/KristiyanGK/cloudcooking/models"
	"github.com/KristiyanGK/cloudcooking/persistence"
	"github.com/jinzhu/gorm"
)

// CategoryStore is a store for categories
// Implements contracts/ICategoryStore
type CategoryStore struct {
	db *gorm.DB
}

// NewCategoryStore creates a new CategoryStore
func NewCategoryStore() *CategoryStore {
	return &CategoryStore{persistence.GetDb()}
}

// GetAllCategories returns all categories from db
func (cs *CategoryStore) GetAllCategories() []models.Category {
	var result []models.Category

	cs.db.Find(&result)

	return result
}
