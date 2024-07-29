package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rept0id/birthday-planetary-backend/apis/main/model"
	stars "github.com/rept0id/birthday-planetary-backend/apis/main/routes/stars"
)

func Routes(dbs model.DBs, ginEngine *gin.Engine) {
	stars.Routes(dbs, ginEngine)
}
