package routes

import (
	"github.com/kataras/iris"
	DB "server/persistence"
	CO "server/config"
)

func Call(ctx iris.Context) {

	var msg string

	err := DB.AddCategory("Prueba")

	if (err == nil){
		msg = "OK"
	} else {
		msg = err.(*CO.AppError).Description
	}

	//DB.RemoveCategory("Prueba")

	json(ctx, map[string]interface{}{
		"mensaje": msg,
	})

}