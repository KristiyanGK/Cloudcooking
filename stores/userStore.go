package stores

import (
	uvm "github.com/KristiyanGK/cloudcooking/api/viewmodels/users"
	"github.com/KristiyanGK/cloudcooking/models"
	"github.com/KristiyanGK/cloudcooking/persistence"
	"github.com/jinzhu/gorm"
	"database/sql"
)

// UserStore is a store for users
// Implements contracts/IUserStore
type UserStore struct {
	db *gorm.DB
}

// NewUserStore creates a new UserStore
func NewUserStore() *UserStore {
	return &UserStore{persistence.GetDb()}
}

// AddUser adds new user to the store and returns it
// If the user already exists returns error
func (us *UserStore) AddUser(registerInfo uvm.UserRegisterVM) (models.User, error) {
	user := models.User{}

	us.db.Where("name = ?", "HomeCook").First(&user.Role)

	user.Email = registerInfo.Email
	user.Password = registerInfo.Password
	user.Username = registerInfo.Username

	us.db.Create(&user)

	return user, nil
}

// GetUserByUsername receives a username a finds the user by it
func (us *UserStore) GetUserByUsername(username string) models.User {
	var user models.User

	row := us.db.Table("users AS u").Select("u.id, u.username, u.email, u.picture, r.name").Joins("JOIN roles AS r ON u.role_id = r.id").Where("u.username = ?", username).Row()

	var picture sql.NullString

	row.Scan(&user.ID,&user.Username,&user.Email,&picture,&user.Role.Name)

	user.Picture = picture.String

	return user
}
