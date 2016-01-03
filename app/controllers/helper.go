package controllers

import (
	"golang.org/x/crypto/bcrypt"
	"revel_cms/models"
)

func (c App) connected() *models.Person {
	if c.RenderArgs["user"] != nil {
		return c.RenderArgs["user"].(*models.Person)
	}
	if userEmail, ok := c.Session["user"]; ok {
		return c.getUser(userEmail)
	}
	return nil
}

func (c App) getUser(email string) *models.Person {
	p, err := new(models.Person).Login(email)
	if err == nil {
		return &p
	}
	return nil
}

func genereteCode(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, 12)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

func compare(hash, pass []byte) error {
	err := bcrypt.CompareHashAndPassword(hash, pass)
	return err
}
