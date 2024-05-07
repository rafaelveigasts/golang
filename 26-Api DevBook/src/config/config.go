package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Server, DbUser, DbPassword, DbPort, DbHost, DbName string
	
	Port = 0
	UrlConnection string
)

func LoadConfig(){
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Port, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Port = 9000
	}

	db_port := os.Getenv("Db_Port")
	
println(os.Getenv(db_port))	

    UrlConnection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("Db_User"),
		os.Getenv("Db_Password"),
		os.Getenv("Db_Name"),
		os.Getenv("Db_Host"),
		os.Getenv("Db_Port"),
	)

	println(UrlConnection)


} 
