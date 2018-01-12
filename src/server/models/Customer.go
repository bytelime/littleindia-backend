package models

type Customer struct {
  Id int
  Name string
  Email string
  Phone string
}

func NewCustomer(name string, email string, phone string ) *Customer {

  c := Customer{-1, name, email, phone}
  return &c
}
