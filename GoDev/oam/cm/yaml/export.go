// Package yaml implements YAML support for the Go language.
// Based on gopkg.in/yaml.v2.
package yaml

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"sync"
)

var rwlock sync.RWMutex

// Load out from yaml file
func Load(file string, out interface{}) {
	rlogger.Trace(types.ModuleOamCm, rlogger.DEBUG, nil, "load data from [yaml]: %s\n", file)
	rwlock.RLock()
	defer rwlock.RUnlock()
	content, err := ioutil.ReadFile(file)
	checkError(err)

	err = yaml.Unmarshal(content, out)
	checkError(err)
}

// Dump in to yaml file
// to do RWlock
func Dump(in interface{}, file string) {
	rlogger.Trace(types.ModuleOamCm, rlogger.DEBUG, nil, "dump data to [yaml]: %s\n", file)
	rwlock.Lock()
	defer rwlock.Unlock()
	data, err := yaml.Marshal(in)
	checkError(err)

	err = ioutil.WriteFile(file, data, 0777)
	checkError(err)
}
