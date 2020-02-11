package controllers

import (
	"github.com/KristiyanGK/cloudcooking/api/auth"
	"encoding/json"
	"net/http"
	uvm "github.com/KristiyanGK/cloudcooking/api/viewmodels/users"
)

func (a *App) Register(w http.ResponseWriter, r *http.Request) {
	var registerInfo uvm.UserRegisterVM

	decoder := json.NewDecoder(r.Body)
	var err error

	if err = decoder.Decode(&registerInfo); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	//TODO validate user data
}

func (a *App) Login(w http.ResponseWriter, r *http.Request) {
	var loginInfo uvm.UserLoginVM

	decoder := json.NewDecoder(r.Body)
	var err error

	if err = decoder.Decode(&loginInfo); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	user := a.UserStore.GetUserByUsername(loginInfo.Username)

	token := auth.GenerateToken(user)

	w.Header().Set("Authorization", "Bearer " + token)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logged in!"))
}
