package controllers

import (
	"github.com/go-chi/chi"
)

// RegisterRoutes registers routes of the app
func (a *App) RegisterRoutes() {
	a.Router.Post("/api/login", a.Login)

	a.Router.Route("/api/recipes", func(r chi.Router) {
		r.Get("/", a.ListRecipes)
		r.Post("/", a.CreateRecipe)

		r.Route("/{recipeID}", func(r chi.Router) {
			r.Get("/", a.GetRecipeByID)
			r.Put("/", a.UpdateRecipe)
			r.Delete("/", a.DeleteRecipe)
		})
	})
}
