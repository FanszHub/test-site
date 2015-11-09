package Api

import (
	"github.com/gorilla/mux"
	"github.com/FanszHub/test-site/Models"
)

type Env struct {
	Db Models.Datastore
}

func Handlers(env *Env) *mux.Router{
	r := mux.NewRouter()

	r.Handle("/users", UserIndex(env)).Methods("GET")
	r.Handle("/users", CreateUser(env)).Methods("POST")
	r.Handle("/", HomeHandler()).Methods("GET")

	return r
}