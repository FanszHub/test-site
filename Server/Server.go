package Server

import (
	"fmt"
	"log"
	"github.com/facebookgo/inject"
	"github.com/FanszHub/test-site/Data"
	"github.com/FanszHub/test-site/Env"
	"github.com/FanszHub/test-site/Api")

func StartMyApp(port int, dbName string) {

	var env Env.Env
	var g inject.Graph

	db, e := Data.NewDB(dbName)

	log.Println(e)

	err := g.Provide(
		&inject.Object{Value: &env},
		&inject.Object{Value: db},
	)

	if err != nil {
		log.Fatalf("Error providing dependencies: ", err.Error())
	}

	if err := g.Populate(); err != nil {
		log.Fatalf("Error providing dependencies: ", err.Error())
	}

	routes := Api.Routes(env)

	routes.Run(fmt.Sprintf(":%v", port))
}
