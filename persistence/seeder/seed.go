package seeder

import (
	"github.com/KristiyanGK/cloudcooking/persistence"
	"github.com/jinzhu/gorm"
	"github.com/KristiyanGK/cloudcooking/models"
)

var db *gorm.DB

// Seed seeds sample data into the database
func Seed() {
	db = persistence.GetDb()

	if db == nil {
		panic("Cannot get database to seed")
	}

	seedRoles()
	seedCategories()
	seedRecipes()
}

var roles = []models.Role {
	{
		Name: "HomeCook",
	},
	{
		Name: "Administrator",
	},
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

var categories = [] models.Category {
	{
		Name: "Cat1",
		Description: "This is a description for cat 1 ",
	},
	{
		Name: "Special Stuff",
		Description: `Special StuffSpecial StuffSpecial StuffSpecial StuffSpecial Stuff
		vSpecial StuffSpecial Stuff`,
	},
	{
		Name: "god tier food",
		Description: `Food fir for the gods`,
	},
}

func seedCategories() {
	var catCount int
	db.Model(&models.Category{}).Count(&catCount)

	if catCount > 0 {
		return
	}

	for _, c := range categories {
		db.Model(&models.Category{}).Create(&c)
	}
}

var recipes = [] models.Recipe {
	{
		Title: "Recipe1",
		Description: `This is a great description`,
		UsedProducts: "stuff1;stuff2;stuff3",
		Picture: "/something",
		CookingTime: 512,
	},
	{
		Title: "Recipe2",
		Description: `This is a great description for the whole family
			m	onff fdf sdf sdf sdf dsf sd`,
		UsedProducts: "stuff1;stuff2;stuff3;stuff6;stuff95324;434343434e",
		Picture: "/something/other",
		CookingTime: 420,
	},
	{
		Title: "Recipe infinite de great",
		Description: `Ezio alditore la firence`,
		UsedProducts: "stuff1;stuff2;stuff3;stuff6;stuff95324;434343434e",
		Picture: "/something/other",
		CookingTime: 420,
	},
	{
		Title: "Pizza",
		Description: `A recipe for the most cultured of people`,
		UsedProducts: "dough;tomatos;sauce;peperoni",
		Picture: "/something/other",
		CookingTime: 5556,
	},
	{
		Title: "Recipe4",
		Description: `This Recipe will be remebered throughout the ages`,
		UsedProducts: "ahem;item2;item3;item4;item5",
		Picture: "/actual/pic",
		CookingTime: 6,
	},
}

func seedRecipes() {
	var recipesCount int
	db.Model(&models.Recipe{}).Count(&recipesCount)

	if recipesCount > 0 {
		return
	}

	var cats []models.Category
	db.Find(&cats)

	for i, r := range recipes {
		r.CategoryID = cats[i % (len(cats))].ID
		db.Model(&models.Recipe{}).Create(&r)
	}
}
