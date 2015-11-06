package Api

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/FanszHub/test-site/Models"
)

type Env struct {
	Db Models.Datastore
}

func Handlers(env *Env) *mux.Router{
	r := mux.NewRouter()

	r.Handle("/users", UserHandler(env)).Methods("GET")
	r.Handle("/", HomeHandler()).Methods("GET")

	http.Handle("/", r)

	return r
}
