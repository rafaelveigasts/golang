package main

import (
	"crud_basico/server"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter() 
	router.HandleFunc("/users", server.CreateUser).Methods(http.MethodPost) // rota, função, método http
	router.HandleFunc("/users", server.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", server.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", server.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/user/{id}", server.DeleteUser).Methods(http.MethodDelete)

	println("Server running at port 8001")
	log.Fatal(http.ListenAndServe(":8001", router))
}