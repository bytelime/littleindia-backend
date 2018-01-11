package persistence

import (
	CO "server/config"
)

func AddCategory(name string) error {

	db := DB()
	defer db.Close()

	var categoryCount int

	db.QueryRow("SELECT COUNT(id) AS categoryCount FROM Categories WHERE name=?", name).Scan(&categoryCount)

	if (categoryCount > 0){
		return CO.ThrowError("alreadyExists")
	}

	stmt, _ := db.Prepare("INSERT INTO Categories(name) VALUES(?)")
	_, err := stmt.Exec(name)
	
	return err

}

func RemoveCategory(name string) error{

	db := DB()
	defer db.Close()

	_, err := db.Exec("DELETE FROM Categories WHERE name=?", name)
	
	return err

}
