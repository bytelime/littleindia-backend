package routes

import (
	CO "server/config"
	"github.com/kataras/iris"
)

// Index route
func Index(ctx iris.Context) {
	loggedIn(ctx, "/welcome")

	id, _ := CO.AllSessions(ctx)
	db := CO.DB()
	var (
		postID    int
		title     string
		content   string
		createdBy int
		createdAt string
	)
	feeds := []interface{}{}

	stmt, _ := db.Prepare("SELECT posts.postID, posts.title, posts.content, posts.createdBy, posts.createdAt from posts, follow WHERE follow.followBy=? AND follow.followTo = posts.createdBy ORDER BY posts.postID DESC")
	rows, qErr := stmt.Query(id)
	CO.Err(qErr)

	for rows.Next() {
		rows.Scan(&postID, &title, &content, &createdBy, &createdAt)
		feed := map[string]interface{}{
			"postID":    postID,
			"title":     title,
			"content":   content,
			"createdBy": createdBy,
			"createdAt": createdAt,
		}
		feeds = append(feeds, feed)
	}

	renderTemplate(ctx, "index", iris.Map{
		"title":   "Home",
		"session": ses(ctx),
		"posts":   feeds,
		"GET":     CO.Get,
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
