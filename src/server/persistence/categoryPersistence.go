package persistence

import (
	CO "server/config"
	"errors"
	"log"
)

func AddCategory(name string) error{

	log.Printf("Entro al add")
	db := CO.DB()
	log.Printf("Obtengo db")
	defer db.Close()

	var categoryCount int

	log.Printf("Checkeo cantidad")
	db.QueryRow("SELECT COUNT(id) AS categoryCount FROM Categories WHERE name=?", name).Scan(&categoryCount)

	if (categoryCount > 0){
		return errors.New(CO.DBerrorDescriptions["alreadyExists"])
	}

	log.Printf("En teoria no existe, asi que preparo el insert")

	stmt, _ := db.Prepare("INSERT INTO Categories(name) VALUES(?)")
	log.Printf("Listo para ejecutar")

	_, err := stmt.Exec(name)
	log.Printf("Ejecuta3")
	
	return err

}

func RemoveCategory(name string) error{

	db := CO.DB()
	defer db.Close()

	_, err := db.Exec("DELETE FROM Categories WHERE name=?", name)
	
	return err

}
