package config

import (
	"log"
	"time"
	"os"
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
}


func init() {
	
	var logpath = os.ExpandEnv("$GOPATH/src/server/logs/errors.log")

	var file, err1 = os.Create(logpath)
	
	if err1 != nil {
		panic(err1)
	}

	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
}

func (e *AppError) Error() string {  
    return e.Code
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