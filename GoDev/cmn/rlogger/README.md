## HOW TO use rlogger

Exam: `example/main.go` is a simple example using rlogger.

Set rlogger to file.  
`rlogger.SetLogOutput("logrus.log")`

Set rlogger to console.  
`rlogger.SetLogOutput(os.Stderr)`

Set rlogger stand level.  
`rlogger.SetLogLevel(rlogger.TRACE)`

Log msg.  
`rlogger.Trace(moduleName, rlogger.INFO, nil, "log info in cmn")`  
`rlogger.Trace(moduleName, rlogger.INFO, nil, "log info in %v cmn", num)`