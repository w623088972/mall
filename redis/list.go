package redis

import (
	"github.com/gomodule/redigo/redis"
)

//LPop 从列表左边移除一个数据
func LPop(redisKey string) string {
	// 从池里获取连接
	rc := RedisClient.Self.Get()

	// 用完后将连接放回连接池
	defer func() {
		rc.Close()
	}()

	value, err := redis.String(rc.Do("LPOP", redisKey))
	if err != nil {
		return ""
	}

	return value
}

//LPush 从列表右边插入数据
func LPush(redisKey, value string) error {
	// 从池里获取连接
	rc := RedisClient.Self.Get()

	// 用完后将连接放回连接池
	defer func() {
		rc.Close()
	}()

	_, err := rc.Do("LPUSH", redisKey, value)
	if err != nil {
		return err
	}

	return nil
}

//RPush 从列表右边插入数据
func RPush(redisKey, value string) error {
	// 从池里获取连接
	rc := RedisClient.Self.Get()

	// 用完后将连接放回连接池
	defer func() {
		rc.Close()
	}()

	_, err := rc.Do("RPUSH", redisKey, value)
	if err != nil {
		return err
	}

	return nil
}

//LRange 获取列表指定范围内的元素
func LRange(redisKey string, start, stop int) ([]interface{}, error) {
	// 从池里获取连接
	rc := RedisClient.Self.Get()

	// 用完后将连接放回连接池
	defer func() {
		rc.Close()
	}()

	res, err := redis.Values(rc.Do("LRANGE", redisKey, start, stop))
	if err != nil {
		return nil, err
	}

	return res, nil
}

//LLen 获取列表长度
func LLen(redisKey string) (int, error) {
	// 从池里获取连接
	rc := RedisClient.Self.Get()

	// 用完后将连接放回连接池
	defer func() {
		rc.Close()
	}()

	res, err := redis.Int(rc.Do("LLEN", redisKey))
	if err != nil {
		return 0, err
	}

	return res, nil
}

//LRem 移除列表元素
//count > 0 : 从表头开始向表尾搜索，移除与 VALUE 相等的元素，数量为 COUNT 。
//count < 0 : 从表尾开始向表头搜索，移除与 VALUE 相等的元素，数量为 COUNT 的绝对值。
//count = 0 : 移除表中所有与 VALUE 相等的值。
func LRem(redisKey, value string) error {
	// 从池里获取连接
	rc := RedisClient.Self.Get()

	// 用完后将连接放回连接池
	defer func() {
		rc.Close()
	}()

	_, err := rc.Do("LREM", redisKey, 0, value)
	if err != nil {
		return err
	}

	return nil
}
