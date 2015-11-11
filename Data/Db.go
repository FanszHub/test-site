package Data

import (
	"log"
	"gopkg.in/mgo.v2"
)

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
