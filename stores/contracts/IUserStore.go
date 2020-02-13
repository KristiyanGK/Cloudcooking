package contracts

import (
	"github.com/KristiyanGK/cloudcooking/models"
	uvm "github.com/KristiyanGK/cloudcooking/api/viewmodels/users"
)

// IUserStore is an interface for a user container
type IUserStore interface {
	GetUserByUsername(username string) models.User
	AddUser(registerInfo uvm.UserRegisterVM) (models.User, error)
}