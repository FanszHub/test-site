package Api

import (
	"github.com/FanszHub/test-site/Env"
	"github.com/gin-gonic/gin"
	"github.com/FanszHub/test-site/Models"
)

func BlahIndex(env *Env.Env) (gin.HandlerFunc) {

	return func(c *gin.Context) {

		var blahRepository = env.BlahRepository
		var blahs []*Models.Blah

		blahs, err := blahRepository.AllBlahs()

		if err != nil {
			c.JSON(400, nil)
		}

		c.JSON(200, blahs)
	}
}

func CreateBlah(env *Env.Env) (gin.HandlerFunc) {

	return func(c *gin.Context) {

		var blahRepository = env.BlahRepository
		var blah Models.Blah

		c.BindJSON(&blah)

		blahRepository.AddBlah(&blah)

		c.JSON(200, nil)
	}

}