package redis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"log"
)

func Exists(key string) bool {
	//从池里获取连接
	rc := RedisClient.Self.Get()

	//用完后将连接放回连接池
	defer func() {
		_ = rc.Close()
		log.Println("redis Exists end redis ActiveCount:", RedisClient.Self.ActiveCount())
	}()
	log.Println("redis Exists start redis ActiveCount:", RedisClient.Self.ActiveCount())

	exists, err := redis.Bool(rc.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

//Set 会将data转为json, 所以data须是结构或其他复合类型
func Set(key string, data interface{}, time int) error {
	//从池里获取连接
	rc := RedisClient.Self.Get()

	//用完后将连接放回连接池
	defer func() {
		_ = rc.Close()
		log.Println("redis Set end redis ActiveCount:", RedisClient.Self.ActiveCount())
	}()
	log.Println("redis Set start redis ActiveCount:", RedisClient.Self.ActiveCount())

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

func Get(key string) ([]byte, error) {
	//从池里获取连接
	rc := RedisClient.Self.Get()

	//用完后将连接放回连接池
	defer func() {
		_ = rc.Close()
		log.Println("redis Get end redis ActiveCount:", RedisClient.Self.ActiveCount())
	}()
	log.Println("redis Get start redis ActiveCount:", RedisClient.Self.ActiveCount())

	reply, err := redis.Bytes(rc.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func Delete(key string) (bool, error) {
	//从池里获取连接
	rc := RedisClient.Self.Get()

	//用完后将连接放回连接池
	defer func() {
		_ = rc.Close()
		log.Println("redis Delete end redis ActiveCount:", RedisClient.Self.ActiveCount())
	}()
	log.Println("redis Delete start redis ActiveCount:", RedisClient.Self.ActiveCount())

	return redis.Bool(rc.Do("DEL", key))
}

func LikeDeletes(key string) error {
	//从池里获取连接
	rc := RedisClient.Self.Get()

	//用完后将连接放回连接池
	defer func() {
		_ = rc.Close()
		log.Println("redis LikeDeletes end redis ActiveCount:", RedisClient.Self.ActiveCount())
	}()
	log.Println("redis LikeDeletes start redis ActiveCount:", RedisClient.Self.ActiveCount())

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

func Incr(key string) (int, error) {
	//从池里获取连接
	rc := RedisClient.Self.Get()

	//用完后将连接放回连接池
	defer func() {
		_ = rc.Close()
		log.Println("redis Incr end redis ActiveCount:", RedisClient.Self.ActiveCount())
	}()
	log.Println("redis Incr start redis ActiveCount:", RedisClient.Self.ActiveCount())

	reply, err := redis.Int(rc.Do("Incr", key))
	if err != nil {
		return 0, err
	}

	return reply, nil
}
