package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/projects/ai/models"
)

func UsersIndex(c buffalo.Context) error {
	users := []models.User{}
	models.DB.All(&users)
	return c.Render(200, r.Auto(c, users))
}

func UsersShow(c buffalo.Context) error {
	user := models.User{}
	models.DB.Find(&user, c.Param("id"))
	return c.Render(200, r.Auto(c, user))
}

func UsersCreate(c buffalo.Context) error {
	user := models.User{}
	c.Bind(&user)
	models.DB.Create(&user)
	return c.Render(200, r.Auto(c, user))
}

func UsersDelete(c buffalo.Context) error {
	user := models.User{}
	models.DB.Find(&user, c.Param("id"))
	models.DB.Destroy(&user)
	return c.Render(200, r.Auto(c, user))
}
