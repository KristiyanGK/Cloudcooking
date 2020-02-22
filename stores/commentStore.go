package stores

import (
	"github.com/KristiyanGK/cloudcooking/models"
	"github.com/KristiyanGK/cloudcooking/persistence"
	"github.com/jinzhu/gorm"
)

// CommentStore is a store for comments
// Implements contracts/ICommentStore
type CommentStore struct {
	db *gorm.DB
}

// NewCommentStore creates a new CommentStore
func NewCommentStore() *CommentStore {
	return &CommentStore{persistence.GetDb()}
}

// GetRecipeComments returns all comments of given recipe
func (cs *CommentStore) GetRecipeComments(recipeID models.ModelID) []models.Comment {
	var comments []models.Comment

	cs.db.
	Preload("User").
	Where("recipe_id = ?", recipeID).
	Find(&comments)

	return comments
}

// AddComment adds a comment to the store
func (cs *CommentStore) AddComment(comment models.Comment) models.Comment {
	cs.db.Create(&comment)

	return comment
}