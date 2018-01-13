package services

import (
  . "server/models"
  DB "server/persistence"
)
type CategoryService interface {
  AddCategory(*Category) (int, error)
  UpdateCategory(*Category, string) error
  GetCategoryByName(name string) (*Category, error)
  GetAllCategories() ([]*Category, error)
  RemoveCategory(*Category) error
}

func AddCategory(c *Category) (int, error){
  return DB.AddCategory(c)
}

func UpdateCategory(c *Category, name string) error{
  return DB.UpdateCategory(c, name)
}

func GetCategoryByName(name string) (*Category, error){
  return DB.GetCategoryByName(name)
}

func GetAllCategories() ([]*Category, error){
  return DB.GetAllCategories()
}

func RemoveCategory(c *Category) error{
  return DB.RemoveCategory(c)
}
