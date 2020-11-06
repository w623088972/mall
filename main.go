package main

import (
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	k8s "github.com/micro/examples/kubernetes/go/web"
	"github.com/micro/go-micro/web"
	"github.com/spf13/viper"
	"myself/mall/api"
	"myself/mall/model"
	"myself/mall/redis"
)

func main() {
	Init()

	var service web.Service
	if viper.GetString("project.env") == "local" {
		service = web.NewService(web.Name("go-code"), web.Address(viper.GetString("project.port")))
	} else {
		if os.Getenv("isy_deploy_env") == "dev" {
			service = k8s.NewService(web.Name("mall-dev"))
		} else {
			service = k8s.NewService(web.Name("mall"))
		}
	}

	//创建记录日志的文件
	logName := "storage/log/" + time.Now().Format("20060102150405") + ".log"
	f, _ := os.Create(logName)
	gin.DefaultWriter = io.MultiWriter(f)

	//路由
	router := gin.Default()
	api.InitRouter(router)
	service.Handle("/", router)

	_ = service.Init()
	_ = service.Run()
}

func Init() {
	filePath := os.Getenv("filepath")
	if filePath == "" {
		filePath = "./config"
	}
	//配置文件
	viper.AddConfigPath(filePath)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	viper.WatchConfig()

	model.DB.Init()

	//mysql
	mysqlHost := viper.GetString("mysql.host")
	mysqlPort := viper.GetString("mysql.port")
	mysqlDatabase := viper.GetString("mysql.database")
	mysqlUsername := viper.GetString("mysql.username")
	mysqlPassword := viper.GetString("mysql.password")
	mysqlLog := viper.GetBool("mysql.log")
	model.Init(mysqlHost, mysqlPort, mysqlDatabase, mysqlUsername, mysqlPassword, mysqlLog)

	//redis
	redisHost := viper.GetString("redis.host")
	redisPort := viper.GetString("redis.port")
	redisDatabase := viper.GetString("redis.database")
	redisPassword := viper.GetString("redis.password")
	redis.Init(redisHost, redisPort, redisDatabase, redisPassword)
}
