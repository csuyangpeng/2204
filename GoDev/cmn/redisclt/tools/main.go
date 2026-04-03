/** Copyright(C),2020-2022
* Author: zmj
* Date: 12/9/20 3:43 PM
* Description:
 */
package main

import (
	"fmt"
	"lite5gc/cmn/redisclt"
	"time"
)

func main() {
	// create the redis clt for communication with amf main process
	clt, err := redisclt.New(
		redisclt.Options{
			Addr:     "10.18.1.52:6379",
			Password: "",
			Prefix:   "cn_",
		})

	if err != nil {
		panic(err)
	}
	err = clt.CheckHealth()
	if err != nil {
		fmt.Println("failed to connect to redis server")
		return
	}

	for i := 0; i < 10; i++ {
		clt.LPush("amf_sc_1", "test,test")
		time.Sleep(time.Second * 1)
	}

}
