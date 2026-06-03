package pm

import (
	"github.com/rcrowley/go-metrics"
)

// constructs a new StandardCounter, and register the given metric under the given name.
func CreateCounter(r metrics.Registry, name string) metrics.Counter {
	c := metrics.NewCounter()
	_ = r.Register(name, c)
	return c
}

// get the counter by the given name
func GetCounter(r metrics.Registry, name string) int64 {
	c := metrics.GetOrRegisterCounter(name, r)
	return c.Count()
}

// increments the counter by the given amount
func IncCounter(c metrics.Counter, number int64) {
	c.Inc(number)
}

// increments the counter by the given amount
func PegCounter(c metrics.Counter) {
	c.Inc(1)
}

// decrements the counter by the given amount
func DecCounter(c metrics.Counter, number int64) {
	c.Dec(number)
}

// decrements the counter by the given amount
func UnPegCounter(c metrics.Counter) {
	c.Dec(1)
}

// clear the counter
func ClearCounter(c metrics.Counter) {
	c.Clear()
}
