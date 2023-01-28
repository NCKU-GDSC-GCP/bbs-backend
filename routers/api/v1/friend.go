package v1

import "github.com/gin-gonic/gin"

type FriendRouter interface {
	Setup(rg *gin.RouterGroup)
}

type friendRouter struct {
}

func NewFriendRouter() FriendRouter {
	return &friendRouter{}
}

func (r *friendRouter) Setup(rg *gin.RouterGroup) {
	user := rg.Group("v1/friends")
	user.POST("/", r.AddFriend)
	user.GET("/:id", r.GetFriends)
	user.PUT("/:id", r.UpdateFriend)
	user.DELETE("/:id", r.DeleteFriend)
}

// CreateFriend	 godoc
// @Summary      Add Friend
// @Tags         friends
// @Accept       json
// @Produce      json
// @Success      201  {object}  nil
// @Failure      500  {object}  nil
// @Router       /friends [post]
func (r *friendRouter) AddFriend(c *gin.Context) {

}

// GetFriends 	 godoc
// @Summary      Get the user's all friends
// @Tags         friends
// @Accept       json
// @Produce      json
// @Success      200  {object}  nil
// @Failure      500  {object}  nil
// @Router       /friends/:id [get]
func (r *friendRouter) GetFriends(c *gin.Context) {

}

// UpdateFriend	 godoc
// @Summary      Update the friend
// @Tags         friends
// @Accept       json
// @Produce      json
// @Success      200  {object}  nil
// @Failure      500  {object}  nil
// @Router       /friends/:id [put]
func (r *friendRouter) UpdateFriend(c *gin.Context) {

}

// DeleteFriend	 godoc
// @Summary      Delete the friendship
// @Tags         friends
// @Accept       json
// @Produce      json
// @Success      200  {object}  nil
// @Failure      500  {object}  nil
// @Router       /friends/:id [delete]
func (r *friendRouter) DeleteFriend(c *gin.Context) {

}
