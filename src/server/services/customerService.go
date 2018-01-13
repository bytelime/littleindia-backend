package services

import (
  . "server/models"
  DB "server/persistence"
)

type CustomerService interface {
  AddCustomer(*Customer) (int, error)
  UpdateCustomer(*Customer, string) error
  GetCustomerByOrderId(int) (*Customer, error)
  GetAllCustomers() ([]*Customer, error)
  RemoveCustomer(*Customer) error
}

func AddCustomer(c *Customer) (int, error){
  return DB.AddCustomer(c)
}

func UpdateCustomer(*Customer, string) error{
  return nil
}

func GetCustomerByOrderId(id int) (*Customer, error){
  return DB.GetCustomerByOrderId(id)
}

func GetAllCustomers() ([]*Customer, error){
  return DB.GetAllCustomers()
}

func RemoveCustomer(c *Customer) error{
  return DB.RemoveCustomer(c)
}

