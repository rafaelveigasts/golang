package controllers

import (
	"api-devbook/src/database"
	"api-devbook/src/models"
	"api-devbook/src/repository"
	"api-devbook/src/responses"
	"encoding/json"
	"io"
	"strconv"
	"strings"

	"net/http"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	inputRequest, err := io.ReadAll(r.Body)
	if err!= nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

		if err= json.Unmarshal(inputRequest, &user); err != nil {
			responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("create"); err != nil {
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
	repository.CreateUser(user)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	nameOrAlias := strings.ToLower(r.URL.Query().Get("user"))

	db, erro := database.ConnectToDatabase()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repository.NewUserRepository(db)

	users, erro := repository.GetUsers(nameOrAlias)

	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)

	userID := parameter["id"]
	parsedUserID, err := strconv.ParseInt(userID, 10, 64)

	if err != nil {
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

	user, err := repository.GetUserByID(parsedUserID)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)

	userID := parameter["id"]
	parsedUserID, err := strconv.ParseInt(userID, 10, 64)

	println(userID, parsedUserID)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	inputRequest, err := io.ReadAll(r.Body)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var userUpdate models.User

	if err = json.Unmarshal(inputRequest, &userUpdate); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = userUpdate.Prepare("update"); err != nil {
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

	if err = repository.UpdateUser(parsedUserID, userUpdate); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)

	userID := parameter["id"]
	parsedUserID, err := strconv.ParseInt(userID, 10, 64)

	if err != nil {
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

	if err = repository.DeleteUser(parsedUserID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}