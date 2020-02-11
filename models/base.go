package models

// BaseModel is the base model for all models
type BaseModel struct {
	ID uint `gorm:"primary_key"`
}
