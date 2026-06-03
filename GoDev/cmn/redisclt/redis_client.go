/** Copyright(C),2020-2022
* Author: zmj
* Date: 11/25/20 4:39 PM
* Description:
 */
package redisclt

import (
	"fmt"
	"lite5gc/cmn/types/configure"
)

// global redis client
var Agent *Cacher

func RedisCltInit() error {
	var err error
	// create the redis clt for communication with amf main process
	Agent, err = New(
		Options{
			Addr:     fmt.Sprintf("%s:%d", configure.SysConf.RedisAddr.Ip, configure.SysConf.RedisAddr.Port),
			Password: "",
			Prefix:   "cn_",
		})

	if err != nil {
		panic(err)
	}
	err = Agent.CheckHealth()
	if err != nil {
		fmt.Println("failed to connect to redis server")
		return fmt.Errorf("failed to connect to redis server")
	}

	return nil
}
