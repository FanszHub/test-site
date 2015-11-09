package Api

import (
"net/http"
"github.com/FanszHub/test-site/Models"
	"io/ioutil"
	"io"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func UserIndex(env *Env) gin.HandlerFunc{
	return func(c *gin.Context){

		var users []*Models.User
		users, err := env.Db.AllUsers()

		if err != nil {
			panic(err)
		}

		c.Writer.Header().Set("Content-Type", "application/json;charset=UTF-8")
		c.Writer.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(c.Writer).Encode(users); err != nil {
			panic(err)
		}
	}
}

func CreateUser(env *Env) gin.HandlerFunc{
	return func(c *gin.Context){

		var user *Models.User

		body, err := ioutil.ReadAll(io.LimitReader(c.Request.Body, 1048576))

		if err != nil {
			panic(err)
		}

		json.Unmarshal(body,&user)

		env.Db.AddUser(user)

		c.Writer.Header().Set("Content-Type", "application/json;charset=UTF-8")
		c.Writer.WriteHeader(http.StatusOK)

	}
}