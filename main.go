package main

import (
	"bbs_backend/loaders"
	r "bbs_backend/routers"

	_ "bbs_backend/docs"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			bbs-backend
//	@version		1.0
//	@description	A Bulletin Board System backend using golang
//	@schemes		http,https
//	@host			localhost:8080
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host			127.0.0.1:8080
//	@BasePath		/api/v1

func main() {
	loaders.Init()

	server := gin.Default()
	r.Routes(server)

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	server.Run(":8080")
}
