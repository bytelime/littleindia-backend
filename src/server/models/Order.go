package models

type Order struct {
  Id int
  Customer Customer
  Product Product
}

func NewOrder(id int, c Customer, p Product) *Order {

  o := Order{id, c, p}
  return &o
}
