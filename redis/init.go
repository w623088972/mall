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
			redisHost := viper.GetString("redis.host")
			redisPort := viper.GetString("redis.port")
			redisPassword := viper.GetString("redis.password")
			c, err := redis.Dial("tcp", redisHost+redisPort, redis.DialPassword(redisPassword))
			if err != nil {
				conf.LOG.Self.WithFields(logrus.Fields{
					"redisHost":     redisHost,
					"redisPassword": redisPassword,
					"err":           err.Error,
				}).Info("openRedisDb Dial err")
				return nil, err
			}

			/* 			c, err := redis.Dial("tcp", redisdbHost)
			   			if err != nil {
			   				conf.LOG.Self.WithFields(logrus.Fields{
			   					"redisdbHost": redisdbHost,
			   					"redisdbPw":   redisdbPw,
			   					"err":         err.Error,
			   				}).Info("openRedisDB Dial err")
			   				return nil, err
			   			}
			   			if redisdbPw != "" {
			   				if _, err := c.Do("AUTH", redisdbPw); err != nil {
			   					c.Close()
			   					if err != nil {
			   						conf.LOG.Self.WithFields(logrus.Fields{
			   							"redisdbHost": redisdbHost,
			   							"redisdbPw":   redisdbPw,
			   							"err":         err.Error,
			   						}).Info("openRedisDB AUTH err")
			   						return nil, err
			   					}
			   				}
			   			} */

			// 选择db
			c.Do("SELECT", viper.GetInt("redis.database"))
			return c, nil
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

	conf.LOG.Self.Info("Database openRedis done")
	return rs
}
