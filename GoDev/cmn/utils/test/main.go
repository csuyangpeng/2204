package main

import (
	"fmt"
	"lite5gc/cmn/utils"
	"time"
)

func main() {
	for i := 0; i <= 32; i++ {
		utils.UpdateRestartCounter("amf_rst.dat")
		count := utils.GetRestartCounter()
		fmt.Println(count)
		time.Sleep(time.Second * 1)
	}
}
