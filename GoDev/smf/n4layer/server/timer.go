package server

import (
	"fmt"
	"time"
)

// T1
func StartTimer(t time.Duration, f func(interface{}), para interface{}) (*time.Timer, error) {
	fn := func() {
		fmt.Printf("Now time :       %v.\n", time.Now().Format("2006-01-02 15:04:05"))
		p := para.(string)
		fmt.Println(p)
		f(p)
	}

	tr := time.AfterFunc(t, fn)
	return tr, nil
}

func T1Handler(para interface{}) {
	p := para.(string)

	fmt.Printf("Now time :       %v.\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(p)

}
