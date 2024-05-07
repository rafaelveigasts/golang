package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDatabase() (*sql.DB, error) {

	// fmt.Println(config.UrlConnection)
	// db, erro := sql.Open("mysql", config.UrlConnection )

	// print("Connected to database", db)

	// if erro != nil {
	// 	return nil, erro
	// }

	// if erro = db.Ping(); erro != nil {
	// 	db.Close()
	// 	return nil, erro
	// }

	// 	return db, nil

//___________________________________________
			// dbUser := "golang"
			// dbPass := "golang"
			// dbName := "golang"
			// dbHost := "localhost"
			// dbPort := "3306"
		dbUser := os.Getenv("Db_User")
    dbPass := os.Getenv("Db_Password")
    dbName := os.Getenv(("Db_Name"))
    dbHost := os.Getenv("Db_Host")
    dbPort := os.Getenv("Db_Port")

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