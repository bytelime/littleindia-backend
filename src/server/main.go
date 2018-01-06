package main

import (
	R "server/routes"
	"github.com/kataras/iris"
	//"github.com/kataras/iris/core/host"
	"github.com/kataras/iris/middleware/recover"
	//"net/url"
	//"os"
)

func main() {
	app := iris.New()
	app.Use(recover.New())

	app.RegisterView(iris.HTML("./views", ".html"))
	app.StaticWeb("/", "./public")

	app.Get("/", R.Index)
	app.Get("/welcome", R.Welcome)
	app.Get("/404", R.NotFound)
	
	//app.Get("/profile/:id", R.Profile)
	//app.Get("/profile", R.NotFound)

	api := app.Party("/api")
	{
		api.Post("/call", R.Call)
	}

	/* El redirect no anda
	//80 el port en PROD
	target, _ := url.Parse("https://localhost:7668")
	go host.NewProxy("localhost:7667", target).ListenAndServe()

	//443 el port en PROD
	app.Run(iris.TLS("localhost:7668", "certs/mycert.cert", "certs/mykey.key"))
	*/

	app.Run(iris.Addr("0.0.0.0:8080"))

}
