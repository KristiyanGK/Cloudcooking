package controllers

import (
	"github.com/go-chi/chi"
	"encoding/json"
	"github.com/KristiyanGK/cloudcooking/models"
	"strconv"
	"github.com/go-playground/validator/v10"
	"net/http"
)

//ListRecipes GET /api/recipes
func (a *App) ListRecipes(w http.ResponseWriter, r *http.Request) {
	recipes := a.RecipeStore.GetAllRecipes()

	respondWithJSON(w, http.StatusOK, recipes)
}

//CreateRecipe POST /api/recipes
func (a *App) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	recipe := &models.Recipe{}

	decoder := json.NewDecoder(r.Body)
	var err error

	if err = decoder.Decode(recipe); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	defer r.Body.Close()

	err = a.Validator.Struct(recipe)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		respondWithValidationError(errs.Translate(a.Translator), w)
		return
	}

	a.RecipeStore.AddRecipe(*recipe)

	w.WriteHeader(http.StatusCreated)
}

//GetRecipeByID GET /recipe/{recipeID}
func (a *App) GetRecipeByID(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "recipeID")

	recipeID, err := strconv.Atoi(param)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	recipe, err := a.RecipeStore.GetRecipeByID(uint(recipeID))

	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, recipe)
}

//DeleteRecipe DELETE /api/recipes/{recipeID}
func (a *App) DeleteRecipe(w http.ResponseWriter, r *http.Request) {	
	param := chi.URLParam(r, "recipeID")

	recipeID, err := strconv.Atoi(param)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	err = a.RecipeStore.DeleteRecipeByID(uint(recipeID))

	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

//UpdateRecipe PUT /api/recipes/{recipeID}
func (a *App) UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "recipeID")

	recipeID, err := strconv.Atoi(param)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	var recipe models.Recipe

	decoder := json.NewDecoder(r.Body)

	if err = decoder.Decode(&recipe); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	defer r.Body.Close()

	err = a.RecipeStore.UpdateRecipeByID(uint(recipeID), recipe)

	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

//GetRecipeComments GET /api/recipes/{recipeID}/comments
func (a *App) GetRecipeComments(w http.ResponseWriter, r *http.Request) {

}

//AddComment POST /api/recipes/{recipeID}/comments
func (a *App) AddComment(w http.ResponseWriter, r *http.Request) {
}
