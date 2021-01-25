package conf

import (
	"io"
	"os"
	"syscall"
	"time"

	"github.com/beijibeijing/viper"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

//LogLogger 日志
type LogLogger struct {
	Self *logrus.Logger
}

//LOG 日志
var LOG *LogLogger

//Init 日志初始化
func (log *LogLogger) Init() {
	LOG = &LogLogger{
		Self: GetSelfLog(),
	}
}

//GetSelfLog 日志初始化
func GetSelfLog() *logrus.Logger {
	//fileTruePath := viper.GetString("log.truePath")
	filePath := viper.GetString("log.path")
	fileName := viper.GetString("log.name")

	var log = logrus.New()

	//设置日志格式
	if viper.GetBool("log.json") {
		log.SetFormatter(&logrus.JSONFormatter{ // 为当前logrus实例设置消息输出格式为json格式。
			TimestampFormat: "2006-01-02 15:04:05.000 ",
		})
	} else {
		log.SetFormatter(&logrus.TextFormatter{
			//ForceColors:            viper.GetBool("log.color"),
			ForceColors:            true,
			DisableLevelTruncation: true,
			TimestampFormat:        "2006-01-02 15:04:05.000",
			DisableColors:          false,
			FullTimestamp:          true,
		})
	}
	//log.SetFormatter(log.Formatter)

	//设置日志等级
	log.SetLevel(logrus.TraceLevel)
	//设置日志输出
	//log.Out = os.Stdout //log.Out = logFile
	//log.SetOutput(colorable.NewColorableStdout())

	logFile, err := os.OpenFile(filePath+fileName, syscall.O_RDWR|syscall.O_CREAT|syscall.O_APPEND, 0666)
	if err == nil {
		gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
		log.SetOutput(gin.DefaultWriter) //日志输入
	} else {
		log.Info("Failed to log to file, using default stderr")
		log.Info("os.OpenFile err is:" + err.Error())
	}

	// 设置 rotatelogs
	var logWriter *rotatelogs.RotateLogs
	//if viper.GetBool("localTest") { //本地没有软连接
	logWriter, _ = rotatelogs.New(
		// 分割后的文件名称
		filePath+"%Y%m%d%H.log",

		// 生成软链，指向最新日志文件
		//rotatelogs.WithLinkName(fileName+"now.log"),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1小时)
		rotatelogs.WithRotationTime(time.Hour),
		// WithMaxAge和WithRotationCount二者只能设置一个，
		// WithMaxAge设置文件清理前的最长保存时间，
		// WithRotationCount设置文件清理前最多保存的个数。
		//rotatelogs.WithMaxAge(time.Hour*24),
		//rotatelogs.WithRotationCount(24*7),
	)
	/*
		} else {
			logWriter, err = rotatelogs.New(
				filePath+"%Y%m%d%H.log",
				rotatelogs.WithLinkName(linkName),
				// 设置最大保存时间(7天)
				rotatelogs.WithMaxAge(7*24*time.Hour),
				// 设置日志切割时间间隔(1小时)
				rotatelogs.WithRotationTime(time.Hour),
			)

			if err != nil {
				log.Info("rotatelogs.New err is:" + err.Error())
			}
		}
	*/

	writeMap := lfshook.WriterMap{
		logrus.TraceLevel: logWriter,
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		//lfHook := lfshook.NewHook(gin.DefaultWriter, &logrus.TextFormatter{
		//ForceColors:            viper.GetBool("logColor"),
		//DisableLevelTruncation: true,
		TimestampFormat: "2006-01-02 15:04:05.000",
		//DisableColors:          true,
	})

	// 新增 Hook
	log.AddHook(lfHook)

	log.Info("LOG Init done")

	return log
}
