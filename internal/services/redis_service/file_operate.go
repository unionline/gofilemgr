/*
@Time : 2020/5/15 16:59
@Author : FB
@File : file_operate
@Software: GoLand
*/
package redis_service

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	. "gofilemgr/internal/env/redis"
	"log"
	"time"
)

func SetRedisValue(key string, value interface{}, seconds ...int) (err error) {

	var sed time.Duration
	if len(seconds) == 0 {
		sed = 3600 * 24
	}

	value_json, err := json.Marshal(value)
	if err != nil {
		return
	}

	err = Redis.Set(key, value_json, time.Second*sed).Err()
	if err != nil {
		panic(err)
	}

	return
}
func GetRedisValueForModel(key string, model ...interface{}) (interface{}, error) {

	var item interface{}

	val, err := Redis.Get(key).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist", key)
		return nil, err
	} else if err != nil {
		//panic(err)
		log.Println("GetRedisValue err", err)
		return nil, err
	} else {
		fmt.Println("key", key)

		if len(model) == 0 {
			err = json.Unmarshal([]byte(val), &item)
			if err != nil {
				return nil, nil
			}

		} else if len(model) > 1 {
			return nil, nil
		}

	}

	return item, err
}

func GetRedisValue(key string) (interface{}, error) {

	val, err := Redis.Get(key).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist", key)
	} else if err != nil {
		//panic(err)
		log.Println("GetRedisValue err", err)
	} else {
		fmt.Println("key", key)
	}

	return val, err
}
