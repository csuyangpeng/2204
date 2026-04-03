package rlogger

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func getprocname() string {
	exe, err := os.Executable()
	if err != nil {
		fmt.Println("getprocname ", err)
	}
	return filepath.Base(exe)
}

func getpid() int {
	return os.Getpid()
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip + 2) // this param `2` should vary as the caller' stack
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func funcEntryCallerName() string {
	funcName, _, _, ok := runtime.Caller(2)
	if ok {
		fn := runtime.FuncForPC(funcName).Name()
		fn = filepath.Ext(fn)                   // .foo
		fn = strings.TrimPrefix(fn, ".") + "()" // foo
		return fn
	}
	return ""
}