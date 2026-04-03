package metric

import (
	"fmt"
	"github.com/rcrowley/go-metrics"
	"log"
	"os"
	"time"
)

// 创建统计指标map
func NewMetricsMap() Registry {
	return metrics.NewRegistry()
}

// 创建统计变量 counter
func NewCounter(name string, MetricsMap Registry) Counter {
	c := metrics.NewRegisteredCounter(name, MetricsMap)

	return c
}

// 创建统计变量 meter
func NewMeter(name string, MetricsMap Registry) Meter {
	c := metrics.NewRegisteredMeter(name, MetricsMap)

	return c
}
func MeterStop(name string, MetricsMap Registry) error {
	m := GetMeter(name, MetricsMap)
	if m == nil {
		return fmt.Errorf("meter is nill,%v", name)
	}
	m.Stop()
	return nil
}

func Log(r Registry, freq time.Duration) {
	metrics.Log(r, freq, log.New(os.Stdout, "metrics: ", log.Lmicroseconds))
	//r.Get(metric.UpLinkReceivedPacketPerS).(metrics.Meter).Count()
}

func CountersCreate(Metrics []string, r Registry) error {
	//MetricItems := NewMetricsMap()
	for _, v := range Metrics {
		NewCounter(v, r)
	}
	return nil
}

func MetersCreate(Metrics []string, r Registry) error {
	//MetricItems := NewMetricsMap()
	for _, v := range Metrics {
		NewMeter(v, r)
	}
	return nil
}

// notes: Meter required release
func MetersRelease(Metrics []string, r Registry) error {
	for _, v := range Metrics {
		MeterStop(v, r)
	}
	return nil
}

//type UpfCountMapT metrics.Registry
//type UpfCounterT metrics.Counter

type Counter interface {
	metrics.Counter
}

type Meter interface {
	metrics.Meter
}

type Registry interface {
	metrics.Registry
}

var DefaultMetricsMap metrics.Registry

func init() {
	DefaultMetricsMap = NewMetricsMap()
}

func Get(name string, metricItems Registry) Counter {
	v := metricItems.Get(name)
	v1, _ := v.(Counter)
	return v1
}

func GetMeter(name string, metricItems Registry) Meter {
	v := metricItems.Get(name)
	v1, _ := v.(Meter)
	return v1
}

// add 1
func MetricItemAdd(name string, MetricItems Registry) {
	counter := Get(name, MetricItems) //metric.UpLinkSessionReceivedPacket
	counter.Inc(1)
}

// sub 1
func MetricItemSub(name string, MetricItems Registry) {
	counter := Get(name, MetricItems) //metric.UpLinkSessionReceivedPacket
	counter.Dec(1)
}

// Clear
func MetricItemClear(name string, MetricItems Registry) {
	counter := Get(name, MetricItems) //metric.UpLinkSessionReceivedPacket
	counter.Clear()
}

// count
func MetricItemCount(name string, MetricItems Registry) {
	counter := Get(name, MetricItems) //metric.UpLinkSessionReceivedPacket
	counter.Count()
}

// add x
func MetricItemAddX(name string, MetricItems Registry, x int64) {
	counter := Get(name, MetricItems) //metric.UpLinkSessionReceivedPacket
	counter.Inc(x)
}
