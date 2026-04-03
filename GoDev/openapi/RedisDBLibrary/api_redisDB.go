package RedisDBLibrary

import (
	"lite5gc/cmn/logger"
	"lite5gc/cmn/types"
	"lite5gc/nrf/dbmgr"
	"lite5gc/oam/cm/configure"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strconv"
)

func RestfulAPIGetOne(collName string, filter string) []interface{} {
	//portStr := strconv.Itoa(configure.NrfConf.REDIS.SerAddr.Port)
	//address := fmt.Sprintf("%s%s",configure.NrfConf.REDIS.SerAddr.Ipv4,portStr)
	c, err := redis.Dial("tcp", "10.180.8.74:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
	}
	fmt.Println("Connects redis success!")

	selectRedisDB := dbmgr.SelectRedisDB(collName)
	if selectRedisDB == int8(dbmgr.RedisDBFalse) {
		logger.Trace(MODULE_ID, types.DEBUG, nil, "select Redis DB failed")
	}

	//select redis db
	_, err = c.Do("select", selectRedisDB)
	if err != nil {
		fmt.Println("select error", err)
	}
	defer c.Close()

	result, err := redis.Values(c.Do("lrange", filter, 0, -1))
	if err != nil {
		fmt.Println("Irange err", err.Error())
	}
	return result
}

func RestfulAPIPutOne(collName string, filter string, putData []byte) bool {
	//conn redis serve
	portStr := strconv.Itoa(configure.NrfConf.REDIS.SerAddr.Port)
	address := fmt.Sprintf("%s%s", configure.NrfConf.REDIS.SerAddr.Ipv4, portStr)
	c, err := redis.Dial("tcp", address)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return false
	}
	fmt.Println("Connects redis success!")

	selectRedisDB := dbmgr.SelectRedisDB(collName)
	if selectRedisDB == int8(dbmgr.RedisDBFalse) {
		logger.Trace(MODULE_ID, types.DEBUG, nil, "select Redis DB failed")
		return false
	}

	//select redis db
	_, err = c.Do("select", selectRedisDB)
	if err != nil {
		fmt.Println("select error", err)
	}

	defer c.Close()

	result, err := redis.Values(c.Do("lrange", filter, 0, -1))
	if err != nil {
		fmt.Println("Irange err", err.Error())
	}

	for _, v := range result {
		_, err := c.Do("lrem", filter, 1, v)
		if err != nil {
			fmt.Println("lrem error", err)
		}
	}

	c.Send("rpush", filter, putData)
	c.Flush()

	return true
}

func RestfulAPIGetMany(collName string, filter string) []interface{} {
	//portStr := strconv.Itoa(configure.NrfConf.REDIS.SerAddr.Port)
	//address := fmt.Sprintf("%s%s",configure.NrfConf.REDIS.SerAddr.Ipv4,portStr)
	c, err := redis.Dial("tcp", "10.180.8.74:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
	}
	fmt.Println("Connects redis success!")

	selectRedisDB := dbmgr.SelectRedisDB(collName)
	if selectRedisDB == int8(dbmgr.RedisDBFalse) {
		logger.Trace(MODULE_ID, types.DEBUG, nil, "select Redis DB failed")
	}

	//select redis db
	_, err = c.Do("select", selectRedisDB)
	if err != nil {
		fmt.Println("select error", err)
	}
	defer c.Close()

	result, err := redis.Values(c.Do("lrange", filter, 0, -1))
	if err != nil {
		fmt.Println("Irange err", err.Error())
	}

	return result
}

func RestfulAPIDeleteMany(collName string, filter string) bool {
	//portStr := strconv.Itoa(configure.NrfConf.REDIS.SerAddr.Port)
	//address := fmt.Sprintf("%s%s",configure.NrfConf.REDIS.SerAddr.Ipv4,portStr)
	c, err := redis.Dial("tcp", "10.180.8.74:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
	}
	fmt.Println("Connects redis success!")

	selectRedisDB := dbmgr.SelectRedisDB(collName)
	if selectRedisDB == int8(dbmgr.RedisDBFalse) {
		logger.Trace(MODULE_ID, types.DEBUG, nil, "select Redis DB failed")
	}

	//select redis db
	_, err = c.Do("select", selectRedisDB)
	if err != nil {
		fmt.Println("select error", err)
	}
	defer c.Close()

	_, err = c.Do("del", filter)
	if err != nil {
		fmt.Println("Irange err", err.Error())
		return false
	}

	return true
}

func RestfulAPIJSONPatch(collName string, filter string, patchJson []byte) bool {
	portStr := strconv.Itoa(configure.NrfConf.REDIS.SerAddr.Port)
	address := fmt.Sprintf("%s%s", configure.NrfConf.REDIS.SerAddr.Ipv4, portStr)
	c, err := redis.Dial("tcp", address)
	if err != nil {
		fmt.Println("Connect to redis error", err)
	}
	fmt.Println("Connects redis success!")

	selectRedisDB := dbmgr.SelectRedisDB(collName)
	if selectRedisDB == int8(dbmgr.RedisDBFalse) {
		logger.Trace(MODULE_ID, types.DEBUG, nil, "select Redis DB failed")
	}

	//select redis db
	_, err = c.Do("select", selectRedisDB)
	if err != nil {
		fmt.Println("select error", err)
	}
	defer c.Close()

	return true
}

func RestfulAPIPost(collName string, filter string, putData []byte) bool {
	//conn redis serve
	//portStr := strconv.Itoa(configure.NrfConf.REDIS.SerAddr.Port)
	//address := fmt.Sprintf("%s%s",configure.NrfConf.REDIS.SerAddr.Ipv4,portStr)
	c, err := redis.Dial("tcp", "10.180.8.74:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return false
	}
	fmt.Println("Connects redis success!")

	selectRedisDB := dbmgr.SelectRedisDB(collName)
	if selectRedisDB == int8(dbmgr.RedisDBFalse) {
		logger.Trace(MODULE_ID, types.DEBUG, nil, "select Redis DB failed")
		return false
	}

	//select redis db
	_, err = c.Do("select", selectRedisDB)
	if err != nil {
		fmt.Println("select error", err)
	}

	defer c.Close()

	c.Send("rpush", filter, putData)
	c.Flush()

	return true
}
