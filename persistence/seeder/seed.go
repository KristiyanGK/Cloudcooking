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

	var users = []models.User {
		models.User{
			Username: "Ivan43",
			Email: "ivan@test.bg",
			Password: "123",
			FirstName: "Ivan",
			LastName: "Ivanov",
			Picture: "",
			Role: roles[0],
			Recipes: []models.Recipe{
				{
					Title: "Test Recipe",
					Description: `
						This is a test description.
						ba balba ddv gdsg dsg sdg sdg
					`,
					UsedProducts: "bread;wheat;flower;stuff",
					Picture: "",
					CookingTime: 15,
					Category: models.Category{
						Name: "Test1",
						Description: `
							Test category description mate
						`,
					},
				},
				{
					Title: "Ivan best recipe",
					Description: `
						This is a test description. For IVansd
						ba balba ddv gdsg dsg sdg sdg
					`,
					UsedProducts: "bricks;tteeh;he;egd",
					Picture: "",
					CookingTime: 123,
					Category: models.Category{
						Name: "Test2",
						Description: `
							Another Test category
						`,
					},
				},
			},
		},
		models.User{
			Username: "pesho5664",
			Email: "pesh@abv.bg",
			Password: "421",
			FirstName: "pesho",
			LastName: "peshov",
			Picture: "",
			Role: roles[1],
		},
	}

	for _, u := range users {
		db.Create(&u)
	}
}

func seedCategories() {

}

func seedRecipes() {

}
