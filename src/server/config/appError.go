package config

import (
	"log"
	"time"
	"os"
	"path/filepath"
	"fmt"

)

type AppError struct {
    Code string
	Description string
}

var (
  Log *log.Logger
)

var DBerrorDescriptions = map[string]AppError{
  "open": {"100","Could not ping database."},
  "alreadyExists": {"101","Already exists in database."},
  "doesntExist": {"102","Doesnt exist in database."},

}


func init() {
	
	logpath , _ := filepath.Abs(filepath.Dir(os.Args[0]) + "/logs/errors.log")

	var file, _ = os.Create(logpath)

	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
}

func (e *AppError) Error() string {  
    return fmt.Sprintf("%d - %s", e.Code, e.Description)
}


func ThrowError(key string) *AppError{

	err := new(AppError)
	t := time.Now()

	err.Code = DBerrorDescriptions[key].Code
	err.Description = DBerrorDescriptions[key].Description

	Log.Println(t.Format("2006-01-02 15:04:05") + " -> (" + err.Code + ")" + err.Description)

	return err

}


//formating
//	s := fmt.Sprintf("%s %s %s %s %s", user, password, host, port, _db)
//   log.Printf(s)