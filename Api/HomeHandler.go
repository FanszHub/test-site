package Api

import (
	"github.com/gin-gonic/gin"
)

func HomeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "ok")
	}
}