package Server

import (
	"net/http"
	"fmt"
	"github.com/FanszHub/test-site/Api"
	"github.com/FanszHub/test-site/Models"
	"log"
)

func StartMyApp(port int){

	db, err := Models.NewDB("/home/travis/gopath/src/github.com/FanszHub/test-site/Tmp/MyDatabase")

	if err != nil {
		log.Fatal(err)
	}

	env := &Api.Env{Db:db}

	http.ListenAndServe(fmt.Sprintf(":%v",port), Api.Handlers(env))
}