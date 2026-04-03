package metric_test

import (
	"errors"
	"fmt"
	"github.com/rcrowley/go-metrics"
	"lite5gc/cmn/metric"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestNewCounter(t *testing.T) { /*
		metric.DefaultMetricsMap = metric.NewMetricsMap()

		fmt.Println(unsafe.Sizeof(metric.DefaultMetricsMap))
		fmt.Printf("%T\n", metric.DefaultMetricsMap)

		c := metric.NewCounter(upfmetric.UpLinkReceivedPacket, metric.DefaultMetricsMap)
		c.Inc(1)
		fmt.Println(upfmetric.UpLinkReceivedPacket, c.Count())


		v := metric.DefaultMetricsMap.Get(upfmetric.UpLinkReceivedPacket)

		v1, _ := v.(metric.Counter)
		v1.Inc(2)
		fmt.Println(upfmetric.UpLinkReceivedPacket, v1.Count())

		v1.Dec(1)
		fmt.Println(upfmetric.UpLinkReceivedPacket, v1.Count())
		v1.Clear()
		fmt.Println(upfmetric.UpLinkReceivedPacket, v1.Count(), c.Count())
		metric.DefaultMetricsMap.UnregisterAll()

		c = metric.NewCounter(upfmetric.UpLinkReceivedPacket, metric.DefaultMetricsMap)
		fmt.Println(upfmetric.UpLinkReceivedPacket, c, metric.DefaultMetricsMap.GetAll())
		println(c)
		c.Inc(1)
		c = metric.NewCounter(upfmetric.UpLinkSendPacket, metric.DefaultMetricsMap)
		fmt.Println(upfmetric.UpLinkSendPacket, c, metric.DefaultMetricsMap.GetAll())
		println(c)
		c.Inc(1)
		// 重复返回err
		c = metric.NewCounter(upfmetric.UpLinkReceivedPacket, metric.DefaultMetricsMap)
		// 没有更新到UpfCounter
		fmt.Println(len(metric.DefaultMetricsMap.GetAll()), metric.DefaultMetricsMap.GetAll())
		println(c)
		metric.DefaultMetricsMap.Each(func(s string, i interface{}) {
			fmt.Println("key:", s, "value:", i.(metrics.Counter).Count())
		})
		metric.DefaultMetricsMap.RunHealthchecks() // null

		metrics.Log(metric.DefaultMetricsMap,
			5*time.Second,
			log.New(os.Stdout, "\nmetrics: ", log.Lmicroseconds))*/

}

const fanout = 5

func TestNewMetric(t *testing.T) {
	m := metrics.NewMeter()
	metric.DefaultMetricsMap.Register("send pps", m)
	fmt.Println("send pps", metric.DefaultMetricsMap.GetAll())
	//m.Mark(1)
	for i := 0; i < fanout; i++ {
		go func() {
			for {
				m.Mark(1)
				time.Sleep(1e9)
			}
		}()
		go func() {
			for {
				m.Mark(1)
				time.Sleep(1e9)
			}
		}()
	}
	fmt.Println(m.Count())
	fmt.Println(m.Snapshot().Count())
	fmt.Println(m.Rate1())
	fmt.Println(m.RateMean())

	fmt.Println("send pps", metric.DefaultMetricsMap.GetAll())

	metrics.Log(metric.DefaultMetricsMap, 3e9, log.New(os.Stderr, "metrics: ", log.Lmicroseconds))
}

func TestHistogram(t *testing.T) {
	s := metrics.NewExpDecaySample(1028, 0.015)
	h := metrics.NewHistogram(s)
	metric.DefaultMetricsMap.Register("bang", h)
	for i := 0; i < fanout; i++ {
		go func() {
			for {
				h.Update(19)
				time.Sleep(1e9)
			}
		}()
		go func() {
			for {
				h.Update(47)
				time.Sleep(1e9)
			}
		}()
	}
	fmt.Println("send pps", metric.DefaultMetricsMap.GetAll())
	metrics.Log(metric.DefaultMetricsMap, 3e9, log.New(os.Stderr, "metrics: ", log.Lmicroseconds))
}

func TestHealthcheck(t *testing.T) {
	hc := metrics.NewHealthcheck(func(h metrics.Healthcheck) {
		if 0 < rand.Intn(2) {
			h.Healthy()
		} else {
			h.Unhealthy(errors.New("baz warning 100%cpu"))
		}
	})
	metric.DefaultMetricsMap.Register("baz", hc)
	fmt.Println("Healthcheck", metric.DefaultMetricsMap.GetAll())
	metrics.Log(metric.DefaultMetricsMap, 3e9, log.New(os.Stderr, "metrics: ", log.Lmicroseconds))
}

// 统计代码块的执行时间：次数/秒 与单个执行的统计时间
func TestTimer(t *testing.T) {
	t1 := metrics.NewTimer()
	metric.DefaultMetricsMap.Register("hooah", t1)
	for i := 0; i < fanout; i++ {
		go func() {
			for {
				t1.Time(func() { time.Sleep(1e9) })
			}
		}()
		go func() {
			for {
				t1.Time(func() { time.Sleep(1e9) })
			}
		}()
	}
	fmt.Println("send pps", metric.DefaultMetricsMap.GetAll())
	metrics.Log(metric.DefaultMetricsMap, 3e9, log.New(os.Stderr, "metrics: ", log.Lmicroseconds))
}
