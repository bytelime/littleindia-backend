package config 

import (
	"database/sql"
	"os"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

// DB function
func DB() *sql.DB {

	user := os.Getenv("RDS_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	_db := os.Getenv("DB")

	log.Print(user)

	db, _ := sql.Open("mysql", user+":"+password+"@tcp("+host+":" + port +")/"+_db)
	err := db.Ping()

	if err != nil {
		log.Panic(DBerrorDescriptions["open"])
	}

	return db
}
