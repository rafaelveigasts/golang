package main

import (
	"api-devbook/src/config"
	"api-devbook/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()
	println("rodando na porta", config.Port)

	r := router.GenerateRouter()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))


}