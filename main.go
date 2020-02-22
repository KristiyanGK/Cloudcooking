package main

import (
	"github.com/KristiyanGK/cloudcooking/api"
)
/*
For the app to work a .env file is required with the following data
	API_SECRET
	DB_DRIVER
	DB_HOST
	DB_USER
	DB_PASSWORD
	DB_NAME
	DB_PORT
*/

func main() {
	api.Run(":8080")
}
