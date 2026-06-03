package rlogger

import (
	"fmt"
	"os"
	_ "sort"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

//type Logger struct {
//	// The logs are `io.Copy`'d to this in a mutex. It's common to set this to a
//	// file, or leave it default which is `os.Stderr`. You can also set this to
//	// something more adventurous, such as logging to Kafka.
//	Out io.Writer
//	// Hooks for the logger instance. These allow firing events based on logging
//	// levels and log entries. For example, to send errors to an error tracking
//	// service, log to StatsD or dump the core on fatal errors.
//	Hooks LevelHooks
//	// All log entries pass through the formatter before logged to Out. The
//	// included formatters are `TextFormatter` and `JSONFormatter` for which
//	// TextFormatter is the default. In development (when a TTY is attached) it
//	// logs with colors, but to a file it wouldn't. You can easily implement your
//	// own that implements the `Formatter` interface, see the `README` or included
//	// formatters for examples.
//	Formatter Formatter
//
//	// Flag for whether to log caller info (off by default)
//	ReportCaller bool
//
//	// The logging level the logger should log at. This is typically (and defaults
//	// to) `logrus.Info`, which allows Info(), Warn(), Error() and Fatal() to be
//	// logged.
//	Level Level
//	// Used to sync writing to the log. Locking is enabled by Default
//	mu MutexWrap
//	// Reusable empty entry
//	entryPool sync.Pool
//	// Function to exit the application, defaults to `os.Exit()`
//	ExitFunc exitFunc
//}

// time="2020-11-13 11:17:33" level=info exe=main pid=22232 mod=amf_idmgr file="main.go:17" msg="f in main.go"
var order = map[string]int{
	"time":  1,
	"level": 2,
	"exe":   3,
	"pid":   4,
	"mod":   5,
	"file":  6,
	"msg":   7,
}

var logger = logrus.New()
var pid = getpid()
var procname = getprocname()

func init() {
	logger.Out = os.Stderr
	//setReportCaller(true)
	setLogFormatter(true)
}

// set log prefix
func SetPrefix(prefix string) {
	//l.mu.Lock()
	//defer l.mu.Unlock()
	//l.prefix = prefix
}

func setOutput(ifce interface{}) {
	switch t := ifce.(type) {
	case string:
		logger.SetOutput(&lumberjack.Logger{
			Filename:   t,
			MaxSize:    50, // megabytes
			MaxBackups: 3,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
		})
		setLogFormatter(false)
	case *os.File:
		logger.SetOutput(t)
		setLogFormatter(true)
		// to do for other output engineer
	default:
		fmt.Println("Not support", t)
	}
}

func trace(mod ModuleTag, l logrus.Level, f interface{}, v ...interface{}) {
	switch l {
	case logrus.DebugLevel:
		debugTrace(mod, f, v...)
	case logrus.InfoLevel:
		infoTrace(mod, f, v...)
	case logrus.WarnLevel:
		warnTrace(mod, f, v...)
	case logrus.ErrorLevel:
		errorTrace(mod, f, v...)
	case logrus.FatalLevel:
		fatalTrace(mod, f, v...) //log之后会调用os.Exit(1)
	case logrus.PanicLevel:
		panicTrace(mod, f, v...) //log之后会panic()
	}
}

// get log level
func getLogLevel() logrus.Level {
	return logger.Level
	//logger.Level = level
}

func translateLevel(level LogLevel) (logrus.Level, error) {
	switch level {
	case TRACE:
		return logrus.TraceLevel, nil
	case DEBUG:
		return logrus.DebugLevel, nil
	case WARN:
		return logrus.WarnLevel, nil
	case INFO:
		return logrus.InfoLevel, nil
	case ERROR:
		return logrus.ErrorLevel, nil
	case FATAL:
		return logrus.FatalLevel, nil
	case PANIC:
		return logrus.PanicLevel, nil
	}

	var l logrus.Level
	return l, fmt.Errorf("not a valid logrus Level: %q", level)
}

//// lite5gcFormatter implement interface Formatter
//type lite5gcFormatter struct {
//	Prefix string
//	Suffix string
//}
//
//// Format implement the Formatter interface
//func (mf *lite5gcFormatter) Format(entry *logrus.Entry) ([]byte, error) {
//	var b *bytes.Buffer
//	if entry.Buffer != nil {
//		b = entry.Buffer
//	} else {
//		b = &bytes.Buffer{}
//	}
//	// entry.Message 就是需要打印的日志
//	b.WriteString(fmt.Sprintf("%s - %s - %s", mf.Prefix, entry.Message, mf.Suffix))
//	return b.Bytes(), nil
//}

func setLogFormatter(color bool) {
	//formatter := &logrus.TextFormatter{
	//	DisableColors:   true,
	//	TimestampFormat: "2006-01-02 15:04:05",
	//FieldMap: logrus.FieldMap{
	//	        "time":  "@timestamp",
	//	        "level": "@level",
	//	        //"file":  "@file",
	//	        "msg":   "@message"},
	//	SortingFunc: func(keys []string) {
	//		sort.Slice(keys, func(i, j int) bool {
	//			return order[keys[i]] < order[keys[j]]
	//		})
	//	},
	//}
	formatter := &nested.Formatter{
		TimestampFormat: "2006-01-02 15:04:05.0000",
		NoFieldsSpace:   true,
		HideKeys:        true,
		NoColors:        true,
		FieldsOrder:     []string{"time", "level", "exe", "pid", "mod", "file", "msg"},
	}
	if color == true {
		formatter.NoColors = false
		formatter.NoFieldsColors = true
	}

	logger.Formatter = formatter
}

func setReportCaller(b bool) {
	logger.SetReportCaller(b)
}

func debugTrace(mod ModuleTag, f interface{}, args ...interface{}) {
	if getLogLevel() >= logrus.DebugLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Data["pid"] = pid
		entry.Data["exe"] = procname
		entry.Data["mod"] = mod
		if args == nil {
			entry.Debug(f)
		} else {
			s := f.(string)
			entry.Debugf(s, args...)
		}
	}
}

