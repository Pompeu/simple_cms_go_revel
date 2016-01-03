package models

import (
	"github.com/pompeu/db"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Id       bson.ObjectId `json:"id" bson:"_id`
	Name     string        `json:"name" bson:"name"`
	Email    string        `json:"email" bson:"email"`
	Password string        `json:"password" bson:"password"`
}

func (p *Person) Create(name, email, hash string) error {
	p.Id = bson.NewObjectId()
	p.Name = name
	p.Email = email
	p.Password = hash
	session := db.SimpleSession("persons")
	defer session.Close()
	err := session.DB("test").C("persons").Insert(p)
	return err
}

func (p *Person) Login(email string) (Person, error) {
	session := db.SimpleSession("persons")
	err := session.DB("test").C("persons").Find(bson.M{"email": email}).One(&p)
	defer session.Close()
	return *p, err
}
