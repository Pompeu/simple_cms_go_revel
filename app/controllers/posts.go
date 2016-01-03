package controllers

import (
	"github.com/revel/revel"
	"regexp"
	"revel_cms/app/models"
	"strings"
)

func (c App) Posts(id string) revel.Result {

	if key, _ := regexp.MatchString("^[a-f0-9]{24}$", id); !key {
		return c.Redirect(App.Index)
	}
	post := new(models.Post).GetPost(id)
	if user := c.connected(); user != nil {
		return c.Render(post, user)
	}
	return c.Render(post)
}

func (c App) FormPost(title, text, tags string) revel.Result {
	if user := c.connected(); user == nil {
		return c.Redirect(App.Login)
	}
	if c.Request.Method == "GET" {
		return c.Render()
	} else {
		splitTags := strings.Split(tags, " ")
		if _, err := new(models.Post).Create(
			title, text, splitTags); err != nil {
			c.Flash.Error("erro ao gravar")
			return c.Render()
		} else {
			return c.Redirect(App.Index)
		}

	}
	return c.Render()
}

func (c App) EditPost(id, title, text, tags string) revel.Result {
	if c.Request.Method == "GET" {
		post := new(models.Post).GetPost(id)
		tags = strings.Join(post.Tags, " ")
		return c.Render(post, tags)
	} else {
		splitTags := strings.Split(tags, " ")
		if err := new(models.Post).Update(
			id, title, text, splitTags); err == nil {
			return c.Redirect(App.Index)
		} else {
			return c.Render()
		}
	}

	return c.Render()
}

func (c App) RemovePost(id string) revel.Result {

	if key, err := regexp.MatchString(
		`^[a-f0-9]{24}$`, id); key && err == nil {
		new(models.Post).RemovePost(id)
		return c.Redirect(App.Index)
	}
	return c.Redirect(App.Index)
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
