package Api

import (
	"github.com/FanszHub/test-site/Env"
	"github.com/gin-gonic/gin"
)

func Routes(env Env.Env) *gin.Engine{
	routes := gin.Default()

	routes.GET("/users", UserIndex(&env))
	routes.POST("/users", CreateUser(&env))

	routes.GET("/blahs", BlahIndex(&env))
	routes.POST("/blahs", CreateBlah(&env))

	routes.GET("/", HomeHandler)

	return routes
}
