package Api

import (
	"net/http"
	"log"
)

func HomeHandler() http.Handler{

	log.Print("hi")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		log.Print("hi2")

		w.Write([]byte("ok"))

		return
	})
}
