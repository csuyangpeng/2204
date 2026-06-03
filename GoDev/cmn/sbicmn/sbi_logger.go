package sbicmn

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func LoggerToFile(fileName string) gin.HandlerFunc {

	// logger file
	//fileName := path.Join(logFilePath, logFileName)
	// write file
	src, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	//Initialize
	logger := logrus.New()
	// set logger output
	logger.Out = src
	// set logger level
	logger.SetLevel(logrus.DebugLevel)
	// set rotatelogs
	logWriter, err := rotatelogs.New(
		// file name after split
		fileName+".%Y%m%d.log",
		// soft link to the latest file
		rotatelogs.WithLinkName(fileName),
		// max save time for logger, 7 days
		rotatelogs.WithMaxAge(7*24*time.Hour),
		// logger split time slots, 1 day
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap,
		&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	// new hook
	logger.AddHook(lfHook)
	return func(c *gin.Context) {
		// start time
		startTime := time.Now()
		// process time
		c.Next()
		// finished time
		endTime := time.Now()
		// impement time
		latencyTime := endTime.Sub(startTime)
		// request method
		reqMethod := c.Request.Method
		// request url
		reqUri := c.Request.RequestURI
		// status
		statusCode := c.Writer.Status()
		// request ip address
		clientIP := c.ClientIP()
		// logger formate
		//logger.WithFields(logrus.Fields{
		//	"status_code":  statusCode,
		//	"latency_time": latencyTime,
		//	"client_ip":    clientIP,
		//	"req_method":   reqMethod,
		//	"req_uri":      reqUri,
		//}).Info()
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}

// logger to MongoDB
func LoggerToMongo() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

// logger to ES
func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

// logger to MQ
func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
