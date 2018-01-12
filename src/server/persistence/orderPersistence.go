package persistence

import (
  CO "server/config"
  . "server/models"
  "github.com/go-sql-driver/mysql"
  "database/sql"
)

type OrderPersistenceManager interface {
  AddOrder(*Order) (int, error)
  UpdateOrder(*Order, string) error
  RemovePartOrder(*Order) error
  RemoveFullOrder(*Order) error
  GetOrderById(int) (*Order, error)
  GetOrdersByCustomerId(int) (*Order, error)
  GetAllOrders(int) ([]*Order, error)
}

func AddOrder(o *Order) (int, error){
  db:= DB()
  defer db.Close()

  stmt, _ := db.Prepare("INSERT INTO Orders(id, productId, customerId) VALUES(?, ?)")

  res, err := stmt.Exec(o.Id, o.Product.Id, o.Customer.Id)

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

func UpdateOrder(){}

func RemovePartOrder(o *Order) error {
  db:= DB()
  defer db.Close()

  stmt, _ := db.Prepare("DELETE FROM Orders WHERE id = ? AND productId = ?")
  _ , err := stmt.Exec(o.Id, o.Product.Id)

  return err
}

func RemoveFullOrder(o *Order) error{
  db:= DB()
  defer db.Close()

  stmt, _ := db.Prepare("DELETE FROM Orders WHERE id = ?")

  _ , err := stmt.Exec(o.Id)

  return err
}

func GetAllOrders() ([]*Order, error){
  return GetOrders(0, 0)
}

func GetOrdersById(oid int) ([]*Order, error){
  if oid <= 0 {
    return nil, nil //TODO: ERROR HANDLING
  }
  return GetOrders(oid, 0)
}

func GetOrdersByCustomerId(cid int) ([]*Order, error){
  if cid <= 0 {
    return nil, nil //TODO: ERROR HANDLING
  }
  return GetOrders(0, cid)
}

func GetOrders(oid int, cid int) ([]*Order, error){
  db := DB()
  defer db.Close()

  var oList []*Order
  var rows *sql.Rows
  var err error
  var (
    id int
    productId int
    pName string
    pUrl string
    pDesc string
    pPrice int
    pStock int
    pCat int
    pSub int
    customerId int
    cName string
    cPhone string
    cMail string
    query string
  )
  query = "SELECT o.id, " +
  "c.id, c.name, c.phone, c.email, " +
    "p.id, p.name, p.photoUrl, p.description, p.price, p.hasStock, p.categoryId, p.subCategoryId " +
    "FROM Orders o JOIN Customers c ON o.customerId = c.id " +
    "JOIN Products p ON o.productId = p.id"

  if oid > 0 {
    query += " WHERE o.id = ?"
    rows,err = db.Query(query,oid)
  } else if cid > 0 {
    query += " WHERE c.id = ? ORDER BY o.id DESC"
    rows,err = db.Query(query,cid)
  } else {
    query += " ORDER BY o.id DESC"
    rows, err = db.Query(query)
  }

  if err != nil {
    return oList, err
  }

  defer rows.Close()
  for rows.Next() {
    errScan := rows.Scan(&id, &customerId, &cName, &cPhone, &cMail, &productId, &pName, &pUrl, &pDesc, &pPrice, &pStock, &pCat, &pSub)

    if errScan != nil {
      return oList, errScan
    }
    prod := NewProduct(pName, pUrl, pDesc, pPrice, pSub, pCat, pSub)
    prod.Id = productId
    cust := NewCustomer(cName, cMail, cPhone)
    cust.Id = customerId
    order:= NewOrder(id, *cust, *prod)


    oList = append(oList, order)
  }

  rowErr := rows.Err()

  if err != nil {
    return oList, rowErr
  }

  return oList, nil

}


