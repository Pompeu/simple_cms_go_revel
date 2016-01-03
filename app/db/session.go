package db

import (
	"gopkg.in/mgo.v2"
)

func SimpleSession(col string) *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost/" + col)
	if err != nil {
		panic(err)
	}
	return session
}
