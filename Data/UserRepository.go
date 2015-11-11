package Data
import (
	"github.com/FanszHub/test-site/Models"
	"log"
)

type UserRepository interface {
	AllUsers() ([]*Models.User, error)
	AddUser(*Models.User) (error)
}

func (db *DB) AllUsers() ([]*Models.User, error) {

	var users []*Models.User

	log.Println(db)

	err := db.C("Users").Find(nil).All(&users)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (db *DB) AddUser(user *Models.User) (error) {
	return db.C("Users").Insert(user)
}