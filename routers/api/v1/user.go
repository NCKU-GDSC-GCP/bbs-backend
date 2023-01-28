package v1

import (
	model "bbs_backend/models"
	service "bbs_backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserRouter a contract what this router can do
type UserRouter interface {
	Setup(rg *gin.RouterGroup)
}

type userRouter struct {
}

func NewUserRouter() UserRouter {
	return &userRouter{}
}

func (r *userRouter) Setup(rg *gin.RouterGroup) {
	rg.GET("/", r.GetUsers)
	rg.POST("/", r.CreateUser)
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
