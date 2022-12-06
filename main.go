package main

import (
	r "bbs_backend/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	r.Routes(server)
	
	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})

	server.Run(":8080")
}