// Fields type
//type Fields logrus.Fields
//func DebugWithFields(l interface{}, f Fields) {
//	if getLogLevel() >= logrus.DebugLevel {
//		entry := logger.WithFields(logrus.Fields(f))
//		entry.Data["file"] = fileInfo(2)
//		entry.Debug(l)
//	}
//}

func infoTrace(mod ModuleTag, f interface{}, args ...interface{}) {
	if getLogLevel() >= logrus.InfoLevel {
		entry := logger.WithFields(logrus.Fields{})
		//entry := logger.WithField("topics", []string{"topic_test"})
		entry.Data["file"] = fileInfo(2)
		entry.Data["pid"] = pid
		entry.Data["exe"] = procname
		entry.Data["mod"] = mod
		if args == nil {
			entry.Info(f)
		} else {
			s := f.(string)
			entry.Infof(s, args...)
		}
	}
}

//func InfoWithFields(l interface{}, f Fields) {
//	if getLogLevel() >= logrus.InfoLevel {
//		entry := logger.WithFields(logrus.Fields(f))
//		entry.Data["file"] = fileInfo(2)
//		entry.Info(l)
//	}
//}

func warnTrace(mod ModuleTag, f interface{}, args ...interface{}) {
	if getLogLevel() >= logrus.WarnLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Data["pid"] = pid
		entry.Data["exe"] = procname
		entry.Data["mod"] = mod
		if args == nil {
			entry.Warn(f)
		} else {
			s := f.(string)
			entry.Warnf(s, args...)
		}
	}
}

//func warnWithFields(l interface{}, f Fields) {
//	if logger.Level >= logrus.WarnLevel {
//		entry := logger.WithFields(logrus.Fields(f))
//		entry.Data["file"] = fileInfo(2)
//		entry.Warn(l)
//	}
//}

func errorTrace(mod ModuleTag, f interface{}, args ...interface{}) {
	if getLogLevel() >= logrus.ErrorLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Data["pid"] = pid
		entry.Data["exe"] = procname
		entry.Data["mod"] = mod
		if args == nil {
			entry.Error(f)
		} else {
			s := f.(string)
			entry.Errorf(s, args...)
		}
	}
}

//func ErrorWithFields(l interface{}, f Fields) {
//	if logger.Level >= logrus.ErrorLevel {
//		entry := logger.WithFields(logrus.Fields(f))
//		entry.Data["file"] = fileInfo(2)
//		entry.Error(l)
//	}
//}

func fatalTrace(mod ModuleTag, f interface{}, args ...interface{}) {
	if getLogLevel() >= logrus.FatalLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Data["pid"] = pid
		entry.Data["exe"] = procname
		entry.Data["mod"] = mod
		if args == nil {
			entry.Fatal(f)
		} else {
			s := f.(string)
			entry.Fatalf(s, args...)
		}
	}
}

//func FatalWithFields(l interface{}, f Fields) {
//	if logger.Level >= logrus.FatalLevel {
//		entry := logger.WithFields(logrus.Fields(f))
//		entry.Data["file"] = fileInfo(2)
//		entry.Fatal(l)
//	}
//}

func panicTrace(mod ModuleTag, f interface{}, args ...interface{}) {
	//if getLogLevel() >= logrus.PanicLevel {
	entry := logger.WithFields(logrus.Fields{})
	entry.Data["file"] = fileInfo(2)
	entry.Data["pid"] = pid
	entry.Data["exe"] = procname
	entry.Data["mod"] = mod
	if args == nil {
		entry.Panic(f)
	} else {
		s := f.(string)
		entry.Panicf(s, args...)
	}
	//}
}

//func PanicWithFields(l interface{}, f Fields) {
//	if logger.Level >= logrus.PanicLevel {
//		entry := logger.WithFields(logrus.Fields(f))
//		entry.Data["file"] = fileInfo(2)
//		entry.Panic(l)
//	}
//}
