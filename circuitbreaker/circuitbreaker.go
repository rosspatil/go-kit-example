package circuitbreaker

import (
	"fmt"
	"time"

	"github.com/sony/gobreaker"
)

// Cb -
var Cb *gobreaker.CircuitBreaker

const (
	interval = 30
)

func init() {
	st := gobreaker.Settings{}
	st.Interval = time.Second * interval
	st.Name = "go-kit-example-circuitbreaker"
	st.MaxRequests = 1
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		fmt.Println(counts)
		return counts.ConsecutiveFailures > 2
	}
	st.OnStateChange = func(name string, from gobreaker.State, to gobreaker.State) {
		fmt.Println(name, from.String(), to.String())
	}
	Cb = gobreaker.NewCircuitBreaker(st)
}
