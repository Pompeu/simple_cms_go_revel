package controllers

import (
	"github.com/revel/revel"
	"revel_cms/app/models"
)

func (c App) Login(email, password string) revel.Result {
	person := c.getUser(email)
	if person != nil {
		err := compare([]byte(person.Password), []byte(password))
		if err == nil {
			c.Session["user"] = person.Email
			c.Session.SetDefaultExpiration()
			c.Flash.Success("Bem Vindo ," + person.Name)
			return c.Redirect(App.Index)
		}
		c.Flash.Error("Login Failed try again")
		return c.Render()
	}
	return c.Render()
}

func (c App) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(App.Index)
}

func (c App) Registrar(name, email, password string) revel.Result {
	c.Validation.Clear()
	if c.Request.Method == "POST" {
		c.Validation.Required(name).Message("Name is Required")
		c.Validation.Email(email).Message("Email is Required")
		c.Validation.Required(password).Message("Password is Required")
		if c.Validation.HasErrors() {
			c.Validation.Keep()
			c.FlashParams()
			return c.Render()
		} else {
			p := &models.Person{}
			hash := genereteCode([]byte(password))
			if err := p.Create(name, email, hash); err == nil {
				c.Session["user"] = name
				c.Flash.Success("Bem Vindo ," + name)
				return c.Redirect(App.Index)
			}

		}
	}
	return c.Render()
}
