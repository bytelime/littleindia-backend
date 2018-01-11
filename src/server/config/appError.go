package config

import (
	"log"
	"time"
	"os"
)

type appError struct {
    code string
}

var (
  Log *log.Logger
)

var DBerrorDescriptions = map[string]string{
  "open": "100",
  "alreadyExists": "101",
}

var DBErrorsCodeToDescription = map[string]string{
  "100": "Could not ping database.",
  "101": "Already exists in database.",
}

func init() {
	
	var logpath = os.ExpandEnv("$GOPATH/src/server/logs/errors.log")

	var file, err1 = os.Create(logpath)
	
	if err1 != nil {
		panic(err1)
	}

	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
}

func (e *appError) Error() string {  
    return e.code
}


func ThrowError(key string) *appError{

	err := new(appError)
	t := time.Now()

	err.code = DBerrorDescriptions[key]

	Log.Println(t.Format("2006-01-02 15:04:05") + " -> (" + err.code + ")" + DBErrorsCodeToDescription[err.code])

	return err

}

func GetErrorDescription (e error) string{
	return DBErrorsCodeToDescription[e.Error()]
}

//formating
//	s := fmt.Sprintf("%s %s %s %s %s", user, password, host, port, _db)
//   log.Printf(s)