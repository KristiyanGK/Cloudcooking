package stores

import (
	"fmt"
	"github.com/KristiyanGK/cloudcooking/models"
	"github.com/KristiyanGK/cloudcooking/persistence"
	"github.com/jinzhu/gorm"
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

func (us *UserStore) GetUserByUsername(username string) models.User {
	var user models.User

	query := `
		SELECT u.id, u.username, u.email, u.picture, r.id, r.name
		FROM users AS u
		INNER JOIN roles AS r ON u.role_id = r.id
		WHERE u.username = ?
	`

	rows, err := us.db.DB().Query(query, username)

	if err != nil {
		fmt.Println(err)
	}

	/*
	SELECT r.id, r.name
		FROM users AS u
		JOIN roles AS r ON u.role_id = r.id
		WHERE u.username = ?
	*/
	//u.id, u.username, u.email, u.picture, 
	//uid, u, email, picture,
	//&uid,&u,&email,&picture,

	var uid, u, email, picture, rid, name string

	for rows.Next() {
		rows.Scan(&uid,&u,&email,&picture,&rid,&name)
	}

	// rows, err := us.db.DB().Query(query, username)

	// rows, err := us.db.Table("users").Joins("left join roles on roles.id = users.role_id").Where("users.username=?", username).Rows()

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// defer rows.Close()

	// for rows.Next() {
	// 	rows.Scan(&user.ID, &user.Username, &user.Email, &user.Picture, &role.ID, &role.Name)
	// }

	//var role models.Role
	
	//us.db.Where("username = ?", username).First(&user)

	return user
}