package redis

import (
	"myself/mall/conf"
	"time"

	"github.com/beijibeijing/viper"
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
)

//RedisDb 结构体
type RedisDb struct {
	Self *redis.Pool
}

//RedisClient 全局引用
var RedisClient *RedisDb

//Init 初始化 待优化
func (r *RedisDb) Init() {
	RedisClient = &RedisDb{
		Self: openRedisDb(),
	}
}

func openRedisDb() *redis.Pool {
	rs := &redis.Pool{
		MaxIdle:     viper.GetInt("redis.maxIdle"),
		MaxActive:   viper.GetInt("redis.maxActive"),
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			host := viper.GetString("redis.host")
			port := viper.GetString("redis.port")
			password := viper.GetString("redis.password")
			conn, err := redis.Dial("tcp", host+port, redis.DialPassword(password))
			if err != nil {
				conf.LOG.Self.WithFields(logrus.Fields{
					"host":     host,
					"port":     port,
					"password": password,
					"err":      err.Error,
				}).Info("openRedisDb Dial err")
				return nil, err
			}

			//选择db
			_, err = conn.Do("SELECT", viper.GetInt("redis.database"))
			if err != nil {
				conf.LOG.Self.WithFields(logrus.Fields{
					"host":     host,
					"port":     port,
					"password": password,
					"err":      err.Error,
				}).Info("openRedisDb Do err")
				return nil, err
			}

			return conn, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				conf.LOG.Self.WithFields(logrus.Fields{
					"err": err.Error,
				}).Info("openRedisDb TestOnBorrow PING err")
			}
			return err
		},
	}

	conf.LOG.Self.Info("openRedisDb openRedis done")
	return rs
}
