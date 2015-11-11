package Api

import (
	"github.com/FanszHub/test-site/Models"
	"github.com/gin-gonic/gin"
	"github.com/FanszHub/test-site/Env"
	"log"
)

func UserIndex(env *Env.Env) (gin.HandlerFunc) {

	return func(c *gin.Context) {
		var userRepository = env.UserRepository;

		var users []*Models.User

		users, err := userRepository.AllUsers()

		if err != nil {
			c.JSON(400, nil)
		}

		c.JSON(200, users)
	}
}

func CreateUser(env *Env.Env) (gin.HandlerFunc) {

	return func(c *gin.Context) {

		var userRepository = env.UserRepository;
		var user Models.User

		err := c.Bind(&user)

		if err == nil {
			userRepository.AddUser(&user)
			c.JSON(200, gin.H{})
		} else {
			log.Println(err)
			c.JSON(400, gin.H{})
		}
	}
}
