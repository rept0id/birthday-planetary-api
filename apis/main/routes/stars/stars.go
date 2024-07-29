package stars

import (
	"github.com/gin-gonic/gin"
	t_stars "github.com/rept0id/birthday-planetary-backend/apis/main/dbs/stars/tables/stars"
	"github.com/rept0id/birthday-planetary-backend/apis/main/model"
)

func Routes(dbs model.DBs, ginEngine *gin.Engine) {
	ginEngine.GET("/stars/find_by_iso_date", func(c *gin.Context) {
		isoDate := c.Query("iso_date")

		res, err := t_stars.FindByISODate(dbs, isoDate)
		if err != nil {
			c.Status(500)
		}

		c.JSON(200, res)
	})
}
