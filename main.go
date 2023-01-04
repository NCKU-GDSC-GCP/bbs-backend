package main

import (
	"bbs_backend/loaders"
	r "bbs_backend/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	loaders.Init()

	server := gin.Default()
	r.Routes(server)

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})

	server.Run(":8080")
}
