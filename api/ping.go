package api

import (
	"github.com/beijibeijing/viper"
	"github.com/gin-gonic/gin"
	"myself/mall/conf"
	"time"
)

func Ping(c *gin.Context) {
	data := make(map[string]interface{})
	data["project_name"] = viper.GetString("project.name")
	data["time"] = time.Now()

	conf.SendResponse(c, nil, "", data, "chs")
}
