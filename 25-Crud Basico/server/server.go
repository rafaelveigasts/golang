package server

import (
	"crud_basico/database"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"
)

type user struct {
	ID   int32    `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}


// CreateUser is a function that creates a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, erro := io.ReadAll(r.Body)

	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error while reading the request body"))
		return
	}

	var user user 

	if erro = json.Unmarshal(body, &user); erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error while parsing the request body"))
		return
	}

	w.WriteHeader(http.StatusCreated)

	db, erro := database.ConnectToDatabase()
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error while connecting to the database"))
		return
	}

	defer db.Close()

	//prepare statement

	statement, erro := db.Prepare("insert into users (name, email) values (?, ?)")
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error while creating the statement: %s", erro.Error())))

		return
	}

	defer statement.Close()

	//execute statement
	insert, erro := statement.Exec(user.Name, user.Email)
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error while executing the statement: %s", erro.Error())))
		return
	}

	id, erro := insert.LastInsertId()
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error while getting the last insert id: %s", erro.Error())))
		return
	}

	user.ID = int32(id)

	w.Write([]byte("User created successfully. ID: %d" + string(user.Name)))


}

// GetUsers is a function that returns all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, erro := database.ConnectToDatabase()
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error while connecting to the database"))
		return
	}

	defer db.Close()

	rows, erro := db.Query("select * from users")
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error while executing the query: %s", erro.Error())))
		return
	}

	defer rows.Close()

	var users []user

	for rows.Next() {
		var user user
		if erro = rows.Scan(&user.ID, &user.Name, &user.Email); erro != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error while scanning the rows: %s", erro.Error())))
			return
		}

		users = append(users, user)
	}

	if erro := json.NewEncoder(w).Encode(users); erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error while parsing the users to JSON: %s", erro.Error())))
		return
	}

}

// GetUser is a function that returns a user
func GetUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	ID, erro:= strconv.ParseUint(parameters["id"], 10, 32)
	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error while parsing the id"))
		return
	}

	db, erro := database.ConnectToDatabase()
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error while connecting to the database"))
		return
	}

	defer db.Close()

	rows, erro := db.Query("select * from users where id = ?", ID)
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error while executing the query: %s", erro.Error())))
		return
	}

	defer rows.Close()

	var user user

	if rows.Next() {
		if erro = rows.Scan(&user.ID, &user.Name, &user.Email); erro != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error while scanning the rows: %s", erro.Error())))
			return
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	userJSON, erro := json.Marshal(user)
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error while parsing the user to JSON: %s", erro.Error())))
		return
	}

	w.Write(userJSON)
}

// UpdateUser is a function that updates a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	ID, erro:= strconv.ParseUint(parameters["id"], 10, 32)
	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error while parsing the id"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error while reading the request body"))
		return
	}

	var user user
	if erro = json.Unmarshal(body, &user); erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error while parsing the request body"))
		return
	}

	db, erro := database.ConnectToDatabase()
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error while connecting to the database"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("update users set name = ?, email = ? where id = ?")

	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error while creating the statement: %s", erro.Error())))
		return
	}

	defer statement.Close()

	_, erro = statement.Exec(user.Name, user.Email, ID)
	w.WriteHeader(http.StatusNoContent)

	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error while executing the statement: %s", erro.Error())))
		return
	}

	w.Write([]byte("User updated successfully"))
}

// DeleteUser is a function that deletes a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	ID, erro:= strconv.ParseUint(parameters["id"], 10, 32)
	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error while parsing the id"))
		return
	}

	db, erro := database.ConnectToDatabase()
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error while connecting to the database"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("delete from users where id = ?")

	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error while creating the statement: %s", erro.Error())))
		return
	}

	defer statement.Close()

	_, erro = statement.Exec(ID)
	w.WriteHeader(http.StatusNoContent)

	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error while executing the statement: %s", erro.Error())))
		return
	}

	w.Write([]byte("User deleted successfully"))
}