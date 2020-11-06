package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"myself/mall/handler"
	"myself/mall/util"
)

func Ping(c *gin.Context) {
	data := make(map[string]interface{})
	data["project_name"] = viper.GetString("project.name")
	data["time"] = util.StrNow()

	handler.SendResponse(c, nil, "", data)
}
