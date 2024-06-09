package handler

import (
	"jwt/db"
	"jwt/model"
	"jwt/utils"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	db, err := db.DbIn()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
	var payload model.User
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	userId, err := utils.IsertUser(db, w, payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
	utils.WriteJson(w, http.StatusCreated, userId)
}
