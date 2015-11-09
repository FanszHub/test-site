package Models

import (
	"log"
	"gopkg.in/mgo.v2"
//	"gopkg.in/mgo.v2/bson"
)

type Datastore interface {
	AllUsers() ([]*User, error)
	AddUser(*User) (error)
}

type DB struct {
	*mgo.Database
}

func NewDB(name string) (*DB, error){

	session, err := mgo.Dial("localhost")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	c := session.DB(name)

	return &DB{c}, nil
}
