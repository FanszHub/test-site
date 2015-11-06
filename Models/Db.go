package Models

import "github.com/HouzuoGuo/tiedot/db"

type Datastore interface {
	AllUsers() ([]*User, error)
}

type DB struct {
	*db.DB
}

func NewDB(path string) (*DB, error){
	db, err := db.OpenDB(path)

	if err != nil {
		return nil, err
	}

	err = db.Create("Users")

	return &DB{db}, err
}