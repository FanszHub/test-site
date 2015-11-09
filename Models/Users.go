package Models

type User struct {
	Username string
}

func (db *DB) AllUsers() ([]*User, error) {

	var users []*User

	err := db.C("Users").Find(nil).All(&users)

	if err != nil {
		panic(err)
	}

	return users, nil
}

func (db *DB) AddUser(user *User) (error) {

	err := db.C("Users").Insert(user)

	return err
}