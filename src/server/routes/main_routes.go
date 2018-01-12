package routes

import (
	DB "server/persistence"
	M "server/models"
	"github.com/kataras/iris"
)

// Index route
func Index(ctx iris.Context) {
	//loggedIn(ctx, "/welcome")

	catList, _ := DB.GetAllCategories()
	var emptyArray []*M.Product

	renderTemplate(ctx, "index", iris.Map{
		"title":   "Home",
		"session": "session",
		"products":   emptyArray,
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
