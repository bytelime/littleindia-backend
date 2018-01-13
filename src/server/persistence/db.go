package persistence 

import (
	"database/sql"
	"os"
	"log"
	_ "github.com/go-sql-driver/mysql"
	CO "server/config"
)

// DB function
func DB() *sql.DB {

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	_db := os.Getenv("DB")

	log.Print(user, password, host)

	db, _ := sql.Open("mysql", user+":"+password+"@tcp("+host+":" + port +")/"+_db)
	err := db.Ping()

	if err != nil {
		log.Panic(CO.ThrowError("open"))
	}

	return db
}
