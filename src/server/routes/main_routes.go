package routes

import (
	DB "server/persistence"
	CO "server/config"
	"github.com/kataras/iris"
)

// Index route. No anda la ruta del coso, ni carga los ENV.
func Index(ctx iris.Context) {
	//loggedIn(ctx, "/welcome")

	CO.Err(nil)
	catList, _ := DB.GetAllCategories()
	prodList, _ := DB.GetAllProducts()

	//var emptyArray []*M.Product

	renderTemplate(ctx, "index", iris.Map{
		"title":   "Home",
		"session": "session",
		"products":   prodList,
		"categories": catList,
		"GET":     "xD",
	})
}

// Welcome route
func Welcome(ctx iris.Context) {
	notLoggedIn(ctx)

	renderTemplate(ctx, "welcome", iris.Map{
		"title": "Welcome",
	})
}

// NotFound route
func NotFound(ctx iris.Context) {
	renderTemplate(ctx, "404", iris.Map{
		"title":   "Oops!! Error",
		"session": ses(ctx),
	})
}

func Error(ctx iris.Context) {
	renderTemplate(ctx, "500", iris.Map{
		"title":   "Oops!! Error",
		"session": ses(ctx),
	})
}
