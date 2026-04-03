package main

import (
	"fmt"
	_ "github.com/gfremex/logrus-kafka-hook"
	"lite5gc/cmn/rlogger"
	"time"
)

const moduleTag rlogger.ModuleTag = "cmn"

//func init() {
//	rlogger.SetLogOutput("logrus.log")
//	//rlogger.SetLogOutput(os.Stderr)
//	rlogger.SetLogLevel(rlogger.TRACE)
//	//hook, err := lkh.NewKafkaHook(
//	//	"kh",
//	//	[]logrus.Level{logrus.InfoLevel},
//	//	&logrus.TextFormatter{},
//	//	[]string{"10.18.1.2:9092",},
//	//)
//	//if err != nil {
//	//	return
//	//}
//	//rlogger.AddHook(hook)
//
//}

func f() {
	rlogger.FuncEntry(moduleTag, nil)
	for {
		rlogger.Trace(moduleTag, rlogger.INFO, nil, "f in main.go")
	}
}
type LoggerConfig struct {
	Level   string
	Path    string
	Control bool
}

func main() {
	config := LoggerConfig{
		Level:   "info",
		Path:    "log.txt",
		Control: true,
	}
	err := rlogger.Initilize(config.Control, config.Path, config.Level)
	if err != nil {
		panic(err)
	}

	rlogger.Trace(moduleTag, rlogger.ERROR, nil, "log error in amf config")
	rlogger.Trace(moduleTag, rlogger.WARN, nil, "log warn in amf")
	rlogger.Trace(moduleTag, rlogger.INFO, nil, "log info in cmn")
	num := 3
	str := fmt.Sprintf("log info in %v cmn", num)
	fmt.Println(str)
	rlogger.Trace(moduleTag, rlogger.INFO, nil, "log info in %v cmn", num)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "log info in %v cmn", num)
	rlogger.Trace(moduleTag, rlogger.WARN, nil, "log info in %v cmn", num)

	//go f()
	time.Sleep(time.Second)
}
