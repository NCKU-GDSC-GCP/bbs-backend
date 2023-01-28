package v1

import (
	. "bbs_backend/models"
	service "bbs_backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserRouter a contract what this router can do
type UserRouter interface {
	Setup(rg *gin.RouterGroup)
}

type userRouter struct {
	model UserModel
}

func NewUserRouter(m UserModel) UserRouter {
	return &userRouter{
		model: m,
	}
}

func (r *userRouter) Setup(rg *gin.RouterGroup) {
	user := rg.Group("v1/user")
	user.GET("/", r.GetUsers)
	user.POST("/", r.CreateUser)
}

func (r *userRouter) GetUsers(c *gin.Context) {
	users, _ := r.model.GetAll()
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
