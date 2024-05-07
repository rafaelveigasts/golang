package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	 dbUser := "golang"
    dbPass := "golang"
    dbName := "golang"
    dbHost := "localhost"
    dbPort := "3306"

    urlConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
        dbUser, dbPass, dbHost, dbPort, dbName)


	db, erro := sql.Open("mysql", urlConnection)
	if erro != nil {
		panic(erro.Error())
	}
	defer db.Close()

	if erro = db.Ping(); erro != nil {
		panic(erro.Error())
	}

	println("Connected to database", db)

	linhas, erro := db.Query("show databases")
	if erro != nil {
		panic(erro.Error())
	}

	defer linhas.Close()

	fmt.Println(linhas)
}

// escreva uma linha de comando que cria um container com o banco de dados MySQL a partir dessa string de conexão "golang:golang@/golang?charset=utf8&parseTime=True&loc=Local"

// docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=golang -e MYSQL_USER=golang -e MYSQL_PASSWORD=golang -d mysql:5.7