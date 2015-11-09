package Api

import (
	"github.com/FanszHub/test-site/Models"
	"github.com/gin-gonic/gin"
)

type Env struct {
	Db Models.Datastore
}

func Routers(env *Env) *gin.Engine{

	r := gin.Default()

	r.GET("/users", UserIndex(env))
	r.POST("/users", CreateUser(env))
	r.GET("/", HomeHandler())

	return r
}