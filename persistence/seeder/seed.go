package seeder

import (
	"github.com/KristiyanGK/cloudcooking/persistence"
	"github.com/jinzhu/gorm"
	"github.com/KristiyanGK/cloudcooking/models"
)

var roles = []models.Role {
	{
		Name: "HomeCook",
	},
	{
		Name: "Administrator",
	},
}

var db *gorm.DB

// Seed seeds sample data into the database
func Seed() {
	db = persistence.GetDb()

	if db == nil {
		panic("Cannot get database to seed")
	}

	//seedRoles()
	//seedUsers()
}

func seedRoles() {
	var rolesCount int
	db.Model(&models.Role{}).Count(&rolesCount)

	if rolesCount > 0 {
		return
	}

	for _, r := range roles {
		db.Model(&models.Role{}).Create(&r)
	}
}

func seedUsers() {
	var usersCount int
	db.Model(&models.User{}).Count(&usersCount)

	if usersCount > 0 {
		return
	}
}

func seedCategories() {

}

func seedRecipes() {

}
