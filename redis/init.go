package redis

import (
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

type RedisDb struct {
	Self *redis.Pool
}

var RedisClient *RedisDb

func (r *RedisDb) Init() {
	RedisClient = &RedisDb{
		Self: openRedisDb(),
	}
}

func openRedisDb() *redis.Pool {
	rs := &redis.Pool{
		MaxIdle:     100,
		MaxActive:   1000,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", viper.GetString("redis.address")+viper.GetString("redis.port"), redis.DialPassword(viper.GetString("redis.password")))
			if err != nil {
				log.Println("openRedisDb redis.Dial", err)
				return nil, err
			}
			//选择database
			_, err = c.Do("SELECT", viper.GetInt("redis.database"))
			if err != nil {
				log.Println("openRedisDb c.Do failed.", err)
				return nil, err
			}
			return c, nil
		},
	}

	return rs
}

func Init(address, port, database, password string) {
	rs := &redis.Pool{
		MaxIdle:     100,
		MaxActive:   1000,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", address+port, redis.DialPassword(password))
			if err != nil {
				log.Println("InitRedis redis.Dial failed.", err)
				return nil, err
			}
			//选择database
			_, err = c.Do("SELECT", database)
			if err != nil {
				log.Println("InitRedis c.Do failed.", err)
				return nil, err
			}
			return c, nil
		},
	}
	RedisClient = &RedisDb{
		Self: rs,
	}
}
