package controllers

import (
	"encoding/json"
	"github.com/KristiyanGK/cloudcooking/models"
	"fmt"
	"strconv"
	"github.com/KristiyanGK/cloudcooking/api/errors"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/KristiyanGK/cloudcooking/api/responses"
	"net/http"
)

//GetRecipes GET /api/recipes
func (a *App) GetRecipes(w http.ResponseWriter, r *http.Request) {
	recipes := a.RecipeStore.GetAllRecipes()
	
	responses.JSONResponse(w, recipes, http.StatusOK)
}

//CreateRecipe POST /api/recipes
func (a *App) CreateRecipe(w http.ResponseWriter, r *http.Request) error {
	recipe := &models.Recipe{}

	decoder := json.NewDecoder(r.Body)
	var err error

	if err = decoder.Decode(recipe); err != nil {
		return errors.StatusError{Err:fmt.Errorf("Invalid request body"), Code: http.StatusBadRequest}
	}

	defer r.Body.Close()

	err = a.Validator.Struct(recipe)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		return errors.StatusError{Err:fmt.Errorf(errs.Error()), Code: http.StatusBadRequest}
	}

	res := a.RecipeStore.AddRecipe(*recipe)

	responses.CreatedResponse(w, "/recipes/", res.ID)

	return nil
}

//GetRecipeByID GET /recipe/{recipeID}
func (a *App) GetRecipeByID(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)

	recipeID, err := strconv.Atoi(params["recipeID"])

	if err != nil {
		return &errors.StatusError{Err: fmt.Errorf("Invalid id"), Code: http.StatusBadRequest}
	}

	recipe, err := a.RecipeStore.GetRecipeByID(uint(recipeID))

	if err != nil {
		return &errors.StatusError{Err: err, Code: http.StatusNotFound}
	}

	responses.JSONResponse(w, recipe, http.StatusOK)

	return nil
}

//DeleteRecipe DELETE /api/recipes/{recipeID}
func (a *App) DeleteRecipe(w http.ResponseWriter, r *http.Request) error {	
	params := mux.Vars(r)

	recipeID, err := strconv.Atoi(params["recipeID"])

	if err != nil {
		return &errors.StatusError{Err: fmt.Errorf("Invalid id"), Code: http.StatusBadRequest}
	}

	err = a.RecipeStore.DeleteRecipeByID(uint(recipeID))

	if err != nil {
		return &errors.StatusError{Err: err, Code: http.StatusNotFound}
	}

	w.WriteHeader(http.StatusOK)

	return nil
}

//UpdateRecipe PUT /api/recipes/{recipeID}
func (a *App) UpdateRecipe(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)

	recipeID, err := strconv.Atoi(params["recipeID"])

	if err != nil {
		return &errors.StatusError{Err: fmt.Errorf("Invalid id"), Code: http.StatusBadRequest}
	}

	var recipe models.Recipe

	decoder := json.NewDecoder(r.Body)

	if err = decoder.Decode(&recipe); err != nil {
		return errors.StatusError{Err:fmt.Errorf("Invalid request body"), Code: http.StatusBadRequest}
	}

	defer r.Body.Close()

	err = a.RecipeStore.UpdateRecipeByID(uint(recipeID), recipe)

	if err != nil {
		return &errors.StatusError{Err: err, Code: http.StatusNotFound}
	}

	w.WriteHeader(http.StatusOK)

	return nil
}

//GetRecipeComments GET /api/recipes/{recipeID}/comments
func (a *App) GetRecipeComments(w http.ResponseWriter, r *http.Request) {

}

//AddComment POST /api/recipes/{recipeID}/comments
func (a *App) AddComment(w http.ResponseWriter, r *http.Request) error {
	return nil
}
