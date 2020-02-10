package controllers

import (
	"net/http"
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
	a.Router.HandleFunc("/api/recipes", a.GetRecipes).Methods(http.MethodGet)
	a.Router.Handle("/api/recipes", middlewares.Handler{H: a.CreateRecipe}).Methods(http.MethodPost)
	a.Router.Handle("/api/recipes/{recipeID}", middlewares.Handler{H: a.GetRecipeByID}).Methods(http.MethodPost)
	a.Router.Handle("/api/recipes/{recipeID}", middlewares.Handler{H: a.UpdateRecipe}).Methods(http.MethodPut)
	a.Router.Handle("/api/recipes/{recipeID}", middlewares.Handler{H: a.DeleteRecipe}).Methods(http.MethodDelete)
	a.Router.HandleFunc("/api/recipes/{recipeID}/comments", a.GetRecipeComments).Methods(http.MethodGet)
	a.Router.Handle("/api/recipes/{recipeID}/comments", middlewares.Handler{H: a.AddComment}).Methods(http.MethodPost)
	//a.Router.HandleFunc("/api/recipes/{recipeid}/comments/{commentid}", nil)

	//categories
	a.Router.HandleFunc("/api/categories", nil)
	a.Router.HandleFunc("/api/categories/{id}", nil)
}