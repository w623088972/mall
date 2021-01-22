package redis

import (
	"github.com/gomodule/redigo/redis"
)

//HSet 设置散列缓存值
func HSet(redisKey string, hashKey string, hashValue string) error {
	// 从池里获取连接
	rc := RedisClient.Self.Get()

	// 用完后将连接放回连接池
	defer func() {
		rc.Close()
	}()

	_, err := rc.Do("HSET", redisKey, hashKey, hashValue)
	if err != nil {
		return err
	}
	return nil
}

//HGet 取散列缓存值
func HGet(redisKey string, hashKey string) (string, error) {
	// 从池里获取连接
	rc := RedisClient.Self.Get()

	// 用完后将连接放回连接池
	defer func() {
		rc.Close()
	}()

	hashValue, err := redis.String(rc.Do("HGET", redisKey, hashKey))
	if err != nil {
		return "", err
	}

	return hashValue, nil
}

//HDel 删除散列缓存值
func HDel(redisKey string, hashKey string) error {
	// 从池里获取连接
	rc := RedisClient.Self.Get()

	// 用完后将连接放回连接池
	defer func() {
		rc.Close()
	}()

	_, err := rc.Do("HDEL", redisKey, hashKey)
	if err != nil {
		return err
	}

	return nil
}
