package services

import (
  DB "server/persistence"
  . "server/models"
)

type OrderService interface {
  AddOrder(*Order) (int, error)
  UpdateOrder(*Order, string) error
  RemovePartOrder(*Order) error
  RemoveFullOrder(*Order) error
  GetOrdersById(int) ([]*Order, error)
  GetOrdersByCustomerId(int) ([]*Order, error)
  GetAllOrders() ([]*Order, error)
}


func AddOrder(o *Order) (int, error){
  return DB.AddOrder(o)
}

func UpdateOrder(*Order, string) error{
  return nil
}

func RemovePartOrder(o *Order) error{
  return DB.RemovePartOrder(o)
}

func RemoveFullOrder(o *Order) error {
  return DB.RemoveFullOrder(o)
}

func GetOrderById(id int) ([]*Order, error) {
  return DB.GetOrdersById(id)
}

func GetOrdersByCustomerId(id int) ([]*Order, error){
  return DB.GetOrdersByCustomerId(id)
}

func GetAllOrders(id int) ([]*Order, error){
  return DB.GetAllOrders()
}
