package Models

import (
	"github.com/HouzuoGuo/tiedot/db"
	"log"
)

type Datastore interface {
	AllUsers() ([]*User, error)
}

type DB struct {
	*db.DB
}

func NewDB(path string) (*DB, error){

	log.Println("Opening db")

	log.Println(path)

	db, err := db.OpenDB(path)

	log.Println("Creating users")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println("Creating users")

	//err = db.Create("Users")

	if err != nil {
		log.Println(err)
	}

	log.Println("Created users")

	return &DB{db}, nil
}