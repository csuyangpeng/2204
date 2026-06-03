package main

import (
	"context"
	"fmt"
	tm "lite5gc/cmn/timermgr"
	"math/rand"
	"reflect"
	"time"
)

var counter int

func display(params interface{}) {
	fmt.Println("timeout, call display func")
	args := reflect.ValueOf(params) //interface to value, which is a slice

	lens := args.Len()
	if lens == 2 {
		name := reflect.ValueOf(args.Index(0).Interface()).String()
		age := reflect.ValueOf(args.Index(1).Interface()).Int()
		fmt.Printf("My name is %s, I am %d year's old.\n", name, age)
	}

	//fmt.Printf("Timerout : %d times\n", timerCoutner.IncrementAndGet())
}
func main() {
	fmt.Println("start")
	tMgr := tm.NewTimerMgr(context.Background(), 10, 500)
	counter = 0
	//Do the test for Period timer
	timeOutCB := tm.NewOnTimeOut(display, "nina", 666)

	tMgr.AddPeriodTimer(2, timeOutCB)
	//tMgr.AddAfterTimer(2, timeOutCB)

	time.Sleep(5 * time.Second)

}
func main_1() {

	tMgr := tm.NewTimerMgr(context.Background(), 10, 500)

	//Do the test for Period timer
	timeOutCB := tm.NewOnTimeOut(display, "nina", 666)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		num := time.Duration(rand.Intn(20) + 1)
		_ = tMgr.AddPeriodTimer(num, timeOutCB)
		// fmt.Printf("add timer %d, rand = %d .\n", timerId, num)
		time.Sleep(time.Microsecond * 20)
	}

	//Do the test for After timer
	timeOutCB1 := tm.NewOnTimeOut(display, "aaron", 777)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		num := time.Duration(rand.Intn(20) + 1)
		_ = tMgr.AddAfterTimer(num, timeOutCB1)
		// fmt.Printf("add timer %d, rand = %d .\n", timerId, num)
		time.Sleep(time.Microsecond * 20)
	}

	//monitor
	go func() {
		for {
			time.Sleep(2 * time.Second)

			fmt.Println("***Timer Manger PM Info****")
			fmt.Println("----------------------------------------------------------")
			fmt.Printf("Total timer Currently           : %d\n", tMgr.Size())
			fmt.Printf("After timer Add counter         : %d\n", tMgr.AfterTimerAddCounter.Get())
			fmt.Printf("After timer Timerout counter    : %d\n", tMgr.AfterTimerTimeroutCounter.Get())
			fmt.Printf("After timer Cancel counter      : %d\n", tMgr.AfterTimerCancelCounter.Get())
			fmt.Printf("Period timer Add counter        : %d\n", tMgr.PeriodTimerAddCounter.Get())
			fmt.Printf("Period timer Cancel counter     : %d\n", tMgr.PeriodTimerCancelCounter.Get())
			fmt.Printf("Period timer Timerout counter   : %d\n", tMgr.PeriodTimerTimeroutCounter.Get())
			fmt.Printf("Timer Push Heap counter         : %d\n", tMgr.TimerPushHeapCounter.Get())
			fmt.Printf("Timer Pop Heap counter          : %d\n", tMgr.TimerPopHeapCounter.Get())
			fmt.Println("----------------------------------------------------------")
		}
	}()

	// wg.Wait()
	time.Sleep(time.Second * 1000)
}
