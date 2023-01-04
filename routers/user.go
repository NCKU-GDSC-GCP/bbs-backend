package router

import (
	"bbs_backend/models"
	"bbs_backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserRouter a contract what this router can do
type UserRouter interface {
	GetUsers(c *gin.Context)
	CreateUser(c *gin.Context)
}

type userRouter struct {
}

func NewUserRouter() UserRouter {
	return &userRouter{}
}

func (r *userRouter) GetUsers(c *gin.Context) {
	users, _ := model.UserModel.GetAll()
	c.IndentedJSON(http.StatusOK, users)
}

func (r *userRouter) CreateUser(c *gin.Context) {
	name := c.PostForm("nickname")
	email := c.PostForm("email")
	passwd := c.PostForm("password")
	// headshot := c.DefaultPostForm("headshot", "")
	user, err := service.UserService.CreateUser(name, email, passwd)
	if err != nil {
		c.JSON(http.StatusConflict,
			gin.H{"message": "Failed to create user"})
	}

	c.JSON(201, user)
}
