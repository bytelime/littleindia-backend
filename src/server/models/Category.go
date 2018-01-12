package models

type Category struct {
	Id int
    Name string 
}

func NewCategory(name string) *Category {  
    
    c := Category{-1, name}
    return &c
}
