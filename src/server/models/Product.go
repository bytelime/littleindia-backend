package models

type Product struct {
    Id int
    Name string
    PhotoUrl string
    Description string
    Price int
    HasStock int
    CategoryId int
    SubCategoryId int
}

func NewProduct(name string, photoUrl string, description string, price int, hasStock int, categoryId int, subCategoryId int) *Product {

    p := Product{-1, name, photoUrl, description, price, hasStock, categoryId, subCategoryId}
    return &p
}
