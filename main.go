package main

import "github.com/gin-gonic/gin"
import "github.com/rasyidkaromi/berita/middleware"
import "github.com/rasyidkaromi/berita/controller"

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(middleware.Cors())

	handler := controller.NewStreamController()
	r.POST("/berita", handler.Post)
	r.GET("/berita", handler.Get)
	r.PUT("/berita/:id", handler.Update)
	r.DELETE("/berita/:id", handler.Delete)

	r.Run(":8090")
}
