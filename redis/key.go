package redis

import (
	"github.com/gomodule/redigo/redis"
)

//SetEXPIRE 设置缓存值
func SetExpire(key string, time int) error {
	// 从池里获取连接
	rc := RedisClient.Self.Get()

	// 用完后将连接放回连接池
	defer func() {
		rc.Close()
	}()

	if time != 0 {
		_, err := rc.Do("EXPIRE", key, time)
		if err != nil {
			return err
		}
	}

	return nil
}

//Exists 判断缓存是否存在
func Exists(key string) bool {
	// 从池里获取连接
	rc := RedisClient.Self.Get()

	// 用完后将连接放回连接池
	defer func() {
		rc.Close()
	}()

	exists, err := redis.Bool(rc.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

//Delete 删除缓存值
func Delete(key string) (bool, error) {
	// 从池里获取连接
	rc := RedisClient.Self.Get()

	// 用完后将连接放回连接池
	defer func() {
		rc.Close()
	}()

	return redis.Bool(rc.Do("DEL", key))
}

//LikeDeletes 批量删除缓存值
func LikeDeletes(key string) error {
	// 从池里获取连接
	rc := RedisClient.Self.Get()

	// 用完后将连接放回连接池
	defer func() {
		rc.Close()
	}()

	keys, err := redis.Strings(rc.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}
