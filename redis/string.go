package redis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
)

//Set 设置缓存值
func Set(key string, data interface{}, time int) error {
	// 从池里获取连接
	rc := RedisClient.Self.Get()

	// 用完后将连接放回连接池
	defer func() {
		rc.Close()
	}()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = rc.Do("SET", key, value)
	if err != nil {
		return err
	}

	if time != 0 {
		_, err = rc.Do("EXPIRE", key, time)
		if err != nil {
			return err
		}
	}

	return nil
}

//Get 获取缓存值
func Get(key string) ([]byte, error) {
	// 从池里获取连接
	rc := RedisClient.Self.Get()

	// 用完后将连接放回连接池
	defer func() {
		rc.Close()
	}()

	reply, err := redis.Bytes(rc.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

//Incr 自增
func Incr(key string) (int, error) {
	// 从池里获取连接
	rc := RedisClient.Self.Get()

	// 用完后将连接放回连接池
	defer func() {
		rc.Close()
	}()

	reply, err := redis.Int(rc.Do("INCR", key))
	if err != nil {
		return 0, err
	}

	return reply, nil
}
