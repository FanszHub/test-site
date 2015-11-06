package Server

import (
	"net/http"
	"fmt"
	"github.com/FanszHub/test-site/Api"
//	"github.com/FanszHub/test-site/Models"
//	"log"
)

func StartMyApp(port int){

//	db, err := Models.NewDB("Tmp/MyDatabase")
/*
	if err != nil {
		log.Panic(err)
	}
*/
	env := &Api.Env{Db:nil}

	http.ListenAndServe(fmt.Sprintf(":%v",port), Api.Handlers(env))
}
