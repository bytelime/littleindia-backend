package persistence

import (
  CO "server/config"
  . "server/models"
  "github.com/go-sql-driver/mysql"
)

type CustomerPersistenceManager interface {
  AddCustomer(*Customer) (int, error)
  UpdateCustomer(*Customer, string) error
  GetCustomerByOrderId(int) (*Customer, error)
  GetAllCustomers() ([]*Customer, error)
  RemoveCustomer(*Customer) error
}


func AddCustomer(c *Customer) (int, error){
  db:= DB()
  defer db.Close()

  stmt, _ := db.Prepare("INSERT INTO Customers(name, email, phone) VALUES(?, ?, ?)")

  res, err := stmt.Exec(c.Name, c.Email, c.Phone)

  if err != nil {

    me, ok := err.(*mysql.MySQLError)

    if !ok {
      return -1, err
    }

    if me.Number == 1062 { //duplicate key
      return -1, CO.ThrowError("alreadyExists")
    }

    return -1, err
  }

  lastId, err := res.LastInsertId()

  return int(lastId), err

}

func UpdateCustomer(c* Customer){
  //pensar como hacer bien esto
}

func RemoveCustomer(c* Customer) error {
  db:= DB()
  defer db.Close()

  stmt, _ := db.Prepare("DELETE FROM Customers WHERE name = ? AND email = ? AND phone = ?")

  _ , err := stmt.Exec(c.Name, c.Email, c.Phone)

  return err
}

func GetCustomerByOrderId(orderId int) (*Customer, error){
  db:=DB()
  defer db.Close()
  var customer *Customer
  var (
    id int
    name string
    email string
    phone string
  )


  row := db.QueryRow("SELECT DISTINCT c.id, c.name, c.email, c.phone FROM Customers c JOIN Orders o ON c.id = o.customerId WHERE o.id = ?", orderId)
  err := row.Scan(&id, &name, &email, &phone)

  if err != nil {
    _, ok := err.(*mysql.MySQLError)
    if !ok {
      return customer, err
    }
  }

  customer = NewCustomer(name, email, phone)
  customer.Id = id

  return customer, nil

}

func GetAllCustomers() ([]*Customer, error){
  db := DB()
  defer db.Close()

  var cList []*Customer

  var (
    id int
    name string
    phone string
    mail string
  )

  rows, err := db.Query("SELECT * FROM Customers ORDER BY name DESC")

  if err != nil {
    return cList, err
  }

  defer rows.Close()
  for rows.Next() {
    errScan := rows.Scan(&id, &name, &mail, &phone)

    if errScan != nil {
      return cList, errScan
    }

    cus := NewCustomer(name, mail, phone)
    cus.Id = id

    cList = append(cList, cus)
  }

  rowErr := rows.Err()

  if (err != nil){
    return cList, rowErr
  }

  return cList, nil
}
