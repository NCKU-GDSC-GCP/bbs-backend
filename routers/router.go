package router

import (
	v1 "bbs_backend/routers/api/v1"

	"github.com/gin-gonic/gin"
)

var (
	userRoute v1.UserRouter = v1.NewUserRouter()
)

func Routes(app *gin.Engine) {
	apiRoute := app.Group("api")
	userRoute.Setup(apiRoute)
}
