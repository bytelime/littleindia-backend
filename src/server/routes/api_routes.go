package routes

import (
	"github.com/kataras/iris"
	CO "server/config"
)

func Call(ctx iris.Context) {

	db := CO.DB()

	stmt, _ := db.Prepare("INSERT INTO posts(title, content, createdBy, createdAt) VALUES (?, ?, ?, ?)")
	rs, _ := stmt.Exec("title", "content", "xD", "123")

	json(ctx, map[string]interface{}{
		"mensaje": rs,
	})
}