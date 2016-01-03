package controllers

import (
	"github.com/revel/revel"
	"regexp"
	"revel_cms/app/models"
)

type App struct {
	*revel.Controller
}

func (c App) Index(q string) revel.Result {
	c.Validation.Clear()
	posts := new(models.Post).GetPosts()

	if t, _ := regexp.MatchString("^[a-zA-Z]{2,}$", q); t {
		posts = new(models.Post).FindByName(q)
	}

	if user := c.connected(); user != nil {
		return c.Render(posts, user)
	}
	return c.Render(posts)
}
