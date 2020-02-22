package controllers

import (
	"github.com/KristiyanGK/cloudcooking/api/auth"
	"github.com/go-chi/chi"
	"encoding/json"
	"github.com/KristiyanGK/cloudcooking/models"
	rvm "github.com/KristiyanGK/cloudcooking/api/viewmodels/recipes"
	"github.com/go-playground/validator/v10"
	"net/http"
)

//ListRecipes GET /api/recipes
func (a *App) ListRecipes(w http.ResponseWriter, r *http.Request) {
	recipes := a.RecipeStore.GetAllRecipes()

	recipesVM := []rvm.RecipesListVm{}

	for _, recipe := range recipes {
		recipeVM := rvm.RecipesListVm {
			ID: string(recipe.ID),
			Title: recipe.Title,
			Description: recipe.Description,
			Picture: recipe.Picture,
			Category: recipe.Category,
			CookingTime: recipe.CookingTime,
			UsedProducts: recipe.UsedProducts,
			User: recipe.User.Username,
		}

		recipesVM = append(recipesVM, recipeVM)
	}

	respondWithJSON(w, http.StatusOK, recipesVM)
}

//CreateRecipe POST /api/recipes
func (a *App) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	token := ctx.Value("token").(string)

	claims := auth.ParseToken(a.APISecret, token)

	recipeFormValues := rvm.RecipesFormReceivedVm{}

	decoder := json.NewDecoder(r.Body)
	var err error

	if err = decoder.Decode(&recipeFormValues); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	recipe := &models.Recipe{
		Title: recipeFormValues.Title,
		Description: recipeFormValues.Description,
		Picture: recipeFormValues.Picture,
		CookingTime: recipeFormValues.CookingTime,
		CategoryID: models.ModelID(recipeFormValues.CategoryID),
		UsedProducts: recipeFormValues.UsedProducts,
		UserID: models.ModelID(claims.UserID),
	}

	defer r.Body.Close()

	err = a.Validator.Struct(recipe)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		respondWithValidationError(errs.Translate(a.Translator), w)
		return
	}

	var createdRecipe models.Recipe

	createdRecipe, err = a.RecipeStore.AddRecipe(*recipe)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	recipeVM := *rvm.NewRecipesDetailsVm(createdRecipe)

	respondWithJSON(w, http.StatusCreated, recipeVM)
}

//GetRecipeByID GET /recipe/{recipeID}
func (a *App) GetRecipeByID(w http.ResponseWriter, r *http.Request) {
	recipeID := models.ModelID(chi.URLParam(r, "recipeID"))

	recipe, err := a.RecipeStore.GetRecipeByID(recipeID)

	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	recipeVM := *rvm.NewRecipesDetailsVm(recipe)

	respondWithJSON(w, http.StatusOK, recipeVM)
}

//DeleteRecipe DELETE /api/recipes/{recipeID}
func (a *App) DeleteRecipe(w http.ResponseWriter, r *http.Request) {	
	recipeID := models.ModelID(chi.URLParam(r, "recipeID"))

	err := a.RecipeStore.DeleteRecipeByID(recipeID)

	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

//UpdateRecipe PUT /api/recipes/{recipeID}
func (a *App) UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	recipeID := models.ModelID(chi.URLParam(r, "recipeID"))

	var recipe models.Recipe

	decoder := json.NewDecoder(r.Body)

	var err error

	if err = decoder.Decode(&recipe); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	defer r.Body.Close()

	err = a.RecipeStore.UpdateRecipeByID(recipeID, recipe)

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
