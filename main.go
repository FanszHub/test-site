package main

import (
	"log"
	"github.com/FanszHub/test-site/Server"
)

func main(){

	log.Println("Starting server")

	Server.StartMyApp(3131,"TESTGoNuts")

	log.Println("Started?")
}
