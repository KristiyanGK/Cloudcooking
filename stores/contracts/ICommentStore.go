package contracts

import (
	"github.com/KristiyanGK/cloudcooking/models"
)

type ICommentStore interface {
	GetRecipeComments(recipeID models.ModelID) []models.Comment
	AddComment(comment models.Comment) models.Comment
}