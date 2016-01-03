package models

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
	"revel_cms/app/db"
)

type Post struct {
	Id    bson.ObjectId `json:"id" bson:"_id`
	Title string        `json:"title" bson:"title"`
	Body  string        `json:"body" bson:"body"`
	Tags  []string      `json:"tags" bson:"tags"`
}

func (p *Post) Create(title, body string, tags []string) (Post, error) {
	p.Id = bson.NewObjectId()
	p.Title = title
	p.Body = body
	p.Tags = tags
	session := db.SimpleSession("posts")
	err := session.DB("test").C("posts").Insert(p)
	defer session.Close()
	return *p, err
}

func (p *Post) Update(id, title, body string, tags []string) error {
	oid := bson.ObjectIdHex(id)
	p.Id = oid
	p.Title = title
	p.Body = body
	p.Tags = tags
	session := db.SimpleSession("posts")
	err := session.DB("test").C("posts").Update(bson.M{"id": oid}, &p)
	defer session.Close()
	return err
}
func (p *Post) FindByName(title string) []Post {
	var posts []Post
	session := db.SimpleSession("posts")
	if err := session.DB("test").C("posts").Find(
		bson.M{"title": &bson.RegEx{
			Pattern: title, Options: "i"}}).All(&posts); err != nil {
		panic(err)
	}
	defer session.Close()
	return posts
}

func (p *Post) GetPostsByTag(tag string) []Post {
	var posts []Post
	session := db.SimpleSession("posts")
	if err := session.DB("test").C("posts").Find(bson.M{"tags": &bson.RegEx{Pattern: tag, Options: "i"}}).All(&posts); err != nil {
		panic(err)
	}
	defer session.Close()
	return posts
}

func (p *Post) GetPosts() []Post {
	var posts []Post
	session := db.SimpleSession("posts")
	if err := session.DB("test").C("posts").Find(bson.M{}).All(&posts); err != nil {
		panic(err)
	}
	defer session.Close()
	return posts
}

func (p *Post) GetPost(id string) Post {
	session := db.SimpleSession("posts")
	oid := bson.ObjectIdHex(id)
	if err := session.DB("test").C("posts").Find(bson.M{"id": oid}).One(&p); err != nil {
		panic(err)
	}
	defer session.Close()
	return *p
}

func (p *Post) RemovePost(id string) (done bool, err error) {
	session := db.SimpleSession("posts")
	if !bson.IsObjectIdHex(id) {
		err = errors.New("id invaid")
		return done, err
	}
	oid := bson.ObjectIdHex(id)
	err = session.DB("test").C("posts").Remove(bson.M{"id": oid})
	defer session.Close()
	if err == nil {
		done = true
	}
	return done, err
}
