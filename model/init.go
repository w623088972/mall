package model

import (
	"fmt"
	"myself/mall/conf"

	"github.com/beijibeijing/viper"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

//Database 数据库
type Database struct {
	Self *gorm.DB
	//Docker *gorm.DB
}

//DB 数据库
var DB *Database

/*
func InitDockerDb() *gorm.DB {
	return openDb(
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.database"),
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"))
}

func GetDockerDb() *gorm.DB {
	return InitDockerDb()
}
*/

//Init 数据库初始化
func (db *Database) Init() {
	DB = &Database{
		Self: GetSelfDb(),
		//Docker: GetDockerDb(),
	}
}

//Close 数据库关闭
func (db *Database) Close() {
	DB.Self.Close()
	//DB.Docker.Close()
}

//GetSelfDB 数据库调用
func GetSelfDb() *gorm.DB {
	return InitSelfDb()
}

//InitSelfDb 数据库初始化
// used for cli
func InitSelfDb() *gorm.DB {
	return openDb(
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.database"),
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"))
}

func openDb(host, port, database, username, password string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		username,
		password,
		host+port,
		database,
		true,
		//"Asia/Shanghai"),
		"Local")

	db, err := gorm.Open("mysql", config)
	if err != nil {
		//log.Errorf(err, "Database connection failed. Database name: %s", name)
		conf.LOG.Self.WithFields(logrus.Fields{
			"username": username,
			"password": password,
			"host":     host,
			"port":     host,
			"name":     database,
			"err":      err.Error,
		}).Info("openDb err")
	} else {
		conf.LOG.Self.Info("Database openDb done")
	}

	// set for db connection
	setupDB(db)

	return db
}

func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("mysql.log")) //true打印详细日志

	//db.SetLogger(gorm.Logger{gin.DefaultWriter})
	db.SetLogger(gorm.Logger{conf.LOG.Self})
	db.DB().SetMaxOpenConns(viper.GetInt("mysql.maxOpenConns")) // 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	db.DB().SetMaxIdleConns(viper.GetInt("mysql.maxIdleConns")) // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
}
