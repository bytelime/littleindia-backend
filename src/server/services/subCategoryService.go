package services

import (
  DB "server/persistence"
  . "server/models"
)
type SubCategoryPersistenceManager interface {
  AddSubcategory(*SubCategory) (int, error)
  RemoveSubcategory(*SubCategory) error
  GetSubcatsByCat(*Category) ([]*SubCategory, error)
}

func AddSubcategory(s *SubCategory) (int, error){
  return DB.AddSubcategory(s)
}

func RemoveSubcategory(s *SubCategory) error{
  return DB.RemoveSubcategory(s)
}

func GetSubcatsByCat(c *Category) ([]*SubCategory, error){
  return DB.GetSubcatsByCat(c)
}
