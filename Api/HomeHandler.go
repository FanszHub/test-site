package Api

import (
	"net/http"
)

func HomeHandler() http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		w.Write([]byte("ok"))

		return
	})

}
