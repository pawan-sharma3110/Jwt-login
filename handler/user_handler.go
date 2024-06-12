package handler

import (
	"jwt/db"
	"jwt/model"
	"jwt/utils"
	"net/http"
)

// strings

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	db, err := db.DbIn()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	var payload model.User
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	userId, err := utils.InsertUser(db, payload)
	if err != nil {
		if err.Error() == "user already exists" {
			utils.WriteError(w, http.StatusBadRequest, err)
		} else {
			utils.WriteError(w, http.StatusInternalServerError, err)
		}
		return
	}

	utils.WriteJson(w, http.StatusCreated, map[string]int{"user_id": userId})
}
func Login(w http.ResponseWriter, r *http.Request) {

	db, err := db.DbIn()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	var payload model.User
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	token, err := utils.UserLogin(db, w, payload)
	if err != nil {
		// utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	err = utils.WriteJson(w, http.StatusOK, map[string]string{"token": token})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}
