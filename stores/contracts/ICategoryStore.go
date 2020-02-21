package contracts

import (
	"github.com/KristiyanGK/cloudcooking/models"
)

type ICategoryStore interface {
	GetAllCategories() []models.Category
}