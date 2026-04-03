package rlogger

import (
	"fmt"
	lkh "github.com/gfremex/logrus-kafka-hook"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

// LogLevel type
type LogLevel int32

// ModuleTag type
type ModuleTag string

// LogLevel enum defs
const (
	PANIC LogLevel = iota
	FATAL
	ERROR
	WARN
	INFO
	DEBUG
	TRACE
)

var lmap = map[string]LogLevel{
	"trace": TRACE,
	"debug": DEBUG,
	"info":  INFO,
	"warn":  WARN,
	"error": ERROR,
	"fatal": FATAL,
	"panic": PANIC,
}

func Initialize(ctrl bool, path, level string) error {
	if ctrl == true {
		if path == "" || path == "." {
			return fmt.Errorf("if choose log to file. need config log path ?")
		}
		setOutput(path)
	} else {
		setOutput(os.Stderr)
	}

	if v, ok := lmap[level]; ok {
		SetLogLevel(v)
	} else {
		return fmt.Errorf("initilize log level fail.")
	}
	return nil
}

// You can log to file by `SetOutput("filename")` or SetOutput(os.Stderr)
// or leave it default which is `os.Stderr`
func SetLogOutput(ifce interface{}) {
	setOutput(ifce)
}

// default： InfoLevel
func SetLogLevel(l LogLevel) {
	level, err := translateLevel(l)
	if err != nil {
		fmt.Println(err)
		return
	}
	logger.SetLevel(level)
}

// current log level
func GetLogLevel() LogLevel {
	return LogLevel(logger.GetLevel())
}

func IsLevelDebug() bool {
	return DEBUG <= GetLogLevel()
}

//func Register(mod string) int {
//
//	return
//}

// logger.Trace(ModuleTag, types.INFO, nil, "logger started")
// or 	logger.Trace(logger.MODULE_CMN_ID, logger.INFO, nil,"log info in %v cmn", num)
func Trace(mod ModuleTag, level LogLevel, ctxt interface{}, f interface{}, v ...interface{}) {
	l, err := translateLevel(level)
	if err != nil {
		return
	}
	// to do mod switch

	// to do ctxt

	// to do speeddown

	// to do probability

	// logrus will do this automatically
	//if l > logger.GetLevel() {
	//	return
	//}
	trace(mod, l, f, v...)
}

func FuncEntry(mod ModuleTag, ctxt interface{}) {
	str := fmt.Sprintf("%s entered", funcEntryCallerName())
	l, _ := translateLevel(INFO)
	trace(mod, l, str)
}

func AddHook(hook *lkh.KafkaHook) {
	logger.Hooks.Add(hook)
}

func LoggerToFile() gin.HandlerFunc {

	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()

		latencyTime := endTime.Sub(startTime)

		reqMethod := c.Request.Method

		reqUri := c.Request.RequestURI

		statusCode := c.Writer.Status()

		clientIP := c.ClientIP()

		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}
