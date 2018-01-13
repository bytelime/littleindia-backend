package persistence

import (
  CO "server/config"
  . "server/models"
  "github.com/go-sql-driver/mysql"
)

type SubCategoryPersistenceManager interface {
  AddSubcategory(*SubCategory) (int, error)
  RemoveSubcategory(*SubCategory) error
  GetSubcatsByCat(*Category) ([]*SubCategory, error)
}

func AddSubcategory(s *SubCategory) (int, error){
  db := DB()
  defer db.Close()

  stmt, _ := db.Prepare("INSERT INTO Subcategories(name) VALUES(?)")
  res , err := stmt.Exec(s.Name)

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
  s.Id = int(lastId)

  return s.Id, err
}

func RemoveSubcategory(s *SubCategory) error{
  db := DB()
  defer db.Close()

  _ , err := db.Exec("DELETE FROM Subategories WHERE id=?", s.Id)

  return err
}


func GetSubcatsByCat(c *Category) ([]*SubCategory, error){
  db := DB()
  defer db.Close()

  var sList []*SubCategory

  var (
    id int
    name string
  )

  rows, err := db.Query("SELECT s.id, s.name FROM Subcategories s JOIN Categories c ON s.categoryId = c.id WHERE c.name = ? ORDER BY s.name DESC", c.Name)

  if err != nil {
    return sList, err
  }

  defer rows.Close()
  for rows.Next() {
    errScan := rows.Scan(&id, &name)

    if errScan != nil {
      return sList, errScan
    }

    s := NewSubCategory(name, *c)
    s.Id = id
    sList = append(sList, s)
  }

  rowErr := rows.Err()

  if err != nil {
    return sList, rowErr
  }

  return sList, nil
}
