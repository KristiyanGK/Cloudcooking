package controllers

import (
	"net/http"
)

//ListCategories GET /api/categories
func (a *App) ListCategories(w http.ResponseWriter, r *http.Request) {
	categories := a.CategoryStore.GetAllCategories()

	respondWithJSON(w, http.StatusOK, categories)
}