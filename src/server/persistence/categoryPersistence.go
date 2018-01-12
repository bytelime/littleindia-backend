package persistence

import (
	CO "server/config"
	. "server/models"
	"github.com/go-sql-driver/mysql"
)

type CategoryPersistenceManager interface {
    AddCategory(*Category) (int, error)
    UpdateCategory(*Category, string) error
    GetCategoryByName(name string) (*Category, error)
    GetAllCategories() ([]*Category, error)
    RemoveCategory(*Category) error 
}

//agregar el modelo aca!

func AddCategory(category *Category) (int,error) {

	db := DB()
	defer db.Close()

	stmt, _ := db.Prepare("INSERT INTO Categories(name) VALUES(?)")
	res , err := stmt.Exec(category.Name)

	if err != nil {
		
		me, ok := err.(*mysql.MySQLError)
		
		if !ok {
      		return -1, err
    	}

    	if me.Number == 1062 { //duplicate key
      		return -1, CO.ThrowError("alreadyExists")
    	}

    	return -1, err
	}

	lastId, err := res.LastInsertId()
	category.Id = int(lastId)

	return category.Id, err

}

func UpdateCategory(category *Category, newName string) error{

	db := DB()
	defer db.Close()

	_ , err := db.Exec("UPDATE Categories SET name=? WHERE id=?", newName, category.Id)

	return err

}


func GetCategoryByName(catName string) (*Category, error){

	db := DB()
	defer db.Close()

	var cat *Category

	var (
		id int
		name string
	)

	row := db.QueryRow("SELECT * FROM Categories WHERE name=?", catName)
	err := row.Scan(&id, &name)

	if err != nil {

		me, ok := err.(*mysql.MySQLError)
		
		if !ok {
      		return cat, err
    	}

    	if me.Number == 1032 { //key not found?
      		return cat, CO.ThrowError("doesntExist")
    	}

    	return cat, err
	}

	cat = NewCategory(name)
	cat.Id = id	

	return cat, nil

}


func GetAllCategories() ([]*Category, error){

	db := DB()
	defer db.Close()

	var catList []*Category

	var (
		id int
		name string
	)

	rows, err := db.Query("SELECT * FROM Categories ORDER BY name DESC")

	if err != nil {
    	return catList, err
	}

	defer rows.Close()
	for rows.Next() {
		errScan := rows.Scan(&id, &name)

		if (errScan != nil){
			return catList, errScan
		}

		cat := NewCategory(name)
		cat.Id = id
		catList = append(catList, cat)
	}

	rowErr := rows.Err()

	if (err != nil){
		return catList, rowErr
	}

	return catList, nil

}

func RemoveCategory(category *Category) error{

	db := DB()
	defer db.Close()

	_ , err := db.Exec("DELETE FROM Categories WHERE id=?", category.Id)

	return err

}
