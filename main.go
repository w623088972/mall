package main

import (
	"github.com/beijibeijing/viper"
	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	"myself/mall/api"
	"myself/mall/conf"
	"myself/mall/middleware"
	"myself/mall/model"
	"myself/mall/redis"
)

//初始化相关
func init() {
	//配置文件初始化
	if err := conf.ConfigInit(""); err != nil {
		panic(err)
	}

	conf.LOG.Init()          //日志初始化
	conf.CronC.Init()        //定时任务初始化
	model.DB.Init()          //mysql初始化 	//defer model.DB.Close()
	redis.RedisClient.Init() //redis初始化
	//conf.PrometheusInit()    //监控初始化

	//conf.LOG.Self.Info("Main init done")
	//conf.CronC.Self.Start()
}

func main() {
	defer conf.CatchPanic("main") //捕获异常
	gin.SetMode(viper.GetString("runmode"))

	//r := gin.Default()
	router := gin.New()
	router.Use(middleware.Logging())
	router.Use(ginprom.PromMiddleware(nil)) //prometheus for gin

	api.InitRouter(router)

	router.Run(viper.GetString("project.port"))
}
