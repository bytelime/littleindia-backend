package config 

import (
	"database/sql"
	"os"

	// for mysql
	_ "github.com/go-sql-driver/mysql"
)

// DB function
func DB() *sql.DB {

	user := os.Getenv("RDS_USERNAME")
	password := os.Getenv("RDS_PASSWORD")
	host := os.Getenv("RDS_HOSTNAME")
	port := os.Getenv("RDS_PORT")
	_db := os.Getenv("RDS_DB_NAME")

	db, _ := sql.Open("mysql", user+":"+password+"@tcp("+host+":" + port +")/"+_db)
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
