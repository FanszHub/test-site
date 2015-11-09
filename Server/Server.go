package Server

import (
	"fmt"
	"github.com/FanszHub/test-site/Api"
	"github.com/FanszHub/test-site/Models"
	"log"
)

func StartMyApp(port int, dbName string){

	db, err := Models.NewDB(dbName)

	if err != nil {
		log.Fatal(err)
	}

	env := &Api.Env{Db:db}

	routes := Api.Routers(env)

	routes.Run(fmt.Sprintf(":%v",port))
}
