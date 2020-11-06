package goods

import (
	"github.com/gin-gonic/gin"
)

func AddRoute(group *gin.RouterGroup) {
	//auth := middleware.AuthMiddleware("")
	api := group.Group("")
	{
		api.GET("goods/info", GoodsInfo) //商品详情
	}
}
