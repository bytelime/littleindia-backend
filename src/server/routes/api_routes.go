package routes

import (
	"github.com/kataras/iris"
	DB "server/persistence"
	"fmt"
	"log"
)

func Call(ctx iris.Context) {

	var msg string

	log.Printf("Entro al call")

	err := DB.AddCategory("Prueba")

	log.Printf("Pase el addCategory")

	if (err != nil){
		msg = "OK"
	} else {
		msg = err.Error()
	}

	s := fmt.Sprintf("Hola, el mensaje es %s", msg)
    log.Printf(s)

	json(ctx, map[string]interface{}{
		"mensaje": msg,
	})

}