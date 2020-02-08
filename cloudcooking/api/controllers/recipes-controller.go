package controllers

import (
	"github.com/KristiyanGK/cloudcooking/api/responses"
	"net/http"
)

//GetRecipes GET /api/recipes
func (a *App) GetRecipes(w http.ResponseWriter, r *http.Request) error {
	recipes := a.RecipeStore.GetAllRecipes()
	
	responses.JSONResponse(w, recipes, http.StatusOK)
	
	return nil
}

//CreateUser POST /api/recipes
func CreateUser(w http.ResponseWriter, r *http.Request) error {
	
	return nil
}
