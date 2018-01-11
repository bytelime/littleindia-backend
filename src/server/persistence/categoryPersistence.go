package persistence

import (
	CO "server/config"
	"log"
  "github.com/go-sql-driver/mysql"
)

func AddCategory(name string) error {

	db := DB()
	defer db.Close()

	stmt, _ := db.Prepare("INSERT INTO Categories(name) VALUES(?)")
	res , err := stmt.Exec(name)

	if err != nil {
		
		me, ok := err.(*mysql.MySQLError)
		
		if !ok {
      		return err
    	}

    	if me.Number == 1062 { //duplicate key
      		return CO.ThrowError("alreadyExists")
    	}

    	return err
	}

	lastId, err := res.LastInsertId()
	log.Printf("%d", lastId)

	return err

}

func RemoveCategory(name string) error{

	db := DB()
	defer db.Close()

	_ , err := db.Exec("DELETE FROM Categories WHERE name=?", name)

	return err

}
