package api

import (
	"github.com/KristiyanGK/cloudcooking/api/controllers"
)

var app controllers.App

//Run starts the api server
func Run(addr string) {
	app = controllers.App{}

	app.Init()

	app.Run(addr)
}
