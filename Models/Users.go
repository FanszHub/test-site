package Models
import "encoding/json"

type User struct {
	Username string
}

func (db *DB) AllUsers() ([]*User, error) {

	rows := db.Use("Users")

	users := make([]*User, 0)

	rows.ForEachDoc(func(id int, docContent []byte)(willMoveOn bool){
		var u *User
		_ = json.Unmarshal(docContent, u)
		users = append(users, u)
		return true
	})

	return users, nil
}