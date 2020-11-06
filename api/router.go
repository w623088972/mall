package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	goodsApi "myself/mall/api/goods"
	"myself/mall/middleware"
)

func InitRouter(router *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Content-Type", "Content-Length", "X-Requested-With", "x-requested-with",
		"Origin", "origin", "token", "channel", "register_source", "language", "project_id"}
	config.AllowAllOrigins = true

	router.Use(cors.New(config))
	group := router.Group("mall", middleware.PrepareContext())
	group.GET("ping", Ping)

	goodsApi.AddRoute(group)

}
