package Api

import (
"net/http"
"github.com/FanszHub/test-site/Models"
)

func UserHandler(env *Env) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		var users []*Models.User
		users, err := env.Db.AllUsers()

		if err != nil {
			panic(err)
		}

		w.Write([]byte(users[0].Username))

		return
	})

}
