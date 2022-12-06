package router

import "github.com/gin-gonic/gin"

var (
	user UserRouter = NewUserRouter()
)

func Routes(app *gin.Engine) {
	userRoutes := app.Group("api/user")
	{
		userRoutes.GET("/get", user.GetUsers)
	}
}