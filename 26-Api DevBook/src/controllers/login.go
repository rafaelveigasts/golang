package controllers

import (
	"api-devbook/src/database"
	"api-devbook/src/models"
	"api-devbook/src/repository"
	"api-devbook/src/responses"
	"api-devbook/src/security"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	println(string(body))
	
	var user models.User
	var userDB models.User
	
	if err = json.Unmarshal(body, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.ConnectToDatabase()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repository.NewUserRepository(db)
	userDB, err = repository.GetUserByEmail(user.Email)
	
	
	if err = security.VerifyPassword(userDB.Password, user.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	w.Write([]byte(fmt.Sprintf("Welcome %s", userDB.Name)))
	

}