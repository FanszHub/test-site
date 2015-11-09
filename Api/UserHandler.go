package Api

import (
"net/http"
"github.com/FanszHub/test-site/Models"
	"io/ioutil"
	"io"
	"encoding/json"
)

func UserIndex(env *Env) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		var users []*Models.User
		users, err := env.Db.AllUsers()

		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(users); err != nil {
			panic(err)
		}
	})
}

func CreateUser(env *Env) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		var user *Models.User

		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

		if err != nil {
			panic(err)
		}

		json.Unmarshal(body,&user)

		env.Db.AddUser(user)

		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)

	})
}