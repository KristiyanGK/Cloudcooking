package controllers

import (
	"github.com/KristiyanGK/cloudcooking/models"
	"github.com/KristiyanGK/cloudcooking/api/auth"
	"github.com/KristiyanGK/cloudcooking/api/utils"
	uvm "github.com/KristiyanGK/cloudcooking/api/viewmodels/users"
	"encoding/json"
	"net/http"
)

// Register receives username, password, email, RepeatPassword
// If the credentials are valid 
// Returns a jwt token via json response
func (a *App) Register(w http.ResponseWriter, r *http.Request) {
	var registerInfo uvm.UserRegisterVM

	decoder := json.NewDecoder(r.Body)
	var err error

	if err = decoder.Decode(&registerInfo); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	//TODO validate user data

	registerInfo.Password, err = utils.HashPassword(registerInfo.Password)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not register user")
		return
	}

	var user models.User

	user, err = a.UserStore.AddUser(registerInfo)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	token := auth.GenerateToken(a.APISecret, user)

	result := models.UserResult{
		Username: user.Username,
		Picture: user.Picture,
		Role: user.Role.Name,
		Token: token,
	}

	respondWithJSON(w, http.StatusOK, result)
}

// Login receives username and password and logs in the user if the credentials are valid
// Returns a jwt token via json response 
func (a *App) Login(w http.ResponseWriter, r *http.Request) {
	var loginInfo uvm.UserLoginVM

	decoder := json.NewDecoder(r.Body)
	var err error

	if err = decoder.Decode(&loginInfo); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	//TODO validate user data

	user := a.UserStore.GetUserByEmail(loginInfo.Email)

	if user.ID == "" || utils.CheckPasswordHash(loginInfo.Password, user.Password) {
		respondWithError(w, http.StatusBadRequest, "Invalid credentials")
		return
	}

	token := auth.GenerateToken(a.APISecret, user)

	result := models.UserResult{
		Username: user.Username,
		Picture: user.Picture,
		Role: user.Role.Name,
		Token: token,
	}

	respondWithJSON(w, http.StatusOK, result)
}
