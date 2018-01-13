package persistence

import (
	CO "server/config"
	. "server/models"
	"github.com/go-sql-driver/mysql"
)

type ProductPersistenceManager interface {
    AddProduct(*Product) (int, error)
    UpdateProduct(*Product, string) error
    GetProductByName(name string) (*Product, error)
    GetAllProducts() ([]*Product, error)
    RemoveProduct(*Product) error
}

//agregar el modelo aca!
//mejorar mas objetoso, mejorar calls y hacer llamado por categoria, y por subcat
//paginate

func AddProduct(product *Product) (int,error) {

	db := DB()
	defer db.Close()

	stmt, _ := db.Prepare("INSERT INTO Products(name, photoUrl, description, price, hasStock, categoryId, subCategoryId) VALUES(?, ?, ?, ?, ?, ?, ?)")
	res , err := stmt.Exec(product.Name, product.PhotoUrl, product.Description, product.Price, product.HasStock, product.CategoryId, product.SubCategoryId)

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
	product.Id = int(lastId)

	return product.Id, err

}

//no
func UpdateProduct(product *Product, newName string) error{

	db := DB()
	defer db.Close()

	_ , err := db.Exec("UPDATE Products SET name=? WHERE id=?", newName, product.Id)

	return err

}


func GetProductByName(prodName string) (*Product, error){

	db := DB()
	defer db.Close()

	var prod *Product

	var (
		id int
		name string
		photoUrl string
    	description string
    	price int
    	hasStock int
    	categoryId int
    	subCategoryId int
	)

	row := db.QueryRow("SELECT * FROM Products WHERE name=?", prodName)
	err := row.Scan(&id, &name, &photoUrl, &description, &price, &hasStock, &categoryId, &subCategoryId)

	if err != nil {

		me, ok := err.(*mysql.MySQLError)

		if !ok {
      		return prod, err
    	}

    	if me.Number == 1032 { //key not found?
      		return prod, CO.ThrowError("doesntExist")
    	}

    	return prod, err
	}

	prod = NewProduct(name, photoUrl, description, price, hasStock, categoryId, subCategoryId)
	prod.Id = id

	return prod, nil

}


func GetAllProducts() ([]*Product, error){

	db := DB()
	defer db.Close()

	var prodList []*Product

	var (
		id int
		name string
		photoUrl string
    	description string
    	price int
    	hasStock int
    	categoryId int
    	subCategoryId int
	)

	rows, err := db.Query("SELECT * FROM Products ORDER BY name DESC")

	if err != nil {
    	return prodList, err
	}

	defer rows.Close()
	for rows.Next() {
		errScan := rows.Scan(&id, &name, &photoUrl, &description, &price, &hasStock, &categoryId, &subCategoryId)

		if errScan != nil {
			return prodList, errScan
		}

		prod := NewProduct(name, photoUrl, description, price, hasStock, categoryId, subCategoryId)
		prod.Id = id

		prodList = append(prodList, prod)
	}

	rowErr := rows.Err()

	if err != nil {
		return prodList, rowErr
	}

	return prodList, nil

}

func RemoveProduct(product *Product) error{

	db := DB()
	defer db.Close()

	_ , err := db.Exec("DELETE FROM Products WHERE id=?", product.Id)

	return err

}
