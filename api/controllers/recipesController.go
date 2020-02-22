package controllers

import (
	"fmt"
	"github.com/KristiyanGK/cloudcooking/api/errors"
	"strconv"
	"github.com/KristiyanGK/cloudcooking/api/auth"
	"github.com/go-chi/chi"
	"encoding/json"
	"github.com/KristiyanGK/cloudcooking/models"
	rvm "github.com/KristiyanGK/cloudcooking/api/viewmodels/recipes"
	cvm "github.com/KristiyanGK/cloudcooking/api/viewmodels/comments"
	"github.com/go-playground/validator/v10"
	"net/http"
)

//ListRecipes GET /api/recipes
func (a *App) ListRecipes(w http.ResponseWriter, r *http.Request) {

	limitStr := r.URL.Query().Get("limit")

	limit, err := strconv.Atoi(limitStr)

	if err != nil {
		limit = 5
	}

	offsetSTR := r.URL.Query().Get("offset")

	offset, err := strconv.Atoi(offsetSTR)

	if err != nil {
		offset = 0
	}

	category := r.URL.Query().Get("category")

	fmt.Println(category);

	recipes, count := a.RecipeStore.GetRecipes(limit, offset, category)

	recipesVM := rvm.RecipesListVm{}

	recipesVM.Count = count

	for _, recipe := range recipes {
		recipeVM := rvm.NewRecipesListItemVm(recipe)

		recipesVM.Recipes = append(recipesVM.Recipes, recipeVM)
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

	decoder := json.NewDecoder(r.Body)

	recipeFormValues := rvm.RecipesFormReceivedVm{}

	var err error

	if err = decoder.Decode(&recipeFormValues); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	ctx := r.Context()

	token := ctx.Value("token").(string)

	claims := auth.ParseToken(a.APISecret, token)

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

	err = a.RecipeStore.UpdateRecipeByID(recipeID, *recipe)

	if err != nil {
		
		switch e := err.(type) {
			case errors.StatusError:
				respondWithError(w, e.Code, e.Err.Error())
			default:
				respondWithError(w, http.StatusInternalServerError, e.Error())
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}

//GetRecipeComments GET /api/recipes/{recipeID}/comments
func (a *App) GetRecipeComments(w http.ResponseWriter, r *http.Request) {
	recipeID := models.ModelID(chi.URLParam(r, "recipeID"))

	comments := a.CommentStore.GetRecipeComments(recipeID)

	commentsVM := []cvm.CommentResponseVm{}

	for _, c := range comments {
		commentVM := cvm.CommentResponseVm {
			ID: string(c.BaseModel.ID),
			Content: c.Content,
			CreatedAt: c.CreatedAt,
			User: c.User.Username,
		}

		commentsVM = append(commentsVM, commentVM)
	}

	respondWithJSON(w, http.StatusOK, commentsVM)
}

//AddComment POST /api/recipes/{recipeID}/comments
func (a *App) AddComment(w http.ResponseWriter, r *http.Request) {
	var commentVM cvm.CommentFormReceivedVm

	decoder := json.NewDecoder(r.Body)

	var err error

	if err = decoder.Decode(&commentVM); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	ctx := r.Context()

	token := ctx.Value("token").(string)

	recipeID := models.ModelID(chi.URLParam(r, "recipeID"))

	claims := auth.ParseToken(a.APISecret, token)

	comment := models.Comment {
		Content: commentVM.Content,
		RecipeID: recipeID,
		UserID: models.ModelID(claims.UserID),
	}

	commentResult := a.CommentStore.AddComment(comment)

	commentResultVM := cvm.CommentResponseVm {
		ID: string(commentResult.BaseModel.ID),
		Content: commentResult.Content,
		CreatedAt: commentResult.CreatedAt,
		User: claims.Username,
	}

	respondWithJSON(w, http.StatusCreated, commentResultVM)
}
