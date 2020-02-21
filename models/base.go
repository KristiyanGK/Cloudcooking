package models

import (
	"github.com/google/uuid"
)

// ModelID is the type of id used for the models
type ModelID string

// BaseModel is the base model for all models
type BaseModel struct {
	ID ModelID `gorm:"primary_key" json:"id"`
}

// BeforeCreate hook that creates a uuid for the model
func (m *BaseModel) BeforeCreate() {
	m.ID = ModelID(uuid.New().String())
}