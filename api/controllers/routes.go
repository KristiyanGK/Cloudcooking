package controllers

import (
	"github.com/KristiyanGK/cloudcooking/api/middlewares"
)

// RegisterRoutes registers routes of the app
func (a *App) RegisterRoutes() {
	//users
	a.Router.HandleFunc("/api/users", nil)
	a.Router.HandleFunc("/api/users/{id}", nil)
	a.Router.HandleFunc("/api/login", nil)
	a.Router.HandleFunc("/api/logout", nil)

	//recipes
	a.Router.Handle("/api/recipes", middlewares.Handler{H: a.GetRecipes})
	a.Router.HandleFunc("/api/recipes/{id}", nil)
	a.Router.HandleFunc("/api/recipes/{id}/comments", nil)
	a.Router.HandleFunc("/api/recipes/{recipeid}/comments/{commentid}", nil)

	//categories
	a.Router.HandleFunc("/api/categories", nil)
	a.Router.HandleFunc("/api/categories/{id}", nil)
}