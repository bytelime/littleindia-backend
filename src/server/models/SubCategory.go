package models

type SubCategory struct {
  Id int
  Name string
  Category Category
}

func NewSubCategory(name string, cat Category) *SubCategory {

  s := SubCategory{-1, name, cat}
  return &s
}
