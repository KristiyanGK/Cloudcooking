package controllers

import (
	"github.com/KristiyanGK/cloudcooking/api/middlewares"
	"github.com/go-chi/cors"
	"github.com/go-chi/chi"
)

// RegisterRoutes registers routes of the app
func (a *App) RegisterRoutes() {
	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	//a.Router.Use(middlewares.LoggingMiddleware)
	a.Router.Use(cors.Handler)

	a.Router.Post("/api/login", a.Login)
	a.Router.Post("/api/register", a.Register)

	a.Router.Route("/api/recipes", func(r chi.Router) {
		r.Use(middlewares.AuthenticationMiddleware)
		r.Get("/", a.ListRecipes)
		r.Post("/", a.CreateRecipe)

		r.Route("/{recipeID}", func(r chi.Router) {
			r.Get("/", a.GetRecipeByID)
			r.Put("/", a.UpdateRecipe)
			r.Delete("/", a.DeleteRecipe)
		})
	})

	a.Router.Route("/api/categories", func(r chi.Router) {
		r.Get("/", a.ListCategories)
	})
}
