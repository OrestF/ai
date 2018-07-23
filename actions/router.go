package actions

import (
	"github.com/gobuffalo/buffalo"
)

func ConfigRoutes(app *buffalo.App) {
	app.GET("/", HomeHandler)

	// USERS
	app.GET("/users", UsersIndex)
	app.GET("/users/{id}", UsersShow)
	app.GET("/users/{id}/edit", UsersEdit)
	app.POST("/users", UsersCreate)
	app.POST("/users/{id}/edit", UsersUpdate)
	app.DELETE("/users/{id}", UsersDelete)

	app.ServeFiles("/", assetsBox) // serve files from the public directory
}
