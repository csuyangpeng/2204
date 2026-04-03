package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var restartCounter uint32

func UpdateRestartCounter(filename string) {

	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	defer f.Close()

	contents, err := ioutil.ReadAll(f)
	if err == nil && len(contents) > 0 {
		//get the first byte for restart counter
		strCont := string(contents)
		counter, err := strconv.Atoi(strings.TrimSpace(strCont))
		if err != nil {
			fmt.Println(err.Error())
			panic(err)
		}

		restartCounter = uint32(counter)
	} else {
		restartCounter = 0
	}

	if restartCounter >= 15 {
		restartCounter = 0
	} else {
		restartCounter++
	}

	wdata := []byte(fmt.Sprintf("%2d", restartCounter))
	if _, err := f.WriteAt(wdata, 0); err == nil {
		fmt.Println("update the file with ", restartCounter)
	}
}

func GetRestartCounter() uint32 {
	return restartCounter
}
