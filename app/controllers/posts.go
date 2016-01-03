package controllers

import (
	"github.com/revel/revel"
	"regexp"
	"revel_cms/models"
)

func (c App) Posts(id string) revel.Result {

	if t, _ := regexp.MatchString("^[a-f0-9]{24}$", id); !t {
		return c.Redirect(App.Index)
	}
	post := new(models.Post).GetPost(id)
	if user := c.connected(); user != nil {
		return c.Render(post, user)
	}
	return c.Render(post)
}

func (c App) FormPost(title, text, tags string) revel.Result {
	if user := c.connected(); user != nil {
		return c.Render(user)
	}
	return c.Render()
}

func (c App) Tags(tag string) revel.Result {
	c.Validation.Clear()
	posts := new(models.Post).GetPosts()

	if t, _ := regexp.MatchString("^[a-zA-Z]{2,}$", tag); t {
		posts = new(models.Post).GetPostsByTag(tag)
	}

	if user := c.connected(); user != nil {
		return c.Render(posts, user)
	}
	return c.Render(posts)
}
