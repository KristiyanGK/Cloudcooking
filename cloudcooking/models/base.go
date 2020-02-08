package models

import (
	"github.com/jinzhu/gorm"
)

// BaseModel is the base model for all models
type BaseModel struct {
	gorm.Model
}
