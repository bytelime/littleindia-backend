package services

import (
  DB "server/persistence"
  . "server/models"
)

type ProductService interface {
  AddProduct(*Product) (int, error)
  UpdateProduct(*Product, string) error
  GetProductByName(name string) (*Product, error)
  GetAllProducts() ([]*Product, error)
  RemoveProduct(*Product) error
}

func AddProduct(p *Product) (int, error){
  return DB.AddProduct(p)
}


func GetProductByName(name string) (*Product, error){
  return DB.GetProductByName(name)
}


func UpdateProduct(p *Product, s string) error{
  return nil
}



func GetAllProducts() ([]*Product, error){
  return DB.GetAllProducts()
}


func RemoveProduct(p *Product) error{
  return DB.RemoveProduct(p)
}

