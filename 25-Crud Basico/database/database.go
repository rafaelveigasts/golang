package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // sempre colocar esse import na mão
)

// ConnectToDatabase is a function that connects to the database
func ConnectToDatabase() (*sql.DB, error) {
		dbUser := "golang"
    dbPass := "golang"
    dbName := "golang"
    dbHost := "localhost"
    dbPort := "3306"

    urlConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
        dbUser, dbPass, dbHost, dbPort, dbName)

				println(urlConnection)


	db, erro := sql.Open("mysql", urlConnection)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	println("Connected to database", db)

	return db, nil
}

