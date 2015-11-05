package Server

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

func StartMyApp(port int){
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")

	http.Handle("/", r)

	http.ListenAndServe(fmt.Sprintf(":%v",port),nil)
}