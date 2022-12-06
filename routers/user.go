package router

import (
	"bbs_backend/entity"
	"net/http"
	"github.com/gin-gonic/gin"
)

var users = []entity.User{
	{Id: 1, Name: "Lewis", Email: "robertlewis@nycu.com"},
	{Id: 2, Name: "Steven", Email: "steven@trendm.com"},
	{Id: 3, Name: "Bob", Email: "bobbob@gamil.com"},
}

// UserRouter a contract what this router can do
type UserRouter interface {
	GetUsers(c *gin.Context)
}

type userRouter struct {
	// services
}

func NewUserRouter() UserRouter {
	return &userRouter{}
}

func (r *userRouter) GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}
