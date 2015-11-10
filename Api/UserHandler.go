package Api

import (
	"github.com/FanszHub/test-site/Models"
	"github.com/gin-gonic/gin"
)

func UserIndex(env *Env) gin.HandlerFunc {
	return func(c *gin.Context) {

		var users []*Models.User
		users, err := env.Db.AllUsers()

		if err != nil {
			c.JSON(400, nil)
		}

		c.JSON(200, users)
	}
}

func CreateUser(env *Env) gin.HandlerFunc {
	return func(c *gin.Context) {

		var user *Models.User

		if c.BindJSON(&user) == nil {
			env.Db.AddUser(user)
			c.JSON(200, gin.H{})
		} else {
			c.JSON(400, gin.H{})
		}
	}
}