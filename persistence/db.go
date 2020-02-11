package persistence

import (
	"fmt"
	"log"

	"github.com/KristiyanGK/cloudcooking/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql database driver
)

var db *gorm.DB

// InitializeDb initialized the db connection
func InitializeDb(driver, host, port, user, password, name string) {
	var connectionString string

	if driver == "postgres" {
		connectionString = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, name, password)
	} else if driver == "mysql" {
		connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, name)
	} else {
		log.Fatal("Unsupported db driver!")
	}

	var err error

	db, err = gorm.Open(driver, connectionString)

	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(true)

	db.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Category{},
		&models.Recipe{},
		&models.Comment{})
}

// GetDb Returns gorm database instance
func GetDb() *gorm.DB {
	return db
}