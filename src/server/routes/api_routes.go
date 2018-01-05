package routes

import (
	"github.com/kataras/iris"
)

func Call(ctx iris.Context) {

	json(ctx, map[string]interface{}{
		"mensaje": "Hola bebeto",
	})
}