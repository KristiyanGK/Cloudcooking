package contracts

import (
	"github.com/KristiyanGK/cloudcooking/models"
)

type IUserStore interface {
	GetUserByUsername(username string) models.User
}